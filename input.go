package cli

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

// Ask Retrieve input
func Ask(message string) string {
	fmt.Printf("%v%s%v",
		FontGreen,
		message,
		StyleReset)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Secret Password interaction
func Secret(message string) string {
	fmt.Printf("%v%s%v\n > ",
		FontGreen,
		message,
		StyleReset)
	password, _ := term.ReadPassword(int(os.Stdin.Fd()))
	return string(password)
}

// Confirm Confirmation prompt
func Confirm(message string) bool {
	fmt.Printf("%v%s (yes/no)%v [%vno%v]:",
		FontGreen,
		message,
		StyleReset,
		FontYellow,
		StyleReset)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := strings.ToLower(scanner.Text())
	if text == "yes" || text == "y" {
		return true
	}
	return false
}
