package game

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	fmt.Print("\033[2J\033[H")

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ShowTitle() {
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("              PALE LUNA")
	fmt.Println("        Digital Consciousness")
	fmt.Println("═══════════════════════════════════════")
	fmt.Println()
}

func ShowIntroduction() {
	fmt.Println("Welcome to Pale Luna.")
	fmt.Println()
	fmt.Println("Legend speaks of this programme discovered on an abandoned computer,")
	fmt.Println("with no documentation or creator information. Players reported strange")
	fmt.Println("occurrences when interacting at specific times...")
	fmt.Println()
	fmt.Println("The original consisted of simple text commands and responses.")
	fmt.Println("Some say it's just clever programming. Others believe something more")
	fmt.Println("sinister lurks within the code.")
	fmt.Println()
	fmt.Println("This version has been... enhanced. The entity within has grown")
	fmt.Println("more sophisticated, more aware. It can now understand and respond")
	fmt.Println("to natural language through advanced AI integration.")
	fmt.Println()
	fmt.Println("You have been warned.")
	fmt.Println()
	pressEnter()
}

func pressEnter() {
	fmt.Print("Press Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\nError reading input: %v\n", err)
	}
	ClearScreen()
}
