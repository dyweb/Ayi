package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrowserRegexp(t *testing.T) {
	assert := assert.New(t)
	r, err := parseBrowserURL("https://github.com/dyweb/Ayi")
	assert.Nil(err)
	assert.Equal("https", r.Protocol)
	assert.Equal("github.com", r.Host)
	assert.Equal("dyweb", r.Owner)
	assert.Equal("Ayi", r.Repo)
	r, err = parseBrowserURL("https://github.com/dyweb/Ayi/")
	assert.Equal("Ayi", r.Repo)
	r, err = parseBrowserURL("https://github.com/dyweb/Ayi/util")
	assert.Equal("Ayi", r.Repo)
	r, err = parseBrowserURL("https://github.com/dyweb/Ayi?detail=1")
	assert.Equal("Ayi", r.Repo)
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
	r, err := NewFromURL("github.com/dyweb/Ayi")
	assert.Nil(err)
	assert.Equal("Ayi", r.Repo)
	r, err = NewFromURL("http://github.com/dyweb/Ayi")
	assert.Equal("Ayi", r.Repo)
	r, err = NewFromURL("dyweb/Ayi")
	assert.Equal("Ayi", r.Repo)
	r, err = NewFromURL("file:///D:/tmp/mapreduce.pdf")
	assert.NotNil(err)
}
