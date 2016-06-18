package git

import (
	"regexp"

	"github.com/pkg/errors"
)

// Remote represents a remote git repository
type Remote struct {
	// Name is origin, upstream etc
	Name string
	// Protocol is http, https, ssh
	Protocol string
	// Host is the git host provider domain name, like github.com
	Host string
	// Port is used for remote that does not use default port
	Port int
	// Owner is user name or organization name
	Owner string
	// Repo is repository name
	Repo string
}

// Regular expressions used to match remote info
// browserRegexp extract information from browser url like https://github.com/dyweb/Ayi
var browserRegexp = regexp.MustCompile("(http|https):\\/\\/(.*?)\\/(.*?)\\/(.*)")

// NewFromURL returns a remote based on the url, which could be
// - a url in browser https://github.com/dyweb/Ayi
// - a short url dyweb/Ayi, default host would be github, but it can also use GitLab based on your config
func NewFromURL(url string) (Remote, error) {
	r := Remote{}
	return r, nil
}

func parseBrowserURL(url string) (Remote, error) {
	r := Remote{}
	if !browserRegexp.MatchString(url) {
		return r, errors.New("not a browser url")
	}
	return r, nil
}
