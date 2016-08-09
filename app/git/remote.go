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
	// TODO: store a HostURL and embed a Host object
	Host string
	// Port is used for remote that does not use default ssh port
	Port int
	// Owner is user name or organization name
	Owner string
	// Repo is repository name
	Repo string
	// SupportHTTPS shows if the host is using HTTPS
	// TODO: change is to a method
	SupportHTTPS bool
}

// GetSSH return the ssh clone address
func (r Remote) GetSSH() string {
	// TODO: handle ssh port
	hostsMap = GetAllHostsMap()
	host, exists := hostsMap[r.Host]
	if !exists {
		log.Warn(r.Host + "does not exists")
		return ""
	}
	return fmt.Sprintf("git@%s:%s/%s.git", host.SSHURL, r.Owner, r.Repo)
}

// Regular expressions used to match remote info
// browserRegexp extract information from browser url like https://github.com/dyweb/Ayi,
// trailing slash, query parameters and .git will be ignored
// (?:u/)? and (?:p/)? are added to deal with coding.net, ie: https://coding.net/u/at15/p/apm-v5/git
var browserRegexp = regexp.MustCompile("^(http|https)://(.+?)/(?:u/)?(.+?)/(?:p/)?([^/?]+?)(?:\\.git)?(?:/.*)?(?:\\?.*)?$")

const browserSegmentsCount = 4

var importRegexp = regexp.MustCompile("^([^/]+?)/([^/]+?)/([^/?]+)/*")

const importSegmentsCount = 3

var shortRegexp = regexp.MustCompile("^([^/]+)/([^/]+)/*$")

const shortSegmentsCount = 2

var httpCloneRegexp = regexp.MustCompile("^(http|https)://([^/?]+)/([^/?]+)/([^/?]+).git$")

const httpCloneSegmentsCount = 4

// NewFromURL returns a remote based on the url, which could be
// - url in browser https://github.com/dyweb/Ayi
// - import url, like import "github.com/dyweb/Ayi/util"
// - short url dyweb/Ayi, default host would be github, but it can also use GitLab based on your config
func NewFromURL(url string) (Remote, error) {
	r := Remote{}
	err := errors.New("invalid url")
	switch {
	case httpCloneRegexp.MatchString(url):
		r, err = parseHttpCloneURL(url)
		if err != nil {
			return r, err
		}
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
	case shortRegexp.MatchString(url):
		r, err = parseShortURL(url)
		if err != nil {
			return r, err
		}
	}
	return r, err
}

func parseHttpCloneURL(url string) (Remote, error) {
	r := Remote{}
	segments := httpCloneRegexp.FindStringSubmatch(url)
	if len(segments) != (httpCloneSegmentsCount + 1) {
		return r, errors.New(fmt.Sprintf("not a http clone url, need %d segments but got %d", httpCloneSegmentsCount, len(segments)-1))
	}
	r.Protocol = segments[1]
	r.Host = segments[2]
	r.Owner = segments[3]
	r.Repo = segments[4]
	if r.Protocol == "https" {
		r.SupportHTTPS = true
	}
	return r, nil
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
	if r.Protocol == "https" {
		r.SupportHTTPS = true
	}
	// TODO: loop the config, user maybe using http url for a https host
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
	// TODO: loop the config, determine https by host name, there are only a few common public hosts
	return r, nil
}

// TODO: deal with github.com/dyweb, this is not <owner>/<repo> but <host>/<owner>
func parseShortURL(url string) (Remote, error) {
	r := Remote{}
	segments := shortRegexp.FindStringSubmatch(url)
	if len(segments) != (shortSegmentsCount + 1) {
		return r, errors.New(fmt.Sprintf("not a short url, need %d segments but got %d", shortSegmentsCount, len(segments)-1))
	}
	r.Owner = segments[1]
	r.Repo = segments[2]
	return r, nil
}
