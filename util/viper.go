package util

// monkey patches for viper

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
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
		log.Debugf("Config file NOT found in %s", path)
		return
	}
	viper.SetConfigFile(path)
	err := viper.MergeInConfig()
	if err != nil {
		log.Debugf("Error merge config for %s: %s", path, err.Error())
		return
	}
	log.Debugf("Config read and merged for %s", path)
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

// ViperGetStringOrFail returns string without any convertion
// FIXED: https://github.com/dyweb/Ayi/issues/54
func ViperGetStringOrFail(key string) (string, error) {
	v := viper.Get(key)
	s, ok := v.(string)
	if ok {
		return s, nil
	}
	return "", errors.Errorf("%v is not a string", v)
}
