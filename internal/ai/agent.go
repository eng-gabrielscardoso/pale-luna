package ai

import (
	"github.com/eng-gabrielscardoso/pale-luna/internal/config"
)

type AIAgent interface {
	ProcessCommand(input string, context GameContext) (string, error)
	IsAvailable() bool
}

type AgentManager struct {
	agent  AIAgent
	config *config.Config
}

func NewAgentManager(cfg *config.Config) *AgentManager {
	agent := NewOllamaClient(&cfg.AI)

	return &AgentManager{
		agent:  agent,
		config: cfg,
	}
}

func (am *AgentManager) ProcessInput(input string, context GameContext) string {
	// Try AI agent first
	if am.config.AI.Enabled && am.agent.IsAvailable() {
		response, err := am.agent.ProcessCommand(input, context)
		if err == nil && response != "" {
			return response
		}
	}

	return GetFallbackResponse(input, context)
}

func (am *AgentManager) IsAIAvailable() bool {
	return am.config.AI.Enabled && am.agent.IsAvailable()
}

func (am *AgentManager) GetStatus() map[string]interface{} {
	return map[string]interface{}{
		"ai_enabled":   am.config.AI.Enabled,
		"ai_available": am.IsAIAvailable(),
		"model":        am.config.AI.Model,
		"ollama_url":   am.config.AI.OllamaURL,
	}
}
