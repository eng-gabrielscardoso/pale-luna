# Pale Luna - Makefile
# Simplified commands for development and deployment

.PHONY: help setup start stop restart status build run clean install-deps ollama-setup ollama-start ollama-stop ollama-status

# Default target
help:
	@echo "🌙 Pale Luna - Available Commands"
	@echo "================================"
	@echo ""
	@echo "Setup & Installation:"
	@echo "  setup          Complete setup (Ollama + dependencies + model)"
	@echo "  install-deps   Install Go dependencies"
	@echo ""
	@echo "Ollama Management:"
	@echo "  ollama-setup   Setup Ollama with Docker + install recommended model"
	@echo "  ollama-start   Start Ollama service"
	@echo "  ollama-stop    Stop Ollama service"
	@echo "  ollama-status  Check Ollama status"
	@echo ""
	@echo "Development:"
	@echo "  build          Build the application"
	@echo "  run            Run Pale Luna directly"
	@echo "  test           Run tests"
	@echo "  clean          Clean build artifacts"
	@echo ""
	@echo "Game Commands:"
	@echo "  play           Start Pale Luna (with Ollama check)"

# Complete setup
setup: install-deps ollama-setup
	@echo "✅ Pale Luna setup complete!"
	@echo "Run 'make play' to start the game"

# Install Go dependencies
install-deps:
	@echo "Installing Go dependencies..."
	go mod tidy
	go mod download

# Ollama management
ollama-setup:
	@echo "Setting up Ollama for Pale Luna..."
	./scripts/ollama-docker.sh setup

ollama-start:
	@echo "Starting Ollama service..."
	./scripts/ollama-docker.sh start

ollama-stop:
	@echo "Stopping Ollama service..."
	./scripts/ollama-docker.sh stop

ollama-restart:
	@echo "Restarting Ollama service..."
	./scripts/ollama-docker.sh restart

ollama-status:
	@echo "Checking Ollama status..."
	./scripts/ollama-docker.sh status

# Development commands
build:
	@echo "Building Pale Luna..."
	go build -o pale-luna ./cmd/main.go
	@echo "✅ Build complete: ./pale-luna"

run:
	@echo "Running Pale Luna..."
	go run ./cmd/main.go

test:
	@echo "Running tests..."
	go test ./internal/...

clean:
	@echo "Cleaning build artifacts..."
	rm -f pale-luna pale-luna.exe pale-luna-*
	@echo "✅ Clean complete"

# Game command with Ollama check
play:
	@echo "🌙 Starting Pale Luna..."
	@echo "Checking Ollama status first..."
	@./scripts/ollama-docker.sh status || (echo "❌ Ollama not ready. Run 'make ollama-setup' first." && exit 1)
	@echo "🎮 Starting game..."
	go run ./cmd/main.go

# Cross-platform builds
build-all: clean
	@echo "Building for all platforms..."
	GOOS=linux GOARCH=amd64 go build -o pale-luna-linux ./cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o pale-luna.exe ./cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -o pale-luna-mac ./cmd/main.go
	@echo "✅ Cross-platform builds complete"

# Docker operations (direct)
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f ollama
