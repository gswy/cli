package cli

import "fmt"

// Success Print success information
func Success(message string) {
	fmt.Printf("%v%s%v\n",
		FontGreen,
		message,
		StyleReset)
}

// Error Print error information
func Error(message string) {
	fmt.Printf("%v%s%v\n",
		FontRed,
		message,
		StyleReset)
}

// Warn Print warn information
func Warn(message string) {
	fmt.Printf("%v%s%v\n",
		FontYellow,
		message,
		StyleReset)
}
