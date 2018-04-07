package git

const (
	defaultHost = "github.com"
)

func DefaultHost() string {
	return defaultHost
}

func DefaultUser() string {
	// TODO: check ~/.gitconfig to grab user name
	// TODO: use current user name
	return ""
}
