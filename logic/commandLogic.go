package logic

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/razaibi/gmx/models"
)

func RunCommand(item models.Item) {
	output, err := getOutput(item.Cmd)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Output:\n%s\n", output)
	}
}

func getOutput(commandString string) (string, error) {
	// Split the command string into command and arguments
	parts := strings.Fields(commandString)
	if len(parts) == 0 {
		return "", fmt.Errorf("empty command")
	}

	// The first part is the command, the rest are arguments
	command := parts[0]
	args := parts[1:]

	// Create the command
	cmd := exec.Command(command, args...)

	// Create buffers for stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	// If there's an error, return it along with any stderr output
	if err != nil {
		return "", fmt.Errorf("error executing command: %v\nStderr: %s", err, stderr.String())
	}

	// Return the stdout output
	return stdout.String(), nil
}
