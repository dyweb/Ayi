package util

import (
	"os"
	"os/exec"

	"github.com/kballard/go-shellquote"
	"github.com/pkg/errors"
)

// Command return a Command struct from a full commad
func Command(cmd string) (*exec.Cmd, error) {
	segments, err := shellquote.Split(cmd)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot parse command")
	}
	name := segments[0]
	// FIXME: this is not working ...
	// if (name == "sh") && (segments[1] == "-c") {
	// 	// TODO: this does not support use like go test $(glide novendor)
	// 	fmt.Println(strings.Join(segments[2:], " "))
	// 	return exec.Command("sh", "-c", strings.Join(segments[2:], " "))
	// }
	return exec.Command(name, segments[1:]...), nil
}

// RunCommand runs a commad and show all output in console, block current routine
func RunCommand(cmd string) error {
	command, err := Command(cmd)
	if err != nil {
		return err
	}
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Run()
	if err != nil {
		return errors.Wrap(err, "Failure when executing command")
	}
	return nil
}
