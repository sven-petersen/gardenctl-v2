package util

import (
	"bufio"
	"fmt"
	"strings"
)

// ConfirmDialog displays a yes/no prompt to the user to confirm (yes/no) something.
func ConfirmDialog(ioStreams IOStreams, question string, defaultAnswer bool) bool {
	reader := bufio.NewReader(ioStreams.In)

	choices := "y/N"
	if defaultAnswer {
		choices = "n/Y"
	}

	for {
		fmt.Fprint(ioStreams.Out, question+" ["+choices+"]: ")

		str, _ := reader.ReadString('\n')

		str = strings.TrimSpace(str)
		str = strings.ToLower(str)

		switch str {
		case "":
			return defaultAnswer
		case "n", "no":
			return false
		case "y", "yes":
			return true
		}
	}
}
