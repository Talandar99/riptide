package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(command string) {
	fmt.Println("Running: " + command)
	cmd := exec.Command(command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
