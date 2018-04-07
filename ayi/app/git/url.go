package git

import (
	"regexp"
	"strings"

	"github.com/dyweb/gommon/errors"
)

var (
	httpCloneRegexp = regexp.MustCompile("^([^/]+?)/([^/]+?)/([^/]+?).git$")
	// ?: is non capture group for using group but don't capture content in return result
	// https://stackoverflow.com/questions/3512471/what-is-a-non-capturing-group-what-does-do
	httpCommonRegexp = regexp.MustCompile("^([^/]+?)/([^/]+?)/([^/]+?)(?:\\.git)?(?:[/#?].*)?$")
)

const (
	httpCloneSegments  = 3
	httpCommonSegments = 3
)

func UrlToRepo(u string) (*Repo, error) {
	u = strings.TrimSpace(u)
	switch {
	case strings.HasPrefix(u, "http"):
		return parseHttp(u)
	default:
		return nil, errors.Errorf("unknown url pattern %s", u)
	}
}

// http url
// - http clone url: https://github.com/dyweb/Ayi.git
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
		// common url: https://github.com/dyweb/Ayi/blob/master/README.md
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
