package git

import (
	"regexp"
	"strings"
	"strconv"

	"github.com/dyweb/gommon/errors"
)

var (
	httpCloneRegexp = regexp.MustCompile("^([^/]+?)/([^/]+?)/([^/]+?).git$")
	// ?: is non capture group for using group but don't capture content in return result
	// https://stackoverflow.com/questions/3512471/what-is-a-non-capturing-group-what-does-do
	httpCommonRegexp   = regexp.MustCompile("^([^/]+?)/([^/]+?)/([^/]+?)(?:\\.git)?(?:[/#?].*)?$")
	sshCloneRegexp     = regexp.MustCompile("^(?:ssh://)?git@([^/]+?):(\\d+)?/?([^/]+?)/([^/]+?).git$")
	ownerProjectRegexp = regexp.MustCompile("^([^/]+?)/([^/]+?)(?:.git)?$")
	//projectRegexp = regexp.MustCompile("^([^/]+?)(?:.git)?$")
)

const (
	httpCloneSegments    = 3 // host + owner + repo
	httpCommonSegments   = 3
	sshCloneSegments     = 4 // host + port + owner + repo
	ownerProjectSegments = 2
)

// UrlToRepo detect git repository from url
// http clone url: https://github.com/dyweb/Ayi.git
// http common url: https://github.com/dyweb/Ayi/blob/master/README.md
// ssh clone url: git@github.com:dyweb/Ayi.git ssh://git@git.dongyue.io:6773/at15/tongqu4.git
// project name: go.ice -> github.com/at15/go.ice
// owner/project: dyweb/Ayi -> github.com/dyweb/Ayi
func UrlToRepo(u string) (*Repo, error) {
	u = strings.TrimSpace(u)
	switch {
	case strings.HasPrefix(u, "http"):
		return parseHttp(u)
	case sshCloneRegexp.MatchString(u):
		return parseSsh(u)
	case ownerProjectRegexp.MatchString(u):
		segments := ownerProjectRegexp.FindStringSubmatch(u)
		if len(segments)-1 != 2 {

		}
	default:

	}
	return nil, errors.Errorf("unknown url pattern %s", u)
}

// http url
// - clone url: https://github.com/dyweb/Ayi.git
// - common url: https://github.com/dyweb/Ayi/blob/master/README.md
func parseHttp(u string) (*Repo, error) {
	r := Repo{}
	if strings.HasPrefix(u, "https://") {
		r.Protocol = Https
		u = u[len("https://"):]
	} else if strings.HasPrefix(u, "http://") {
		r.Protocol = Http
		u = u[len("http://"):]
	} else {
		return nil, errors.Errorf("invalid http protocol %s", u)
	}
	switch {
	case httpCloneRegexp.MatchString(u):
		// http clone url: https://github.com/dyweb/Ayi.git
		segments := httpCloneRegexp.FindStringSubmatch(u)
		if len(segments)-1 != httpCloneSegments {
			return nil, errors.Errorf("not a http clone url, got %d segments instead of %d", len(segments)-1, httpCloneSegments)
		}
		r.Host, r.Owner, r.Repository = segments[1], segments[2], segments[3]
	case httpCommonRegexp.MatchString(u):
		// http common url: https://github.com/dyweb/Ayi/blob/master/README.md
		segments := httpCommonRegexp.FindStringSubmatch(u)
		if len(segments)-1 != httpCommonSegments {
			return nil, errors.Errorf("can't infer repo from http url, got %d segments instead of %d", len(segments)-1, httpCommonSegments)
		}
		r.Host, r.Owner, r.Repository = segments[1], segments[2], segments[3]
	default:
		return nil, errors.Errorf("no matched http pattern %s", u)
	}
	return &r, nil
}

func parseSsh(u string) (*Repo, error) {
	r := Repo{}
	r.Protocol = Ssh
	segments := sshCloneRegexp.FindStringSubmatch(u)
	if len(segments)-1 != sshCloneSegments {
		return nil, errors.Errorf("not a git ssh clone url, got %d segments instead of %d", len(segments)-1, sshCloneSegments)
	}
	port := 0
	if segments[2] != "" {
		// ssh port ssh://git@git.dongyue.io:6773/at15/tongqu4.git
		port, err := strconv.Atoi(segments[2])
		if err != nil || port == 0 {
			return nil, errors.Wrap(err, "invalid ssh port")
		}
	}
	// git@github.com:dyweb/Ayi.git
	r.Host, r.Port, r.Owner, r.Repository = segments[1], port, segments[3], segments[4]
	return &r, nil
}
