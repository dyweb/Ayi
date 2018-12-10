package git

import (
	"os/exec"

	"os"

	"github.com/dyweb/gommon/errors"
	shellquote "github.com/kballard/go-shellquote"
)

func Command(cmdWithoutGit string) (*exec.Cmd, error) {
	args, err := shellquote.Split(cmdWithoutGit)
	if err != nil {
		return nil, errors.Wrap(err, "can't split cmd")
	}
	return exec.Command("git", args...), nil
}

// RunCommand execute the command and redirect input/output to standard input/output
func RunCommand(cmdWithoutGit string) error {
	cmd, err := Command(cmdWithoutGit)
	if err != nil {
		return err
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "failed when executing git %s", cmdWithoutGit)
	}
	return nil
}
