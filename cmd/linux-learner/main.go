package main

import (
	"bufio"
	"embed"
	"fmt"
	"linux-learner/challenge"
	"linux-learner/term"
	"os"
	"strings"
)

//go:embed data/*.yaml
var challengeFS embed.FS

func main() {
	// get challenge content
	chal, err := challenge.LoadChallenge(challengeFS, "data/nav_001.yaml")
	if err != nil {
		fmt.Println("Error loading challenge:", err)
		return
	}

	fmt.Println()

	// display challenge details
	displayChallengeHeader(chal)

	// initialise hint counter
	hintsUsed := 0

	// promt
	prompt := fmt.Sprintf(
		"\n%s@%s:~# ",
		term.ColorText("user", term.Colors.Green),
		term.ColorText("hesed-server", term.Colors.Red),
	)

	// main loop to get user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		userInput := strings.TrimSpace(scanner.Text())

		// Handle different commands: exit, hint, clear, help
		switch strings.ToLower(userInput) {
		case "exit":
			fmt.Print(
				term.ColorText(
					"Are you sure you want to exit the challenge? (y/N): ",
					term.Colors.Red,
				),
			)
			if !scanner.Scan() {
				fmt.Println("\nNo confirmation received. Staying in the challenge.")
				continue
			}

			confirm := strings.TrimSpace(scanner.Text())
			if strings.ToLower(confirm) == "y" {
				fmt.Println("Exiting challenge. Goodbye!")
				return
			}

			fmt.Println("Continuing challenge...")
			continue

		case "hint":
			if hintsUsed < len(chal.Hints) {
				fmt.Println("\n",
					term.Bold(
						term.ColorText(
							"[Hint]", term.Colors.Magenta,
						),
					),
					chal.Hints[hintsUsed].Text,
				)
				hintsUsed++
			} else {
				fmt.Println("\n[No more hints available]")
			}

			continue // no need to validate input

		case "clear":
			clearScreen()
			continue

		case "help":
			displayChallengeDetails(chal)
			displayPossibleCmds()
			continue
		}

		// validate user input
		valid := challenge.ValidateCommand(
			userInput, chal.ExpectedCommand, chal.ValidationType, chal.InputFeedback,
		)
		if valid {
			fmt.Println("\n",
				term.Bold(
					term.ColorText(
						chal.SuccessMessage,
						term.Colors.Green,
					),
				),
			)

			// Show reflection prompts after correct input

			fmt.Println(
				term.Bold(
					term.ColorText(
						"Reflection Questions:", term.Colors.Cyan),
				),
			)
			for _, q := range chal.Reflection {
				fmt.Println(" -", term.Italic(q.Question))
			}

			break // exit the main loop
		} else {
			fmt.Println(
				term.Bold(
					term.ColorText(
						"\nIncorrect. Please try again.", term.Colors.Red),
				),
			)
		}
	}
}

func displayChallengeDetails(chal *challenge.Challenge) {
	fmt.Println(
		term.BoldUnderline(
			"\nChallenge Details:",
		),
	)
	fmt.Println(
		term.Bold(
			term.ColorText(
				"Title:", term.Colors.Cyan,
			),
		),
		chal.Title,
	)
	fmt.Println(
		term.Bold(
			term.ColorText(
				"Description:", term.Colors.Cyan,
			),
		),
	)

	scenarioLines := formatTextWithLineBreaks(chal.Description, 60)
	for _, line := range scenarioLines {
		fmt.Println("  " + line) // Indent each line for better visual separation
	}

	fmt.Println(
		term.Bold(
			term.ColorText(
				"Objective:", term.Colors.Cyan,
			),
		),
		chal.Objective,
	)
	fmt.Println(
		term.Bold(
			term.ColorText(
				"Hints Available:", term.Colors.Yellow,
			),
		),
		fmt.Sprintf("%d", len(chal.Hints)),
	)

	fmt.Println(
		term.Bold(
			term.ColorText(
				"Reflection Questions:", term.Colors.Magenta,
			),
		),
		fmt.Sprintf("%d", len(chal.Reflection)),
	)
}

func displayChallengeHeader(chal *challenge.Challenge) {
	fmt.Println(
		term.BoldUnderline(
			term.ColorText(
				"Challenge:", term.Colors.Cyan,
			),
		),
		chal.Title,
	)
	fmt.Print(
		term.BoldUnderline(
			term.ColorText(
				"Scenario:", term.Colors.Blue,
			),
		),
	)

	// Format the scenario text with proper line breaks
	// Assuming a line width of 80 characters for better readability
	scenarioLines := formatTextWithLineBreaks(chal.Scenario, 60)
	for _, line := range scenarioLines {
		fmt.Println("  " + line) // Indent each line for better visual separation
	}

	fmt.Println(
		term.BoldUnderline(
			term.ColorText(
				"Objective:", term.Colors.Green,
			),
		),
		chal.Objective,
	)

	displayPossibleCmds()
}

func displayPossibleCmds() {
	// Predefined list of possible commands
	var possibleCommands = []string{
		"exit - Exit the challenge",
		"hint - Get a hint",
		"clear - Clear the screen",
		"help - Show challenge details",
	}

	fmt.Println(
		term.Bold(
			term.ColorText(
				"\nAvailable Commands:", term.Colors.Yellow,
			),
		),
	)
	for _, cmd := range possibleCommands {
		fmt.Println(
			"  ",
			term.Italic(
				term.ColorText(
					cmd, term.Colors.White,
				),
			),
		)
	}
}

// formatTextWithLineBreaks splits a long text into lines of specified width
func formatTextWithLineBreaks(text string, lineWidth int) []string {
	words := strings.Fields(text)
	lines := []string{}

	currentLine := ""
	for _, word := range words {
		// If adding this word exceeds the line width, start a new line
		if len(currentLine)+len(word)+1 > lineWidth && currentLine != "" {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
		}
	}

	// Add the last line if it's not empty
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

// clearScreen clears the terminal screen using ANSI escape codes
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
