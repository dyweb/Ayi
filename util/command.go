package util

import (
	"os"
	"os/exec"

	"github.com/kballard/go-shellquote"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Command return a Command struct from a full commad
func Command(cmd string) (*exec.Cmd, error) {
	// NOTE: do not use strings.Fields or Split !
	segments, err := shellquote.Split(cmd)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot parse command")
	}
	name := segments[0]
	return exec.Command(name, segments[1:]...), nil
}

// RunCommand runs a commad and show all output in console, block current routine
func RunCommand(cmd string) error {
	command, err := Command(cmd)
	if err != nil {
		return err
	}
	if viper.GetBool("DryRun") == true {
		// TODO: need a new error type for dry run
		return nil
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
