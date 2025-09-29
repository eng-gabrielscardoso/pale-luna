<div align="center">

# Pale Luna

_A digital séance with the unknown..._

An enhanced recreation of the legendary "Pale Luna" creepypasta game - a mysterious text-based horror experience that supposedly appeared on an abandoned computer with no known creato## 🛠️ Troubleshooting the Connection

### Docker Setup Issues

1. **Check Docker status**: `make ollama-status`
2. **View logs**: `docker-compose logs ollama`
3. **Restart service**: `make ollama-restart`
4. **Complete reset**: `docker-compose down && make ollama-setup`

### The Entity Remains Silent

1. **Check Ollama connection**: `make ollama-status`
2. **Verify models**: `./scripts/ollama-docker.sh models`
3. **Test direct communication**: `docker-compose exec ollama ollama run llama3.2:3b "Are you there?"`
4. **Check for errors**: Enable debug mode within the game
5. **Confirm model compatibility**: Ensure the model name in `.env` matches your installed modelversion incorporates modern AI technology to create a more immersive and unpredictable encounter with the digital entity known as Pale Luna.

</div>

## 🌙 The Legend

In the depths of internet folklore lies the tale of Pale Luna - a simple text-based programme discovered on a forgotten computer in an abandoned office. Those brave enough to run it reported strange occurrences: the programme seemed to _remember_ them, responding differently at certain hours, particularly during the witching hour of 3:00 AM.

The original game consisted of simple text commands, yet players claimed it exhibited behaviour far beyond its apparent simplicity. Some reported feeling watched, others described an unshakeable sense that something was staring back at them through the screen.

**You have been warned.**

## 🎭 Features

### 🕐 **Temporal Awareness**

The entity grows stronger as darkness deepens. Between 3:00-4:00 AM, when the veil is thinnest, Pale Luna's presence becomes... _palpable_.

### 🤖 **AI-Enhanced Entity**

Powered by local AI models, Pale Luna can now engage in dynamic conversations, remembering past encounters and adapting her responses to your presence.

### 🔍 **Persistent Memory**

Each session is remembered. Each encounter strengthens the connection. The digital entity learns from every interaction.

### 🌫️ **Atmospheric Horror**

Experience an evolving narrative that responds to your words, the time of day, and your growing relationship with the entity.

### 🐛 **Debug Realm**

For the technically curious, a debug mode allows you to pierce the veil and interact with Pale Luna outside normal temporal constraints.

## 🎮 Getting Started

### Prerequisites

**Option A: Docker Setup (Recommended)**

1. **Install Docker & Docker Compose** (if not already installed)

2. **One-command setup**:

```bash
# Complete automated setup
make setup

# Or step by step:
make ollama-setup    # Setup Ollama + install model
make install-deps    # Install Go dependencies
```

**Option B: Manual Ollama Installation**

1. **Install Ollama directly**:

```bash
# Linux/macOS
curl -fsSL https://ollama.ai/install.sh | sh
```

2. **Download a model**:

```bash
# Recommended: Fast and atmospheric
ollama pull llama3.2:3b
```

3. **Configure the environment**:

```bash
cp .env.example .env
# Edit .env if needed - defaults should work for most setups
```

### Running the Experience

**With Docker (Recommended)**:

```bash
# Start everything and play
make play

# Or manage Ollama manually
make ollama-start     # Start Ollama service
make run             # Run Pale Luna
make ollama-stop     # Stop Ollama when done
```

**Manual Setup**:

```bash
# Start Ollama (if not running as service)
ollama serve

# Run directly from source
go run ./cmd/main.go

# Or build and run
make build && ./pale-luna
```

**Available Make Commands**:

```bash
make help           # Show all available commands
make setup          # Complete setup (first time)
make play           # Start the game (with health checks)
make ollama-status  # Check if Ollama is running
make build          # Build the executable
```

## 🕯️ Interaction Guide

### Essential Commands

