package git

import (
	"os/exec"
)

var cmdName = "git"

// simpleCommand is a interface execute and return output as string
type simpleCommand interface {
	Execute() (stdOut string, stdErr string, err error)
}

// Status execute git status command and return output
type Status struct {
}

// Execute execute git status and return output as string
func (s Status) Execute() (stdOut string, stdErr string, err error) {
	var cmdOut []byte
	err = nil
	cmdArgs := []string{"status"}
	cmdOut, err = exec.Command(cmdName, cmdArgs...).Output()
	// FIXME: cant get correct error output
	if err != nil{
		stdErr = err.Error()
	}
	// TODO: get return code?
	return string(cmdOut), stdErr, err
}
