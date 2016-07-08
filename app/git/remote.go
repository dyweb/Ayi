package git

import (
	"fmt"
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
// TODO: support url with trailng stuff like https://github.com/dyweb/Ayi/issues
var browserRegexp = regexp.MustCompile("(http|https):\\/\\/(.+?)\\/(.+?)\\/([^\\/?]+)/*")

const browserSegmentsCount = 4

var importRegexp = regexp.MustCompile("(.+?)\\/(.+?)\\/([^\\/?]+)/*")

const importSegmentsCount = 3

// NewFromURL returns a remote based on the url, which could be
// - url in browser https://github.com/dyweb/Ayi
// - import url, like import "github.com/dyweb/Ayi/util"
// - short url dyweb/Ayi, default host would be github, but it can also use GitLab based on your config
func NewFromURL(url string) (Remote, error) {
	r := Remote{}
	err := errors.New("invalid url")
	switch {
	case browserRegexp.MatchString(url):
		r, err = parseBrowserURL(url)
		if err != nil {
			return r, err
		}
	case importRegexp.MatchString(url):
		r, err = parseImportURL(url)
		if err != nil {
			return r, err
		}
	}
	return r, err
}

func parseBrowserURL(url string) (Remote, error) {
	r := Remote{}
	// captuer protocol, host, organization and repository
	segments := browserRegexp.FindStringSubmatch(url)
	// segments[0] is the matched string, ie: https://github.com/dyweb/Ayi
	if len(segments) != (browserSegmentsCount + 1) {
		return r, errors.New(fmt.Sprintf("not a browser url, need %d segments but got %d", browserSegmentsCount, len(segments)-1))
	}
	r.Protocol = segments[1]
	r.Host = segments[2]
	r.Owner = segments[3]
	r.Repo = segments[4]
	return r, nil
}

func parseImportURL(url string) (Remote, error) {
	r := Remote{}
	segments := importRegexp.FindStringSubmatch(url)
	if len(segments) != (importSegmentsCount + 1) {
		return r, errors.New(fmt.Sprintf("not a import url, need %d segments but got %d", importSegmentsCount, len(segments)-1))
	}
	r.Host = segments[1]
	r.Owner = segments[2]
	r.Repo = segments[3]
	return r, nil
}
