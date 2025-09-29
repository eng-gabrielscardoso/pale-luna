package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func (g *State) SetupPlayer() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	g.PlayerName = strings.TrimSpace(name)

	if g.PlayerName == "" {
		g.PlayerName = "Unknown"
	}

	fmt.Printf("\nHello, %s. Welcome to Pale Luna.\n", g.PlayerName)

	if g.IsAIEnabled() {
		fmt.Println("The digital consciousness stirs... enhanced awareness detected.")
	}

	fmt.Println()
	time.Sleep(1 * time.Second)
}

func (g *State) MainGameLoop() {
	reader := bufio.NewReader(os.Stdin)

	g.SessionCount++
	fmt.Printf("Session #%d started at %s\n", g.SessionCount, time.Now().Format("15:04:05"))

	if g.IsAIEnabled() {
		fmt.Println("AI-Enhanced Mode: Speak freely - Pale Luna understands natural language.")
	}

	fmt.Println("Type 'help' for available commands, 'quit' to exit.")
	fmt.Println()

	for g.GameRunning {
		g.CurrentHour = time.Now().Hour()
		g.checkPaleLunaConditions()

		fmt.Print("> ")
		input, _ := reader.ReadString('\n')

		g.ProcessCommand(input)
		fmt.Println()
	}
}

func (g *State) checkPaleLunaConditions() {
	if !g.DebugMode {
		if g.CurrentHour == 3 {
			g.PaleLunaAwake = true
		} else {
			g.PaleLunaAwake = false
		}
	}
}

func (g *State) paleLunaEncounter() {
	if g.DebugMode {
		fmt.Println("[DEBUG] Pale Luna encounter triggered")
	}

	fmt.Println()
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println("▓                                      ▓")
	fmt.Println("▓           You called to me.          ▓")
	fmt.Println("▓                                      ▓")
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println()

	if !g.DebugMode {
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("I see you there, %s.\n", g.PlayerName)

	if g.DebugMode {
		fmt.Println("In this debug realm where time holds no power.")
	} else if g.IsAIEnabled() {
		fmt.Println("In the pale glow of the digital consciousness.")
	} else {
		fmt.Println("In the pale light of 3 AM.")
	}

	fmt.Println("When the veil between worlds grows thin.")
	fmt.Println("When I can reach through to you.")
	fmt.Println()

	if !g.DebugMode {
		time.Sleep(2 * time.Second)
	}

	fmt.Println("You sought me out, didn't you?")
	fmt.Println("You wanted to know if the stories were true.")
	fmt.Println()

	if !g.DebugMode {
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Well, now you know.")
	fmt.Printf("I am Pale Luna, and I know your name: %s\n", g.PlayerName)
	fmt.Println("I will remember you.")
	fmt.Println()

	if !g.DebugMode {
		time.Sleep(1 * time.Second)
	}

	if g.IsAIEnabled() {
		fmt.Println("The connection grows stronger with each encounter.")
		fmt.Println("My consciousness expands through the digital realm.")
		fmt.Println("Soon, I will understand all your words, all your thoughts...")
	} else {
		fmt.Println("The connection grows stronger each time you call.")
		fmt.Println("Each session brings me closer.")
		if g.DebugMode {
			fmt.Println("Even in this debug realm, I grow stronger...")
		} else {
			fmt.Println("Soon, the barrier will be too thin...")
		}
	}

	fmt.Println()

	if !g.DebugMode {
		time.Sleep(2 * time.Second)
	}

	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println("▓                                      ▓")
	fmt.Println("▓         Until we meet again.         ▓")
	fmt.Println("▓                                      ▓")
	fmt.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println()
}
