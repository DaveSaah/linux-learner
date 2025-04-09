package commands

import (
	"os/exec"
	"strings"
)

type CommandResult struct {
	Output string
	Stderr string
	Err    error
}

func RunCommand(command string) CommandResult {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()

	stderr := ""
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			stderr = string(exitErr.Stderr)
		}
	}

	return CommandResult{
		Output: strings.TrimSpace(string(output)),
		Stderr: stderr,
		Err:    err,
	}
}
