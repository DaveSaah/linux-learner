package commands

import (
	"os/exec"
	"strings"
)

type CommandResult struct {
	Output string // The standard output of the executed command
	Stderr string // The standard error output (if any) from the command
	Err    error  // Any error that occurred during the execution of the command
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
