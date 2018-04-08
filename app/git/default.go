package git

import (
	"os"
	"os/user"
	"path/filepath"
)

const (
	nonExistUser   = "404"
	defaultHost    = "github.com"
	defaultSshPort = 22
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

func DefaultWorkspace() string {
	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		return filepath.Join(gopath, "src")
	}
	log.Warn("unable to find default workspace")
	return ""
}
