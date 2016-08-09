package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesstHttpCloneRegexp(t *testing.T) {
	assert := assert.New(t)
	r, err := parseHttpCloneURL("https://bitbucket.org/at6/kc-3g.git")
	assert.Nil(err)
	assert.Equal("bitbucket.org", r.Host)
	_, err = parseHttpCloneURL("https://coding.net/u/at15/p/apm-v5/git")
	assert.NotNil(err)
}

func TestSSHCloneRegexp(t *testing.T) {
	assert := assert.New(t)
	segments := sshCloneRegexp.FindStringSubmatch("ssh://git@git.tongqu.me:3022/at15/tongqu4.git")
	assert.NotZero(segments)
	assert.Equal(5, len(segments))
	segments = sshCloneRegexp.FindStringSubmatch("git@github.com:dyweb/Ayi.git")
	assert.Equal(5, len(segments))
	// t.Log(segments)
	// port is empty
	assert.Equal("", segments[2])
}

func TestBrowserRegexp(t *testing.T) {
	assert := assert.New(t)
	r, err := parseBrowserURL("https://github.com/dyweb/Ayi")
	assert.Nil(err)
	assert.Equal("https", r.Protocol)
	assert.Equal("github.com", r.Host)
	assert.Equal("dyweb", r.Owner)
	assert.Equal("Ayi", r.Repo)
	assert.Equal(true, r.SupportHTTPS)
	// ignore trailing slash
	r, err = parseBrowserURL("https://github.com/dyweb/Ayi/")
	assert.Equal("Ayi", r.Repo)
	// ignore extra segements
	r, err = parseBrowserURL("https://github.com/dyweb/Ayi/util")
	assert.Equal("Ayi", r.Repo)
	// ignore query parameters
	r, err = parseBrowserURL("https://github.com/dyweb/Ayi?detail=1")
	assert.Equal("Ayi", r.Repo)
	// ignore .git
	r, err = parseBrowserURL("https://bitbucket.org/at6/kc-3g.git")
	assert.Equal("kc-3g", r.Repo)
	// coding.net has strange url syntax
	r, err = parseBrowserURL("https://coding.net/u/at15/p/apm-v5/")
	assert.Equal("apm-v5", r.Repo)
	// poor tongqu dev can't afford https
	r, err = parseBrowserURL("http://poor.tongqu.me/at15/tongqu4")
	assert.Equal(false, r.SupportHTTPS)
	_, err = parseBrowserURL("file:///D:/tmp/mapreduce.pdf")
	errMsg := "not a browser url"
	assert.Contains(err.Error(), errMsg)
}

func TestImportRegexp(t *testing.T) {
	assert := assert.New(t)
	r, err := parseImportURL("github.com/dyweb/Ayi")
	assert.Nil(err)
	assert.Equal("github.com", r.Host)
	assert.Equal("dyweb", r.Owner)
	assert.Equal("Ayi", r.Repo)
	r, err = parseImportURL("github.com/dyweb/Ayi/")
	assert.Equal("Ayi", r.Repo)
	r, err = parseImportURL("github.com/dyweb/Ayi/util")
	assert.Equal("Ayi", r.Repo)
	r, err = parseImportURL("github.com/dyweb/Ayi/util/")
	assert.Equal("Ayi", r.Repo)
}

func TestShortRegexp(t *testing.T) {
	r, err := parseShortURL("dyweb/web-stuff")
	assert.Equal(t, nil, err)
	assert.Equal(t, "web-stuff", r.Repo)
}

func TestNewFromURL(t *testing.T) {
	assert := assert.New(t)
	r, err := NewFromURL("https://bitbucket.org/at6/kc-3g.git")
	assert.Nil(err)
	assert.Equal("bitbucket.org", r.Host)
	r, err = NewFromURL("ssh://git@git.tongqu.me:3022/at15/tongqu4.git")
	assert.Nil(err)
	assert.Equal(3022, r.Port)
	r, err = NewFromURL("git@github.com:dyweb/Ayi.git")
	assert.Nil(err)
	assert.Equal("dyweb", r.Owner)
	r, err = NewFromURL("github.com/dyweb/Ayi")
	assert.Nil(err)
	assert.Equal("Ayi", r.Repo)
	r, err = NewFromURL("http://github.com/dyweb/Ayi")
	assert.Equal("Ayi", r.Repo)
	r, err = NewFromURL("dyweb/Ayi")
	assert.Equal("Ayi", r.Repo)
	r, err = NewFromURL("file:///D:/tmp/mapreduce.pdf")
	assert.NotNil(err)
}

func TestGetSSH(t *testing.T) {
	assert := assert.New(t)
	// github
	r, err := NewFromURL("github.com/dyweb/Ayi")
	assert.Nil(err)
	assert.Equal("git@github.com:dyweb/Ayi.git", r.GetSSH())
	// gitlab
	r, err = NewFromURL("https://gitlab.com/leanlabsio/kanban")
	assert.Nil(err)
	assert.Equal("git@gitlab.com:leanlabsio/kanban.git", r.GetSSH())
	// bitbucket
	// FIXME: http clone break this
	r, _ = NewFromURL("https://bitbucket.org/at6/kc-3g.git")
	assert.Equal("bitbucket.org", r.Host)
	assert.Equal("git@bitbucket.org:at6/kc-3g.git", r.GetSSH())
	// coding.net
	// FIXME: http clone break this
	r, _ = NewFromURL("https://coding.net/u/at15/p/apm-v5/git")
	assert.Equal("coding.net", r.Host)
	assert.Equal("git@git.coding.net:at15/apm-v5.git", r.GetSSH())
	// oschina
	r, _ = NewFromURL("http://git.oschina.net/caixw/apidoc")
	assert.Equal("git@git.oschina.net:caixw/apidoc.git", r.GetSSH())

	// Non-default port
	// Port is from ssh url
	r, _ = NewFromURL("ssh://git@git.tongqu.me:3022/at15/tongqu4.git")
	assert.Equal("git@git.tongqu.me:3022/at15/tongqu4.git", r.GetSSH())
	// Port is from config file's host config
	r, _ = NewFromURL("http://git.saber.io/mie/draft")
	assert.Equal("git@git.saber.io:10086/mie/draft.git", r.GetSSH())
}
