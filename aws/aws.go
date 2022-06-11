package aws

import (
	"os/exec"
)

type execContext = func(commandName string, args ...string) *exec.Cmd

func Version(cmdContext execContext) ([]byte, error) {
	command := cmdContext("aws", "--version")

	output, err := command.CombinedOutput()
	if err != nil {
		return output, err
	}

	return output, nil
}
