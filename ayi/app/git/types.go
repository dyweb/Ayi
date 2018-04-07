package git

type Protocol string

const (
	Http  Protocol = "http"
	Https Protocol = "https"
	Ssh   Protocol = "ssh"
)

type Host struct {
}

// Repo represents the address came after git clone <repo>
type Repo struct {
	// Protocol is http, https, ssh
	Protocol Protocol
	// Host is the git host provider domain name, like github.com
	Host string
	// Port is used for remote that does not use default ssh port
	Port int
	// Owner is user name or organization name
	Owner string
	// Repository is name of the repo
	Repository string
}
