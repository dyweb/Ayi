package git

import "os/user"

const (
	nonExistUser = "404"
	defaultHost  = "github.com"
)

var (
	cachedDefaultUser = ""
)

func DefaultHost() string {
	return defaultHost
}

func DefaultUser() string {
	if cachedDefaultUser != "" {
		return cachedDefaultUser
	}
	// TODO: check ~/.gitconfig to grab user name
	// git config --global --get user.name
	if u, err := user.Current(); err == nil {
		return u.Username // Username is the login name, Name is display name
	}
	return nonExistUser
}
