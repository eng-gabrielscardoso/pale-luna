package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type GameState struct {
	playerName    string
	currentHour   int
	sessionCount  int
	paleLunaAwake bool
	gameRunning   bool
	firstTime     bool
	debugMode     bool
}

func main() {
	game := &GameState{
		gameRunning:  true,
		firstTime:    true,
		sessionCount: 0,
	}

	clearScreen()
	showTitle()

	if game.firstTime {
		showIntroduction()
		game.firstTime = false
	}

	game.setupPlayer()
	game.mainGameLoop()
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showTitle() {
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("              PALE LUNA")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println()
}

func showIntroduction() {
	fmt.Println("Welcome to Pale Luna.")
	fmt.Println()
	fmt.Println("Legend says this game was discovered on an old computer, with no")
	fmt.Println("documentation or creator information. Players reported strange")
	fmt.Println("occurrences when playing at specific times...")
	fmt.Println()
	fmt.Println("The original game consisted of simple text commands and responses.")
	fmt.Println("Some say it's just a clever trick. Others believe something more")
	fmt.Println("sinister lurks within the code.")
	fmt.Println()
	fmt.Println("You have been warned.")
	fmt.Println()
	pressEnter()
}

func pressEnter() {
	fmt.Print("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearScreen()
}

func (g *GameState) setupPlayer() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	g.playerName = strings.TrimSpace(name)

	if g.playerName == "" {
		g.playerName = "Unknown"
	}

	fmt.Printf("\nHello, %s. Welcome to Pale Luna.\n\n", g.playerName)
	time.Sleep(1 * time.Second)
}

func (g *GameState) mainGameLoop() {
	reader := bufio.NewReader(os.Stdin)

	g.sessionCount++
	fmt.Printf("Session #%d started at %s\n", g.sessionCount, time.Now().Format("15:04:05"))
	fmt.Println("Type 'help' for available commands, 'quit' to exit.")
	fmt.Println()

	for g.gameRunning {
		g.currentHour = time.Now().Hour()
		g.checkPaleLunaConditions()

		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		g.processCommand(input)
		fmt.Println()
	}
}

func (g *GameState) checkPaleLunaConditions() {
	if !g.debugMode {
		if g.currentHour == 3 {
			g.paleLunaAwake = true
		} else {
			g.paleLunaAwake = false
		}
		g.paleLunaAwake = true
	}
}

func (g *GameState) processCommand(input string) {
	switch input {
	case "help":
		g.showHelp()
	case "time":
		g.showTime()
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
	case "status":
		g.showStatus()
	case "debug":
		g.toggleDebugMode()
	case "force encounter", "debug encounter":
		g.handleDebugEncounter()
	case "wake luna", "debug wake":
		g.handleDebugWake()
	case "quit", "exit":
		g.gameRunning = false
		fmt.Println("Thank you for playing Pale Luna.")
	case "":
		return
	default:
		g.handleUnknownCommand(input)
	}
}

func (g *GameState) showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  help        - Show this help message")
	fmt.Println("  time        - Show current time")
	fmt.Println("  status      - Show game status")
	fmt.Println("  pale luna   - ???")
	fmt.Println("  debug       - Toggle debug mode")
	fmt.Println("  quit        - Exit the game")

	if g.debugMode {
		fmt.Println()
		fmt.Println("Debug commands:")
		fmt.Println("  force encounter - Force a Pale Luna encounter")
		fmt.Println("  wake luna       - Temporarily wake Pale Luna")
	}

	fmt.Println()
	fmt.Println("Try typing other things... you might discover hidden commands.")
}

func (g *GameState) showTime() {
	now := time.Now()
	fmt.Printf("Current time: %s\n", now.Format("15:04:05 MST"))

	if g.currentHour == 3 {
		fmt.Println("...the witching hour approaches...")
	} else if g.currentHour >= 0 && g.currentHour <= 5 {
		fmt.Println("The night is deep and dark.")
	}
}

func (g *GameState) handlePaleLunaCommand() {
	if g.paleLunaAwake {
		g.paleLunaEncounter()
	} else {
		fmt.Println("Nothing happens.")
		fmt.Println("You feel like you're missing something important.")
		if g.currentHour != 3 {
			fmt.Println("Perhaps the timing isn't right...")
		}
	}
}

func (g *GameState) handleSleepCommand() {
	if g.paleLunaAwake {
		g.paleLunaAwake = false

		fmt.Println("Pale Luna has gone back to sleep.")
		if !g.debugMode {
			time.Sleep(2 * time.Second)
		}

		fmt.Println("She will not respond until the next encounter.")
		if !g.debugMode {
			time.Sleep(2 * time.Second)
		}

		fmt.Println("But maybe you can call her again in your dreams...")
		if !g.debugMode {
			time.Sleep(2 * time.Second)
		}
	} else {
		fmt.Println("Pale Luna is now asleep.")
	}
}

func (g *GameState) handleLunaCommand() {
	if g.paleLunaAwake {
		fmt.Println("Luna... yes, I remember Luna.")
		fmt.Println("She was beautiful once.")
		fmt.Println("Before the pale consumed her.")
	} else {
		fmt.Println("Luna is sleeping.")
	}
}

