package configutil

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/fsutil"
)

// NOTE: not using gommon/config because it will change very soon
func LoadYAMLFile(path string, cfg interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "can't open file")
	}
	return LoadYAML(f, cfg)
}

func LoadYAML(r io.Reader, cfg interface{}) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "can't drain reader")
	}
	if err := yaml.UnmarshalStrict(b, cfg); err != nil {
		return errors.Wrap(err, "can't parse yaml")
	}
	return nil
}

func SaveYAML(cfg interface{}) ([]byte, error) {
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}
	return b, nil
}

func SaveYAMLFile(path string, cfg interface{}) error {
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}
	return fsutil.WriteFile(path, b)
}
