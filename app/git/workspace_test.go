package git

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
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
	p := os.Getenv("GOPATH")
	// config file is used evne if gopath is not empty
	assert.Equal(filepath.FromSlash("/home/at15/repos"), GetRepoBasePath())
	// gopath is used if config file is not set
	viper.Set("git.repositories", "")
	os.Setenv("GOPATH", p)
	assert.Equal(filepath.FromSlash(p+"/src"), GetRepoBasePath())

}
