package challenge

import (
	"embed"

	"gopkg.in/yaml.v3"
)

// Hint represents a hint provided to the learner.
// Each hint has a level to indicate when it should be shown.
type Hint struct {
	Level int    `yaml:"level"` // Progressive hint level (1 = subtle, higher is more obvious)
	Text  string `yaml:"text"`  // Text of the hint
}

// InputFeedback maps user input patterns to a helpful feedback message.
// This supports more interactive and contextual guidance.
type InputFeedback struct {
	Match    string `yaml:"match"`    // User input pattern to match (can be exact or regex)
	Response string `yaml:"response"` // Response shown when this pattern is matched
}

// ReflectionQuestion represents a post-challenge prompt for deeper learning.
// It can optionally require the learner to provide an answer.
type ReflectionQuestion struct {
	Question string `yaml:"question"` // The reflection question text
}

// Challenge represents the structure of a YAML challenge file.
// It defines the scenario, goals, expected commands, feedback mechanisms, and reflection.
type Challenge struct {
	ID              string               `yaml:"id"`               // Unique challenge identifier
	Title           string               `yaml:"title"`            // Short challenge title
	Topic           string               `yaml:"topic"`            // Category or theme (e.g., navigation, permissions)
	Scenario        string               `yaml:"scenario"`         // Real-world context or backstory for immersion
	Description     string               `yaml:"description"`      // Specific task instructions for the learner
	Objective       string               `yaml:"objective"`        // Learning goal of this challenge
	ExpectedCommand string               `yaml:"expected_command"` // Correct command input
	ValidationType  string               `yaml:"validation_type"`  // Method of validation (e.g., output_match)
	Hints           []Hint               `yaml:"hints"`            // List of progressively revealed hints
	InputFeedback   []InputFeedback      `yaml:"input_feedback"`   // Input-based guidance messages
	Reflection      []ReflectionQuestion `yaml:"reflection"`       // Post-task prompts for deeper learning
	SuccessMessage  string               `yaml:"success_message"`  // Shown when user completes the challenge
}

// LoadChallenge loads and parses a challenge YAML file from the
// embedded filesystem.
// Parameters:
//
//	fs   - embedded filesystem containing challenge files
//	path - path to the YAML file inside the FS
//
// Returns a populated Challenge object or an error if loading fails.
func LoadChallenge(fs embed.FS, path string) (*Challenge, error) {
	// Read YAML file from embedded filesystem
	data, err := fs.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal YAML into Challenge struct
	var challenge Challenge
	err = yaml.Unmarshal(data, &challenge)
	if err != nil {
		return nil, err
	}

	return &challenge, nil
}
