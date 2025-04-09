package challenge

import (
	"embed"

	"gopkg.in/yaml.v3"
)

type Hint struct {
	Level int    `yaml:"level"`
	Text  string `yaml:"text"`
}

type Challenge struct {
	ID              string   `yaml:"id"`
	Title           string   `yaml:"title"`
	Topic           string   `yaml:"topic"`
	Description     string   `yaml:"description"`
	Objective       string   `yaml:"objective"`
	ExpectedCommand string   `yaml:"expected_command"`
	ValidationType  string   `yaml:"validation_type"`
	Hints           []Hint   `yaml:"hints"`
	Reflection      []string `yaml:"reflection"`
}

func LoadChallenge(fs embed.FS, path string) (*Challenge, error) {
	// read file content
	data, err := fs.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// load into data structure
	var challenge Challenge
	err = yaml.Unmarshal(data, &challenge)
	if err != nil {
		return nil, err
	}

	return &challenge, nil
}
