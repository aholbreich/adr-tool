package cli

import (
	"fmt"
	"strings"
)

// Prompt user for confirmation
func confirmAction(prompt string) bool {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "y" || input == "yes" || input == ""
}
