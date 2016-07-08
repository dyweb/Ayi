package util

import (
	"os"
	"os/exec"
	"strings"
)

// Command return a Command struct from a full commad
func Command(cmd string) *exec.Cmd {
	segments := strings.Split(cmd, " ")
	name := segments[0]
	return exec.Command(name, segments[1:]...)
}

// RunCommand runs a commad and show all output in console, block current routine
func RunCommand(cmd string) error {
	command := Command(cmd)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	// TODO: wrap it up using errors
	err := command.Run()
	return err
}