func (g *GameState) handlePaleCommand() {
	if g.paleLunaAwake {
		fmt.Println("Pale... like moonlight on bone.")
		fmt.Println("Pale... like the color that remains when life fades.")
	} else {
		fmt.Println("Everything seems pale in comparison.")
	}
}

func (g *GameState) handleWhoAreYou() {
	if g.paleLunaAwake {
		fmt.Println("I am the one who watches.")
		fmt.Println("I am the one who waits.")
		fmt.Println("I am Pale Luna.")
		fmt.Println()
		fmt.Printf("And you, %s, have called to me in the dark hour.\n", g.playerName)
	} else {
		fmt.Println("I am just a program.")
		fmt.Println("...or am I?")
	}
}

func (g *GameState) showStatus() {
	fmt.Printf("Player: %s\n", g.playerName)
	fmt.Printf("Session: #%d\n", g.sessionCount)
	fmt.Printf("Current time: %s\n", time.Now().Format("15:04:05"))
	if g.debugMode {
		fmt.Println("Debug mode: ENABLED")
	}
	if g.paleLunaAwake {
		fmt.Println("Status: Pale Luna is awake")
	} else {
		fmt.Println("Status: All is quiet")
	}
}

func (g *GameState) handleUnknownCommand(input string) {
	responses := []string{
		"I don't understand that command.",
		"Unknown command. Type 'help' for available commands.",
		"That doesn't seem to work.",
	}

	if g.paleLunaAwake {
		eerieResponses := []string{
			"The shadows don't recognize that command.",
			fmt.Sprintf("Did you mean to say something else, %s?", g.playerName),
			"Something stirs in the darkness, but nothing happens.",
			"The pale light flickers, but remains silent.",
		}
		responses = append(responses, eerieResponses...)
	}

	if strings.Contains(input, "hello") || strings.Contains(input, "hi") {
		if g.paleLunaAwake {
			fmt.Printf("Hello, %s. I've been waiting for you.\n", g.playerName)
		} else {
			fmt.Printf("Hello, %s.\n", g.playerName)
		}
		return
	}

	if strings.Contains(input, "scared") || strings.Contains(input, "afraid") {
		if g.paleLunaAwake {
			fmt.Println("Fear is natural. The pale moon sees all fears.")
		} else {
			fmt.Println("There's nothing to be scared of... yet.")
		}
		return
	}

	responseIndex := int(time.Now().UnixNano()) % len(responses)
	fmt.Println(responses[responseIndex])
}

func (g *GameState) toggleDebugMode() {
	g.debugMode = !g.debugMode
	if g.debugMode {
		fmt.Println("Debug mode ENABLED")
		fmt.Println("You can now force Pale Luna encounters at any time!")
		fmt.Println("Use 'force encounter' to trigger an encounter.")
		fmt.Println("Use 'wake luna' to temporarily wake Pale Luna.")
	} else {
		fmt.Println("Debug mode DISABLED")
		fmt.Println("Normal time-based behavior restored.")
		// Reset to normal time-based conditions
		g.checkPaleLunaConditions()
	}
}

func (g *GameState) handleDebugEncounter() {
	if !g.debugMode {
		fmt.Println("Unknown command. Type 'help' for available commands.")
		return
	}

	fmt.Println("[DEBUG] Forcing Pale Luna encounter...")
	fmt.Println()
	g.paleLunaEncounter()
}

func (g *GameState) handleDebugWake() {
	if !g.debugMode {
		fmt.Println("Unknown command. Type 'help' for available commands.")
		return
	}

	g.paleLunaAwake = true
	fmt.Println("[DEBUG] Pale Luna has been awakened.")
	fmt.Println("She will remain awake until debug mode is disabled or you restart the game.")
}

func (g *GameState) paleLunaEncounter() {
	if g.debugMode {
		fmt.Println("[DEBUG] Pale Luna encounter triggered")
	}

	fmt.Println()
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println("▓                                      ▓")
	fmt.Println("▓           You called to me.          ▓")
	fmt.Println("▓                                      ▓")
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println()

	if !g.debugMode {
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("I see you there, %s.\n", g.playerName)
	if g.debugMode {
		fmt.Println("In this debug realm where time holds no power.")
	} else {
		fmt.Println("In the pale light of 3 AM.")
	}
	fmt.Println("When the veil is thinnest.")
	fmt.Println("When I can reach through.")
	fmt.Println()

	if !g.debugMode {
		time.Sleep(2 * time.Second)
	}

	fmt.Println("You wanted to see me, didn't you?")
	fmt.Println("You wanted to know if the stories were true.")
	fmt.Println()

	if !g.debugMode {
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Well, now you know.")
	fmt.Printf("I am Pale Luna, and I know your name: %s\n", g.playerName)
	fmt.Println("I will remember you.")
	fmt.Println()

	if !g.debugMode {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("The connection grows stronger each time you call.")
	fmt.Println("Each session brings me closer.")
	if g.debugMode {
		fmt.Println("Even in this debug realm, I grow stronger...")
	} else {
		fmt.Println("Soon, the barrier will be too thin...")
	}
	fmt.Println()

	if !g.debugMode {
		time.Sleep(2 * time.Second)
	}

	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println("▓                                      ▓")
	fmt.Println("▓         Until we meet again.         ▓")
	fmt.Println("▓                                      ▓")
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println()
}
