package git

import (
	"fmt"
	"regexp"

	"github.com/go-errors/errors"
)

// browserRegexp extract information from browser url like https://github.com/dyweb/Ayi
var browserRegexp = "(http|https):\\/\\/(.*?)\\/(.*?)\\/(.*)"

const remoteSegmentsLen = 4

// Remote represents a remote git repo
type Remote struct {
	// Protocol is http, https, ssh
	Protocol string
	// Host is the git host provider domain name, like github.com
	Host string
	// Org is user name or organization name
	Org string
	// Repo is repository name
	Repo string
}

func getRemote(url string) (remote Remote, err error) {
	// parse the url
	r, _ := regexp.Compile(browserRegexp)
	if !r.MatchString(url) {
		return remote, errors.New("invalid browser url" + url)
	}
	segments := r.FindStringSubmatch(url)

	// the matched string is captured as segments[0]
	if len(segments) != (remoteSegmentsLen + 1) {
		return remote, errors.New(fmt.Errorf("got %d segments for url", len(segments)))
	}

	remote = Remote{Protocol: segments[1], Host: segments[2], Org: segments[3], Repo: segments[4]}

	return remote, err
}

// transformAddress will turn remote browser address to valid ssh and http clone address
func transformAddress(url string) (sshAddress string, httpAddress string) {
	remote, _ := getRemote(url)
	// git@github.com:dyweb/Ayi.git
	sshAddress = fmt.Sprintf("git@%s:%s/%s.git",
		remote.Host,
		remote.Org,
		remote.Repo,
	)
	// https://github.com/dyweb/Ayi.git
	httpAddress = fmt.Sprintf("%s://%s/%s/%s.git",
		remote.Protocol,
		remote.Host,
		remote.Org,
		remote.Repo,
	)
	return sshAddress, httpAddress
}
