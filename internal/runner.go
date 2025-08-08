package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func runSystemCommands(command []string) {
	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func RunCommand(command Script) {
	fmt.Println("Running: " + command.Path + command.Name)
	runSystemCommands([]string{command.Path + command.Name})
}
func RunRemoteCommand(command Script, remoteAddress string) {
	fmt.Println("Copying " + command.Name + " via scp to " + remoteAddress + "~/" + command.Name)
	runSystemCommands([]string{"scp", command.Path + command.Name, remoteAddress + ":~/" + command.Name})
	fmt.Println("Running: " + command.Name + "Remotely")
}
