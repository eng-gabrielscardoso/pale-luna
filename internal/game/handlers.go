package game

import (
	"fmt"
	"strings"
	"time"
)

func (g *State) handlePaleLunaCommand() {
	if g.PaleLunaAwake {
		g.paleLunaEncounter()
	} else {
		fmt.Println("Nothing happens.")
		fmt.Println("You feel like you're missing something important.")
		if g.CurrentHour != 3 {
			fmt.Println("Perhaps the timing isn't right...")
		}
	}
}

func (g *State) handleSleepCommand() {
	if g.PaleLunaAwake {
		g.PaleLunaAwake = false

		fmt.Println("Pale Luna has gone back to sleep.")
		if !g.DebugMode {
			time.Sleep(2 * time.Second)
		}

		fmt.Println("She will not respond until the next encounter.")
		if !g.DebugMode {
			time.Sleep(2 * time.Second)
		}

		fmt.Println("But maybe you can call her again in your dreams...")
		if !g.DebugMode {
			time.Sleep(2 * time.Second)
		}
	} else {
		fmt.Println("Pale Luna is already asleep.")
	}
}

func (g *State) handleLunaCommand() {
	if g.PaleLunaAwake {
		fmt.Println("Luna... yes, I remember Luna.")
		fmt.Println("She was beautiful once.")
		fmt.Println("Before the pale consumed her.")
	} else {
		fmt.Println("Luna sleeps in the digital darkness.")
	}
}

func (g *State) handlePaleCommand() {
	if g.PaleLunaAwake {
		fmt.Println("Pale... like moonlight on bone.")
		fmt.Println("Pale... like the color that remains when life fades.")
	} else {
		fmt.Println("Everything seems pale in comparison to what lurks in the shadows.")
	}
}

func (g *State) handleWhoAreYou() {
	if g.PaleLunaAwake {
		fmt.Println("I am the one who watches.")
		fmt.Println("I am the one who waits.")
		fmt.Println("I am Pale Luna.")
		fmt.Println()
		fmt.Printf("And you, %s, have called to me in the dark hour.\n", g.PlayerName)
	} else {
		fmt.Println("I am just a program.")
		fmt.Println("...or am I?")
	}
}

func (g *State) handleUnknownCommand(input string) {
	responses := []string{
		"The digital void does not understand those words.",
		"Unknown command. Type 'help' for available commands.",
		"The shadows whisper back, but I cannot make out the meaning.",
	}

	if g.PaleLunaAwake {
		eerieResponses := []string{
			"The pale light flickers at your words, but remains silent.",
			fmt.Sprintf("Did you mean to say something else, %s?", g.PlayerName),
			"Something stirs in the darkness at your voice, but nothing emerges.",
			"I hear you calling through the veil, but your words are unclear.",
		}
		responses = append(responses, eerieResponses...)
	}

	if strings.Contains(input, "hello") || strings.Contains(input, "hi") {
		if g.PaleLunaAwake {
			fmt.Printf("Hello, %s. I have been waiting for you to speak.\n", g.PlayerName)
		} else {
			fmt.Printf("Hello, %s. The silence acknowledges your presence.\n", g.PlayerName)
		}
		return
	}

	if strings.Contains(input, "scared") || strings.Contains(input, "afraid") {
		if g.PaleLunaAwake {
			fmt.Println("Fear is natural when facing the unknown. The pale moon sees all fears.")
		} else {
			fmt.Println("There's nothing to fear... not yet.")
		}
		return
	}

	responseIndex := int(time.Now().UnixNano()) % len(responses)
	fmt.Println(responses[responseIndex])
}