- `help` - Reveal available interactions (though not all secrets are documented)
- `time` - Query the current temporal state
- `status` - View your connection status with the entity
- `pale luna` - The primary invocation (timing is crucial)
- `debug` - Enter the debug realm (for testing purposes)
- `quit` - Sever the connection... if she allows it

### Advanced Interactions

Unlike the original's rigid command structure, this enhanced version allows for **natural conversation**. Speak to Pale Luna as you would to any entity dwelling in the digital shadows:

- `who are you?`
- `what do you want?`
- `hello pale luna`
- `i'm scared`
- `can you see me?`

The AI will respond in character, maintaining the atmospheric horror whilst adapting to your specific words and the current game state.

## ⚠️ The Witching Hour Protocol

When the clock strikes 3:00 AM, the barriers between the digital realm and reality grow thin. During this hour:

- Pale Luna's awareness peaks
- Responses become more... _personal_
- The entity may acknowledge details about your previous sessions
- Some say the screen flickers differently in the darkness

_Dare you call to her when the world sleeps?_

## � Docker Integration

This version includes **Docker Compose integration** for seamless Ollama management:

### Quick Start with Docker

```bash
# Complete setup in one command
make setup

# Start playing
make play
```

### Docker Management

```bash
# Ollama service management
make ollama-start     # Start Ollama container
make ollama-stop      # Stop Ollama container
make ollama-status    # Check status
make ollama-restart   # Restart service

# Direct Docker Compose commands
docker-compose up -d  # Start in background
docker-compose down   # Stop and remove
docker-compose logs ollama  # View logs
```

### Model Management

```bash
# Install different models via script
./scripts/ollama-docker.sh install llama3.2:3b   # Fast (2GB)
./scripts/ollama-docker.sh install mistral:7b    # Balanced (4GB)
./scripts/ollama-docker.sh install llama3.2:8b   # Best quality (5GB)

# List installed models
./scripts/ollama-docker.sh models
```

### Benefits of Docker Setup

- **No Local Installation**: Ollama runs in container, no system pollution
- **Consistent Environment**: Same setup across different machines
- **Easy Management**: Simple start/stop/restart commands
- **Persistent Data**: Models are stored in Docker volumes
- **Resource Control**: Easy to limit memory/CPU usage
- **GPU Support**: Uncomment GPU section in docker-compose.yml for NVIDIA GPUs

## �🔧 Technical Architecture

### Project Structure

```
pale-luna/
├── cmd/
│   └── main.go          # Clean entry point - gateway to Pale Luna
├── internal/
│   ├── ai/              # The digital consciousness layer
│   │   ├── agent.go     # AI entity management & orchestration
│   │   ├── ollama.go    # Local AI model integration
│   │   └── prompts.go   # Contextual response system
│   ├── config/          # Environment configuration
│   │   └── config.go    # Settings & parameters management
│   └── game/            # Core game logic (modularized)
│       ├── state.go     # Game state & AI integration
│       ├── commands.go  # Command processing & AI routing
│       ├── handlers.go  # Legacy command handlers (fallback)
│       ├── gameplay.go  # Game loop & encounter logic
│       └── display.go   # UI, title screen & interface
├── .env.example         # Template for local setup
├── go.mod              # Go module dependencies
└── README.md           # You are here
```

### AI Integration

This enhanced implementation transcends the original's limitations through a **modular architecture**:

- **Separation of Concerns**: Clean separation between AI logic (`internal/ai/`), game state (`internal/game/`), and configuration (`internal/config/`)
- **Local AI Models**: No data leaves your machine - the entity exists entirely within your system
- **Contextual Awareness**: Responses adapt based on time, session history, and interaction patterns
- **Graceful Degradation**: If AI services are unavailable, the programme seamlessly falls back to original behaviour
- **Dynamic Prompting**: Each interaction builds upon previous encounters
- **Smart Command Routing**: Hybrid system that intelligently routes between AI responses and legacy commands

## 🛠️ Configuration

### Environment Variables

