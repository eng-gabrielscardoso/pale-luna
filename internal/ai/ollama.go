package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/eng-gabrielscardoso/pale-luna/internal/config"
)

type OllamaClient struct {
	config     *config.AIConfig
	httpClient *http.Client
	prompts    *PromptBuilder
}

type OllamaRequest struct {
	Model   string                 `json:"model"`
	Prompt  string                 `json:"prompt"`
	Stream  bool                   `json:"stream"`
	Options map[string]interface{} `json:"options,omitempty"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
	Error    string `json:"error,omitempty"`
}

func NewOllamaClient(cfg *config.AIConfig) *OllamaClient {
	return &OllamaClient{
		config: cfg,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
		prompts: NewPromptBuilder(),
	}
}

func (oc *OllamaClient) IsAvailable() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", oc.config.OllamaURL+"/api/version", nil)
	if err != nil {
		return false
	}

	resp, err := oc.httpClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func (oc *OllamaClient) ProcessCommand(input string, gameContext GameContext) (string, error) {
	if !oc.config.Enabled {
		return GetFallbackResponse(input, gameContext), nil
	}

	prompt := oc.prompts.BuildPrompt(input, gameContext)

	reqBody := OllamaRequest{
		Model:  oc.config.Model,
		Prompt: prompt,
		Stream: false,
		Options: map[string]interface{}{
			"temperature": oc.config.Temperature,
			"num_predict": oc.config.MaxTokens,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		if oc.config.FallbackEnabled {
			return GetFallbackResponse(input, gameContext), nil
		}
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), oc.config.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", oc.config.OllamaURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		if oc.config.FallbackEnabled {
			return GetFallbackResponse(input, gameContext), nil
		}
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := oc.httpClient.Do(req)
	if err != nil {
		if oc.config.FallbackEnabled {
			return GetFallbackResponse(input, gameContext), nil
		}
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if oc.config.FallbackEnabled {
			return GetFallbackResponse(input, gameContext), nil
		}
		return "", fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		if oc.config.FallbackEnabled {
			return GetFallbackResponse(input, gameContext), nil
		}
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if ollamaResp.Error != "" {
		if oc.config.FallbackEnabled {
			return GetFallbackResponse(input, gameContext), nil
		}
		return "", fmt.Errorf("API error: %s", ollamaResp.Error)
	}

	response := cleanAIResponse(ollamaResp.Response)
	if response == "" {
		return GetFallbackResponse(input, gameContext), nil
	}

	return response, nil
}

func cleanAIResponse(response string) string {
	response = removePrefix(response, "Pale Luna:")
	response = removePrefix(response, "Response:")
	response = removePrefix(response, "*")
	response = removeSuffix(response, "*")

	return strings.TrimSpace(response)
}

func removePrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return strings.TrimSpace(s[len(prefix):])
	}
	return s
}

func removeSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return strings.TrimSpace(s[:len(s)-len(suffix)])
	}
	return s
}
