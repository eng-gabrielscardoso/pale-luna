package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	AI AIConfig
}

type AIConfig struct {
	Enabled         bool
	OllamaURL       string
	Model           string
	Timeout         time.Duration
	MaxTokens       int
	Temperature     float32
	FallbackEnabled bool
}

func Load() *Config {
	return &Config{
		AI: AIConfig{
			Enabled:         getEnvBool("PALE_LUNA_AI_ENABLED", true),
			OllamaURL:       getEnvString("PALE_LUNA_OLLAMA_URL", "http://localhost:11434"),
			Model:           getEnvString("PALE_LUNA_AI_MODEL", "llama3.2:3b"),
			Timeout:         getEnvDuration("PALE_LUNA_AI_TIMEOUT", 30*time.Second),
			MaxTokens:       getEnvInt("PALE_LUNA_AI_MAX_TOKENS", 150),
			Temperature:     getEnvFloat("PALE_LUNA_AI_TEMPERATURE", 0.8),
			FallbackEnabled: getEnvBool("PALE_LUNA_AI_FALLBACK", true),
		},
	}
}

func getEnvString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if parsed, err := time.ParseDuration(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}

func getEnvFloat(key string, defaultValue float32) float32 {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseFloat(value, 32); err == nil {
			return float32(parsed)
		}
	}
	return defaultValue
}
