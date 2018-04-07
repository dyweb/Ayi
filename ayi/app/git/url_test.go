package git

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestUrlToRepo(t *testing.T) {
	t.Run("http-clone", func(t *testing.T) {
		assert := asst.New(t)
		r, err := UrlToRepo("https://github.com/dyweb/Ayi.git")
		assert.Nil(err)
		assert.Equal(Https, r.Protocol)
		assert.Equal("github.com", r.Host)
		assert.Equal("dyweb", r.Owner)
		assert.Equal("Ayi", r.Repository)
	})
	t.Run("http-common", func(t *testing.T) {
		assert := asst.New(t)
		urls := []string{
			"https://github.com/dyweb/Ayi/blob/master/README.md",
			"https://github.com/dyweb/Ayi/blob/master/README.md#about",
			"https://github.com/dyweb/Ayi?test=1&foo=2",
			"https://github.com/dyweb/Ayi#abc",
			"https://github.com/dyweb/Ayi",
		}
		for _, u := range urls {
			r, err := UrlToRepo(u)
			assert.Nil(err)
			assert.Equal(Https, r.Protocol)
			assert.Equal("github.com", r.Host)
			assert.Equal("dyweb", r.Owner)
			assert.Equal("Ayi", r.Repository)
		}
	})
	t.Run("common", func(t *testing.T) {
		assert := asst.New(t)
		urls := []string{
			"github.com/dyweb/Ayi/blob/master/README.md",
			"github.com/dyweb/Ayi/blob/master/README.md#about",
			"github.com/dyweb/Ayi?test=1&foo=2",
			"github.com/dyweb/Ayi#abc",
			"github.com/dyweb/Ayi",
		}
		for _, u := range urls {
			r, err := UrlToRepo(u)
			assert.Nil(err)
			assert.Equal(Protocol(""), r.Protocol)
			assert.Equal("github.com", r.Host)
			assert.Equal("dyweb", r.Owner)
			assert.Equal("Ayi", r.Repository)
		}
	})
	t.Run("ssh-noport", func(t *testing.T) {
		assert := asst.New(t)
		r, err := UrlToRepo("git@github.com:dyweb/Ayi.git")
		assert.Nil(err)
		assert.NotNil(r)
		assert.Equal(Ssh, r.Protocol)
		assert.Equal("github.com", r.Host)
		assert.Equal("dyweb", r.Owner)
		assert.Equal("Ayi", r.Repository)
	})
	t.Run("ssh-port", func(t *testing.T) {
		assert := asst.New(t)
		r, err := UrlToRepo("ssh://git@git.dongyue.io:6773/at15/tongqu4.git")
		assert.Nil(err)
		assert.NotNil(r)
		assert.Equal(Ssh, r.Protocol)
		assert.Equal("git.dongyue.io", r.Host)
		assert.Equal("at15", r.Owner)
		assert.Equal("tongqu4", r.Repository)
	})
	t.Run("owner/repo", func(t *testing.T) {
		assert := asst.New(t)
		r, err := UrlToRepo("dyweb/Ayi")
		assert.Nil(err)
		assert.Equal(Protocol(""), r.Protocol)
		assert.Equal("github.com", DefaultHost())
		assert.Equal(DefaultHost(), r.Host)
		assert.Equal("dyweb", r.Owner)
		assert.Equal("Ayi", r.Repository)
	})
	t.Run("repo", func(t *testing.T) {
		assert := asst.New(t)
		r, err := UrlToRepo("go.ice")
		assert.Nil(err)
		assert.Equal(Protocol(""), r.Protocol)
		assert.Equal(DefaultHost(), r.Host)
		// FIXME: need to add get current user name to gommon/testutil
		t.Log(r.Owner)
		//assert.Equal("at15", r.Owner)
		assert.Equal("go.ice", r.Repository)
	})
}