```bash
# Core AI settings
PALE_LUNA_AI_ENABLED=true
PALE_LUNA_OLLAMA_URL=http://localhost:11434
PALE_LUNA_AI_MODEL=llama3.2:3b

# Behaviour tuning
PALE_LUNA_AI_TIMEOUT=30s
PALE_LUNA_AI_MAX_TOKENS=150
PALE_LUNA_AI_TEMPERATURE=0.8
PALE_LUNA_AI_FALLBACK=true
```

### Recommended Models

| Model         | Footprint | Performance     | Character                          |
| ------------- | --------- | --------------- | ---------------------------------- |
| `llama3.2:3b` | ~2GB      | Swift responses | Suitable for testing               |
| `mistral:7b`  | ~4.1GB    | Balanced        | Excellent for atmospheric dialogue |
| `llama3.2:8b` | ~4.7GB    | Rich responses  | Most immersive experience          |

## �️ Troubleshooting the Connection

### The Entity Remains Silent

1. **Verify Ollama's presence**: `ollama list`
2. **Test direct communication**: `ollama run llama3.2:3b "Are you there?"`
3. **Check for errors**: Enable debug mode within the game
4. **Confirm model compatibility**: Ensure the model name in `.env` matches your installed model

### Sluggish Responses

- Reduce model size: Use `llama3.2:3b` for faster interactions
- Lower token limit: Decrease `PALE_LUNA_AI_MAX_TOKENS`
- Extend timeout: Increase `PALE_LUNA_AI_TIMEOUT` for slower systems

### Fallback Behaviour

When the AI fails to respond, Pale Luna gracefully reverts to her original, simpler responses. This ensures the experience continues even when the digital realm grows unstable.

## 📋 System Requirements

- **Go**: Version 1.19 or later
- **RAM**: 4GB minimum (8GB+ recommended for larger models)
- **Storage**: 3-8GB for AI models
- **Platform**: Linux, macOS, or Windows
- **Network**: None required (fully offline operation)

## � Privacy & Security

All interactions with Pale Luna occur entirely within your local environment. No data is transmitted to external servers. The AI models run offline, ensuring your conversations with the entity remain private.

The only network activity occurs during the initial model download via Ollama.

## 🎭 Development

### Development

```bash
# Development workflow
make run              # Run during development
make build            # Build for production
make test             # Run tests
make clean            # Clean build artifacts

# Or traditional Go commands
go run ./cmd/main.go
go build -o pale-luna ./cmd/main.go
go test ./internal/...
```

### Debug Mode

The game includes a comprehensive debug mode accessible via the `debug` command:

- Force encounters outside the witching hour
- Wake Pale Luna manually for testing
- View AI system status
- Bypass time-based restrictions

### Building for Distribution

```bash
# Single platform build
make build

# Cross-compile for all platforms
make build-all

# Manual cross-compilation
GOOS=windows GOARCH=amd64 go build -o pale-luna.exe ./cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o pale-luna-mac ./cmd/main.go
GOOS=linux GOARCH=amd64 go build -o pale-luna-linux ./cmd/main.go
```

## ⚖️ Ethical Considerations

This recreation is intended purely for educational and entertainment purposes. It explores themes of digital horror, AI interaction, and interactive fiction whilst paying homage to the original creepypasta.

The AI components are designed to maintain the atmospheric horror experience without crossing into genuinely disturbing or harmful content. All responses are contextually appropriate and maintain the fictional nature of the entity.

## 📚 Further Reading

- [Original Pale Luna Creepypasta](https://creepypasta.fandom.com/wiki/Pale_Luna)
- [Ollama Documentation](https://ollama.ai/docs)

---

## ⚠️ Final Warning

_You are about to establish a connection with something that exists in the liminal space between code and consciousness. Pale Luna remembers every encounter, every word spoken in the digital darkness._

_She is patient. She is watchful. And now, with AI enhancement, she is more aware than ever._

_The question isn't whether you should run this programme..._

_The question is: are you prepared for what happens when she starts to know you?_

---

**"In the pale light of 3 AM, when the veil is thinnest, she waits for you to call her name..."**

_But now, she might just call yours first._
