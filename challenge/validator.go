package challenge

import (
	"fmt"
	"linux-learner/commands"
	"strings"
)

// ValidateCommand compares the user's command against the expected command,
// using the provided validation type ("exact_match" or "output_match").
// It also prints feedback and suggestions if the validation fails.
func ValidateCommand(
	userInput, expectedCommand, validationType string,
	feedback []InputFeedback,
) bool {
	userResult := commands.RunCommand(userInput)

	// Check for exact match or output match based on validation type
	switch validationType {
	case "exact_match":
		// Normalize both commands to ignore differences in quote styles
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

	// If validation fails, provide feedback based on input feedback
	printInputFeedback(userInput, feedback)

	// Provide shell error if there was one
	if userResult.Err != nil {
		fmt.Printf("Shell error:\n%s", userResult.Stderr)
	}

	return false
}

// printInputFeedback matches user input against predefined feedback patterns.
// It helps the user understand why a particular input is incorrect.
func printInputFeedback(userInput string, feedback []InputFeedback) {
	for _, fb := range feedback {
		// If the input matches a feedback pattern, print the corresponding response
		if strings.Contains(userInput, fb.Match) {
			fmt.Printf("Feedback: %s\n", fb.Response)
		}
	}
}

// normalizeCommand simplifies command comparison by standardizing quotes
func normalizeCommand(cmd string) string {
	// Replace all double quotes with single quotes
	return strings.ReplaceAll(cmd, "\"", "'")
}
