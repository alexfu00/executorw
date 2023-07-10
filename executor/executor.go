package executor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ExecuteGoWithFile(code string) (string, error) {

	file, err := os.CreateTemp("", "go_code_*.go")
	if err != nil {
		return "", fmt.Errorf("error creating temporal file: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(code)
	if err != nil {
		return "", fmt.Errorf("error writing to temporal file: %v", err)
	}

	var output bytes.Buffer
	cmd := exec.Command("go", "run", file.Name())
	cmd.Stdout = &output
	cmd.Stderr = &output

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing code: %v", err)
	}

	return output.String(), nil
}
