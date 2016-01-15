package git

// browserRegexp extract information from browser url like https://github.com/dyweb/Ayi
var browserRegexp = "(http|https):\\/\\/(.*?)\\/(.*?)\\/(.*)"

type Remote struct {
	// Protocol is http, https, ssh
	Protocol string
	// Host is the git host provider domain name, like github.com
	Host     string
	// Org is user name or organization name
	Org      string
	// Repo is repository name
	Repo     string
}

func getRemote(url string) (remote Remote, err error) {
	// parse the url
	return remote, err
}

// transformAddress will turn remote browser address to valid ssh and http clone address
func transformAddress(remote string) (sshAddress string, httpAddress string) {
	return "", ""
}

