package git

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// GetRepoBasePath return path in ordr of : config, GOPATH, current directory
func GetRepoBasePath() string {
	// check if repo base is config in file or other way through viper
	r := viper.GetString("git.repositories")
	if r != "" {
		return filepath.FromSlash(r)
	}
	// use go path if provided
	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		return filepath.FromSlash(gopath + "/src")
	}
	return ""
}

// GetCloneDirectory return path for clone destination
func GetCloneDirectory(r Remote) string {
	base := GetRepoBasePath()
	if base == "" {
		return ""
	}
	// git clone does not accept windows format path, so all /
	return strings.Replace(base+"/"+r.Host+"/"+r.Owner+"/"+r.Repo, "\\", "/", -1)
}
