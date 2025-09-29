#!/usr/bin/env bash

# Pale Luna - Ollama Docker Management Script
# This script manages the Ollama service for Pale Luna

set -e

COMPOSE_FILE="docker compose.yml"
RECOMMENDED_MODEL="llama3.2:3b"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_header() {
    echo -e "${BLUE}"
    echo "=================================================="
    echo "    🌙 Pale Luna - Ollama Management 🌙"
    echo "=================================================="
    echo -e "${NC}"
}

print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

check_docker() {
    if ! command -v docker &> /dev/null; then
        print_error "Docker is not installed. Please install Docker first."
        exit 1
    fi

    if ! command -v docker compose &> /dev/null && ! docker compose version &> /dev/null; then
        print_error "Docker Compose is not installed. Please install Docker Compose first."
        exit 1
    fi
}

start_ollama() {
    print_status "Starting Ollama service..."
    docker compose up -d ollama

    print_status "Waiting for Ollama to be ready..."
    sleep 5

    # Wait for Ollama to be available
    for i in {1..30}; do
        if curl -s http://localhost:11434/api/version > /dev/null 2>&1; then
            print_status "Ollama is ready!"
            return 0
        fi
        echo -n "."
        sleep 2
    done

    print_error "Ollama failed to start properly"
    return 1
}

stop_ollama() {
    print_status "Stopping Ollama service..."
    docker compose down
    print_status "Ollama stopped."
}

install_model() {
    local model=${1:-$RECOMMENDED_MODEL}

    print_status "Installing model: $model"
    print_warning "This may take several minutes depending on your internet connection..."

    docker compose exec ollama ollama pull "$model"

    if [ $? -eq 0 ]; then
        print_status "Model $model installed successfully!"
    else
        print_error "Failed to install model $model"
        return 1
    fi
}

list_models() {
    print_status "Installed models:"
    docker compose exec ollama ollama list
}

check_status() {
    print_status "Checking Ollama status..."

    if docker compose ps ollama | grep -q "Up"; then
        print_status "✅ Ollama container is running"

        if curl -s http://localhost:11434/api/version > /dev/null 2>&1; then
            print_status "✅ Ollama API is accessible"
            echo -e "${GREEN}Ollama is ready for Pale Luna!${NC}"
        else
            print_warning "⚠️  Ollama container is running but API is not accessible"
        fi
    else
        print_warning "❌ Ollama container is not running"
        echo "Run: $0 start"
    fi
}

show_logs() {
    print_status "Showing Ollama logs (press Ctrl+C to exit)..."
    docker compose logs -f ollama
}

show_help() {
    print_header
    echo "Usage: $0 [COMMAND]"
    echo ""
    echo "Commands:"
    echo "  start              Start Ollama service"
    echo "  stop               Stop Ollama service"
    echo "  restart            Restart Ollama service"
    echo "  status             Check Ollama status"
    echo "  install [MODEL]    Install AI model (default: $RECOMMENDED_MODEL)"
    echo "  models             List installed models"
    echo "  logs               Show Ollama logs"
    echo "  setup              Complete setup (start + install recommended model)"
    echo "  help               Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 setup                    # Complete setup for Pale Luna"
    echo "  $0 install mistral:7b       # Install Mistral 7B model"
    echo "  $0 install llama3.2:8b      # Install Llama 3.2 8B model"
    echo ""
    echo "Recommended models for Pale Luna:"
    echo "  - llama3.2:3b  (Fast, ~2GB)"
    echo "  - mistral:7b   (Balanced, ~4GB)"
    echo "  - llama3.2:8b  (Best quality, ~5GB)"
}

setup_complete() {
    print_header
    print_status "Setting up Ollama for Pale Luna..."

    start_ollama
    if [ $? -eq 0 ]; then
        install_model "$RECOMMENDED_MODEL"
        if [ $? -eq 0 ]; then
            echo ""
            print_status "🎉 Setup complete! Pale Luna is ready."
            echo -e "${GREEN}You can now run: go run cmd/main.go${NC}"
        fi
    fi
}

# Main script logic
case "${1:-help}" in
    "start")
        check_docker
        start_ollama
        ;;
    "stop")
        check_docker
        stop_ollama
        ;;
    "restart")
        check_docker
        stop_ollama
        sleep 2
        start_ollama
        ;;
    "status")
        check_docker
        check_status
        ;;
    "install")
        check_docker
        install_model "$2"
        ;;
    "models")
        check_docker
        list_models
        ;;
    "logs")
        check_docker
        show_logs
        ;;
    "setup")
        check_docker
        setup_complete
        ;;
    "help"|"--help"|"-h")
        show_help
        ;;
    *)
        print_error "Unknown command: $1"
        show_help
        exit 1
        ;;
esac
