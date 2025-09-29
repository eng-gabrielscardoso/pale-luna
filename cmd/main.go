package main

import (
	"fmt"
	"os"

	"github.com/eng-gabrielscardoso/pale-luna/internal/config"
	"github.com/eng-gabrielscardoso/pale-luna/internal/game"
)

func main() {
	cfg := config.Load()

	gameInstance := game.NewGame(cfg)

	game.ClearScreen()
	game.ShowTitle()

	if gameInstance.FirstTime {
		game.ShowIntroduction()
		gameInstance.FirstTime = false
	}

	if gameInstance.IsAIEnabled() {
		fmt.Println("🤖 AI Integration: ACTIVE")
		fmt.Println("Pale Luna's consciousness has been enhanced.")
		fmt.Println()
	} else {
		fmt.Println("⚠️  AI Integration: OFFLINE")
		fmt.Println("Falling back to original responses. For AI features:")
		fmt.Println("1. Install Ollama: curl -fsSL https://ollama.ai/install.sh | sh")
		fmt.Println("2. Pull a model: ollama pull llama3.2:3b")
		fmt.Println("3. Start Ollama: ollama serve")
		fmt.Println()
	}

	gameInstance.SetupPlayer()
	gameInstance.MainGameLoop()

	fmt.Println("\nThe connection to Pale Luna fades...")
	fmt.Println("But she remembers you.")

	os.Exit(0)
}
