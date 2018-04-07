package git

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/dyweb/gommon/errors"
)

var (
	// ?: is non capture group for using group but don't capture content in return result
	// https://stackoverflow.com/questions/3512471/what-is-a-non-capturing-group-what-does-do
	httpCommonRegexp = regexp.MustCompile("^([^/]+?)/([^/]+?)/([^/]+?)(?:\\.git)?(?:[/#?].*)?$")
	sshCloneRegexp   = regexp.MustCompile("^(?:ssh://)?git@([^/]+?):(\\d+)?/?([^/]+?)/([^/]+?).git$")
	ownerRepoRegexp  = regexp.MustCompile("^([^/]+?)/([^/]+?)(?:.git)?$")
	repoRegexp       = regexp.MustCompile("^([^/]+?)(?:.git)?$")
)

const (
	httpCommonSegments = 3 // host + owner + repo
	sshCloneSegments   = 4 // host + port + owner + repo
	ownerRepoSegments  = 2
	repoSegments       = 1
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
	case httpCommonRegexp.MatchString(u):
		return parseCommon(u)
	case sshCloneRegexp.MatchString(u):
		return parseSsh(u)
	case ownerRepoRegexp.MatchString(u):
		segments := ownerRepoRegexp.FindStringSubmatch(u)
		if len(segments)-1 != ownerRepoSegments {
			return nil, errors.Errorf("invalid <owner>/<repository> %s", u)
		}
		return &Repo{Host: DefaultHost(), Owner: segments[1], Repository: segments[2]}, nil
	case repoRegexp.MatchString(u):
		segments := repoRegexp.FindStringSubmatch(u)
		if len(segments)-1 != repoSegments {
			return nil, errors.Errorf("invalid <repository> %s", u)
		}
		return &Repo{Host: DefaultHost(), Owner: DefaultUser(), Repository: segments[1]}, nil
	default:
		return nil, errors.Errorf("unknown url pattern %s", u)
	}
}

// http url
// - clone url: https://github.com/dyweb/Ayi.git
// - common url: https://github.com/dyweb/Ayi/blob/master/README.md
func parseHttp(u string) (*Repo, error) {
	proto := Https
	if strings.HasPrefix(u, "https://") {
		proto = Https
		u = u[len("https://"):]
	} else if strings.HasPrefix(u, "http://") {
		proto = Http
		u = u[len("http://"):]
	} else {
		return nil, errors.Errorf("invalid http protocol %s", u)
	}
	r, err := parseCommon(u)
	if r != nil {
		r.Protocol = proto
	}
	return r, err
}

// http url with protocol prefix
func parseCommon(u string) (*Repo, error) {
	r := Repo{}

	segments := httpCommonRegexp.FindStringSubmatch(u)
	if len(segments)-1 != httpCommonSegments {
		return nil, errors.Errorf("can't infer repo from url, got %d segments instead of %d", len(segments)-1, httpCommonSegments)
	}
	r.Host, r.Owner, r.Repository = segments[1], segments[2], segments[3]

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
