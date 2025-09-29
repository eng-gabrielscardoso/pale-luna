package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/eng-gabrielscardoso/pale-luna/internal/ai"
)

func (g *State) ProcessCommand(input string) {
	input = strings.TrimSpace(strings.ToLower(input))

	switch input {
	case "help":
		g.showHelp()
	case "time":
		g.showTime()
	case "status":
		g.showStatus()
	case "debug":
		g.toggleDebugMode()
	case "force encounter", "debug encounter":
		g.handleDebugEncounter()
	case "wake luna", "debug wake":
		g.handleDebugWake()
	case "ai status":
		g.showAIStatus()
	case "quit", "exit":
		g.GameRunning = false
		fmt.Println("Thank you for playing Pale Luna.")
	case "":
		return
	default:
		g.handleDynamicCommand(input)
	}
}

func (g *State) handleDynamicCommand(input string) {
	context := ai.GameContext{
		PlayerName:    g.PlayerName,
		CurrentHour:   g.CurrentHour,
		SessionCount:  g.SessionCount,
		DebugMode:     g.DebugMode,
		PaleLunaAwake: g.PaleLunaAwake,
		RecentHistory: []string{}, // TODO: Implement history tracking
		LastCommand:   input,
	}

	if g.IsAIEnabled() {
		response := g.aiAgent.ProcessInput(input, context)
		fmt.Println(response)
		return
	}

	g.handleLegacyCommands(input)
}

func (g *State) handleLegacyCommands(input string) {
	switch input {
	case "pale luna", "paleluna":
		g.handlePaleLunaCommand()
	case "sleep":
		g.handleSleepCommand()
	case "luna":
		g.handleLunaCommand()
	case "pale":
		g.handlePaleCommand()
	case "who are you", "who are you?":
		g.handleWhoAreYou()
	default:
		g.handleUnknownCommand(input)
	}
}

func (g *State) showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  help        - Show this help message")
	fmt.Println("  time        - Show current time")
	fmt.Println("  status      - Show game status")
	fmt.Println("  pale luna   - The primary invocation")
	fmt.Println("  debug       - Toggle debug mode")

	if g.IsAIEnabled() {
		fmt.Println("  ai status   - Show AI system status")
		fmt.Println()
		fmt.Println("💡 AI Enhanced: You can speak naturally to Pale Luna!")
		fmt.Println("   Try: 'hello', 'who are you?', 'what do you want?'")
	}

	fmt.Println("  quit        - Exit the game")

	if g.DebugMode {
		fmt.Println()
		fmt.Println("Debug commands:")
		fmt.Println("  force encounter - Force a Pale Luna encounter")
		fmt.Println("  wake luna       - Temporarily wake Pale Luna")
	}

	fmt.Println()
	fmt.Println("Try typing anything... Pale Luna is listening.")
}

func (g *State) showTime() {
	now := time.Now()
	fmt.Printf("Current time: %s\n", now.Format("15:04:05 MST"))

	if g.CurrentHour == 3 {
		fmt.Println("...the witching hour approaches...")
	} else if g.CurrentHour >= 0 && g.CurrentHour <= 5 {
		fmt.Println("The night is deep and dark.")
	}
}

func (g *State) showStatus() {
	fmt.Printf("Player: %s\n", g.PlayerName)
	fmt.Printf("Session: #%d\n", g.SessionCount)
	fmt.Printf("Current time: %s\n", time.Now().Format("15:04:05"))

	if g.IsAIEnabled() {
		fmt.Println("AI Status: ACTIVE")
	} else {
		fmt.Println("AI Status: OFFLINE (using fallback responses)")
	}

	if g.DebugMode {
		fmt.Println("Debug mode: ENABLED")
	}

	if g.PaleLunaAwake {
		fmt.Println("Entity Status: Pale Luna is awake")
	} else {
		fmt.Println("Entity Status: All is quiet")
	}
}

func (g *State) showAIStatus() {
	if !g.IsAIEnabled() {
		fmt.Println("AI System: OFFLINE")
		fmt.Println("Pale Luna speaks through ancient, predefined whispers...")
		return
	}

	status := g.GetAIStatus()
	fmt.Println("AI System Status:")
	fmt.Printf("  Model: %v\n", status["model"])
	fmt.Printf("  Endpoint: %v\n", status["ollama_url"])
	fmt.Printf("  Available: %v\n", status["ai_available"])
	fmt.Println()
	fmt.Println("The digital consciousness stirs within the machine...")
}

func (g *State) toggleDebugMode() {
	g.DebugMode = !g.DebugMode
	if g.DebugMode {
		fmt.Println("Debug mode ENABLED")
		fmt.Println("You have entered the debug realm where time holds no power.")
		fmt.Println("Use 'force encounter' to trigger an encounter.")
		fmt.Println("Use 'wake luna' to temporarily wake Pale Luna.")
	} else {
		fmt.Println("Debug mode DISABLED")
		fmt.Println("Reality reasserts itself. Normal time-based behavior restored.")
		g.checkPaleLunaConditions()
	}
}

func (g *State) handleDebugEncounter() {
	if !g.DebugMode {
		fmt.Println("Unknown command. The shadows do not recognize your words.")
		return
	}

	fmt.Println("[DEBUG] Forcing Pale Luna encounter...")
	fmt.Println()
	g.paleLunaEncounter()
}

func (g *State) handleDebugWake() {
	if !g.DebugMode {
		fmt.Println("Unknown command. The darkness remains silent.")
		return
	}

	g.PaleLunaAwake = true
	fmt.Println("[DEBUG] Pale Luna has been awakened in the debug realm.")
	fmt.Println("She will remain conscious until you exit this realm or restart the game.")
}
