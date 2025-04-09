package challenge

import (
	"fmt"
	"linux-learner/commands"
	"strings"
)

func ValidateCommand(userInput, expectedCommand, validationType string) bool {
	userResult := commands.RunCommand(userInput)

	switch validationType {
	case "exact_match":
		normalizedInput := normalizeCommand(userInput)
		normalizedExpected := normalizeCommand(expectedCommand)
		if normalizedInput == normalizedExpected {
			return true
		}

	case "output_match":
		expectedResult := commands.RunCommand(expectedCommand)
		if userResult.Err == nil && expectedResult.Err == nil &&
			userResult.Output == expectedResult.Output {
			return true
		}
	}

	// feedback when it fails
	fmt.Println("Incorrect command.")
	if userResult.Err != nil {
		fmt.Printf("Shell error:\n%s\n", userResult.Stderr)
	}

	suggestion := suggestFix(userInput)
	if suggestion != "" {
		fmt.Printf("Suggestion: %s\n", suggestion)
	}

	return false
}

func suggestFix(cmd string) string {
	if strings.Count(cmd, "'")%2 != 0 ||
		strings.Count(cmd, "\"")%2 != 0 {
		return "Check your quotes — looks like one might be missing or unmatched."
	}
	if strings.Contains(cmd, "ecoh") {
		return "Did you mean `echo`?"
	}
	if strings.Contains(cmd, "cd .. /") {
		return "Try removing the space between `..` and `/`."
	}
	return ""
}

func normalizeCommand(cmd string) string {
	// replace all double quotes with single
	return strings.ReplaceAll(cmd, "\"", "'")
}
