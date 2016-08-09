package git

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoPath(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual("", os.Getenv("GOPATH"))
	assert.Equal("", os.Getenv("GOPATHH"))

	p := os.Getenv("GOPATH")
	os.Setenv("GOPATH", "")
	assert.Equal("", os.Getenv("GOPATH"))
	os.Setenv("GOPATH", p)
}

func TestGetRepoBasePath(t *testing.T) {
	assert := assert.New(t)
	// assert.Equal("D:\\workspace\\src", GetRepoBasePath())
	p := os.Getenv("GOPATH")
	os.Setenv("GOPATH", "")
	assert.Equal(filepath.FromSlash("/home/at15/repos"), GetRepoBasePath())
	os.Setenv("GOPATH", p)
	assert.Equal(filepath.FromSlash(p+"/src"), GetRepoBasePath())

}
