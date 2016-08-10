package git

import (
	"os"
	"path/filepath"

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
	// use current folder as repo base path
	cwd, err := os.Getwd()
	if err != nil {
		log.Warn("Can't get working directory " + err.Error())
		return ""
	}
	return cwd
}
