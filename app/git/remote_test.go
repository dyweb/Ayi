package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrowserRegexp(t *testing.T) {
	r, err := parseBrowserURL("https://github.com/dyweb/Ayi")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "https", r.Protocol)
	assert.Equal(t, "github.com", r.Host)
	assert.Equal(t, "dyweb", r.Owner)
	assert.Equal(t, "Ayi", r.Repo)
	_, err = parseBrowserURL("file:///D:/tmp/mapreduce.pdf")
	errMsg := "not a browser url"
	assert.Equal(t, errMsg, err.Error())
}
