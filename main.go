package main

import (
	"bufio"
	"embed"
	"fmt"
	"linux-learner/challenge"
	"os"
	"strings"
)

//go:embed data/*.yaml
var challengeFS embed.FS

// Predefined list of possible commands
var possibleCommands = []string{
	"exit - Exit the challenge",
	"hint - Get a hint",
	"clear - Clear the screen",
	"help - Show challenge details",
}

func main() {
	// get challenge content
	chal, err := challenge.LoadChallenge(challengeFS, "data/nav_001.yaml")
	if err != nil {
		fmt.Println("Error loading challenge:", err)
		return
	}

	// display challenge details
	fmt.Println("Challenge:", chal.Title)
	fmt.Println("Scenario:", chal.Scenario)
	fmt.Println("Objective:", chal.Objective)
	fmt.Println("\nAvailable Commands:")
	for _, cmd := range possibleCommands {
		fmt.Println("  ", cmd)
	}

	// initialise hint counter
	hintsUsed := 0

	// main loop to get user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nuser@hesed-server:~# ")
		scanner.Scan()
		userInput := strings.TrimSpace(scanner.Text())

		// Handle different commands: exit, hint, clear, help
		switch strings.ToLower(userInput) {
		case "exit":
			fmt.Print("Are you sure you want to exit the challenge? (y/N): ")
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
				fmt.Println("\n[Hint]", chal.Hints[hintsUsed].Text)
				hintsUsed++
			} else {
				fmt.Println("\n[No more hints available]")
			}

			continue // no need to validate input

		case "clear":
			clearScreen()
			continue

		case "help":
			// Display challenge details
			fmt.Println("\nChallenge Details:")
			fmt.Println("Title:", chal.Title)
			fmt.Println("Description:", chal.Description)
			fmt.Println("Objective:", chal.Objective)
			fmt.Println("Hints Available:", len(chal.Hints))
			fmt.Println("Reflection Questions:", len(chal.Reflection))
			continue
		}

		// validate user input
		valid := challenge.ValidateCommand(userInput, chal.ExpectedCommand, chal.ValidationType, chal.InputFeedback)
		if valid {
			fmt.Println("\n Correct! Great job.")

			// Show reflection prompts after correct input
			for _, q := range chal.Reflection {
				fmt.Println("-", q)
			}
			break // exit the main loop
		} else {
			fmt.Println("\n Incorrect. Please try again.")
		}
	}
}

// clearScreen clears the terminal screen using ANSI escape codes
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
