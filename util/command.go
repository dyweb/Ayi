package util

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kballard/go-shellquote"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// DryRunError indicate user specified -n flag and the command is only displayed but not actually exectuted
// https://gobyexample.com/errors
// https://blog.golang.org/error-handling-and-go
type DryRunError struct {
	Command string
}

func (e *DryRunError) Error() string {
	return fmt.Sprintf("dry run (-n) flag is specified for command %s", e.Command)
}

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
		return &DryRunError{cmd}
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
