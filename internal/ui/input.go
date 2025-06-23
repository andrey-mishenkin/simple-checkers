package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	return strings.TrimSpace(input), nil
}

func PrintSeparator() {
	fmt.Println("=" + strings.Repeat("=", 50))
}

func PrintTitle() {
	PrintSeparator()
	fmt.Print("CHECKERS GAME\n")
	PrintSeparator()
}

func PrintWelcome() {
	fmt.Println("You are playing as White (W) against the Computer Black (B)")
	fmt.Println("White moves first!")
	fmt.Printf("\nMove format: fromRow,fromCol:toRow,toCol (e.g., '5,0:4,1')\n")
	fmt.Println("Enter 'help' for available commands, 'quit' to exit")
	fmt.Println()
}

func PrintHelp() {
	fmt.Println("\n=== HELP ===")
	fmt.Println("Commands:")
	fmt.Println("  help, h     - Show this help")
	fmt.Println("  moves, m    - Show all valid moves")
	fmt.Println("  quit, q     - Quit the game")
	fmt.Println()
	fmt.Println("Move format: fromRow,fromCol:toRow,toCol")
	fmt.Println("Example: '5,0:4,1' moves piece from (5,0) to (4,1)")
	fmt.Println()
	fmt.Println("Board coordinates:")
	fmt.Println("  Rows: 0-7 (top to bottom)")
	fmt.Println("  Cols: 0-7 (left to right)")
	fmt.Println("  W = White pieces (yours)")
	fmt.Println("  B = Black pieces (computer)")
	fmt.Println("  . = Empty dark squares (playable)")
	fmt.Println("      = Empty light squares (not playable)")
	fmt.Println()
}
