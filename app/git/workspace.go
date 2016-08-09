package git

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func GetRepoBasePath() string {
	// use go path if provided
	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		return filepath.FromSlash(gopath + "/src")
	}
	// check if repo base is config in file or other way through viper
	r := viper.GetString("git.repositories")
	if r != "" {
		return filepath.FromSlash(r)
	}
	return ""
}
