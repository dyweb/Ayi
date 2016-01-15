package git

var cmdName = "git"

// simpleCommand is a interface execute and return output as string
type simpleCommand interface {
	Execute() (stdOut string, stdErr string, err error)
}
