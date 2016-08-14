package util

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

// ViperReadAndMerge read specified config file and merge
func ViperReadAndMerge(path string) {
	// handle $HOME like viper does in `util` `absPathify`
	if strings.HasPrefix(path, "$HOME") {
		path = userHomeDir() + path[5:]
	}
	// TODO: handle other path problems
	path = filepath.FromSlash(path)
	if !FileExists(path) {
		log.WithField("file", path).Debug("Config file NOT found")
		return
	}
	viper.SetConfigFile(path)
	err := viper.MergeInConfig()
	if err != nil {
		log.WithField("file", path).Debug("Error merge config: " + err.Error())
		return
	}
	log.WithField("file", path).Debug("Config read and merged")
}

// copied from viper util since it's a private function
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
