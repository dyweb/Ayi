package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrowserRegexp(t *testing.T) {
	r, err := parseBrowserURL("https://github.com/dyweb/Ayi")
	assert.Equal(t, nil, err)
	assert.Equal(t, "https", r.Protocol)
	assert.Equal(t, "github.com", r.Host)
	assert.Equal(t, "dyweb", r.Owner)
	assert.Equal(t, "Ayi", r.Repo)
	r, err = parseBrowserURL("https://github.com/dyweb/Ayi/")
	assert.Equal(t, "Ayi", r.Repo)
	r, err = parseBrowserURL("https://github.com/dyweb/Ayi/util")
	assert.Equal(t, "Ayi", r.Repo)
	r, err = parseBrowserURL("https://github.com/dyweb/Ayi?detail=1")
	assert.Equal(t, "Ayi", r.Repo)
	_, err = parseBrowserURL("file:///D:/tmp/mapreduce.pdf")
	errMsg := "not a browser url"
	assert.Contains(t, err.Error(), errMsg)
}

func TestImportRegexp(t *testing.T) {
	r, err := parseImportURL("github.com/dyweb/Ayi")
	assert.Equal(t, nil, err)
	assert.Equal(t, "github.com", r.Host)
	assert.Equal(t, "dyweb", r.Owner)
	assert.Equal(t, "Ayi", r.Repo)
	r, err = parseImportURL("github.com/dyweb/Ayi/")
	assert.Equal(t, "Ayi", r.Repo)
	r, err = parseImportURL("github.com/dyweb/Ayi/util")
	assert.Equal(t, "Ayi", r.Repo)
	r, err = parseImportURL("github.com/dyweb/Ayi/util/")
	assert.Equal(t, "Ayi", r.Repo)
}

func TestShortRegexp(t *testing.T) {
	r, err := parseShortURL("dyweb/web-stuff")
	assert.Equal(t, nil, err)
	assert.Equal(t, "web-stuff", r.Repo)
}

func TestNewFromURL(t *testing.T) {
	r, err := NewFromURL("github.com/dyweb/Ayi")
	assert.Equal(t, nil, err)
	assert.Equal(t, "Ayi", r.Repo)
	r, err = NewFromURL("http://github.com/dyweb/Ayi")
	assert.Equal(t, "Ayi", r.Repo)
	r, err = NewFromURL("dyweb/Ayi")
	assert.Equal(t, "Ayi", r.Repo)
	r, err = NewFromURL("file:///D:/tmp/mapreduce.pdf")
	assert.NotEqual(t, nil, err)
}
