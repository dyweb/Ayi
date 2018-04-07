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
}
