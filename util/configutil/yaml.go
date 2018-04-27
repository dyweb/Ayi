package configutil

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"

	"github.com/dyweb/gommon/errors"
)

// NOTE: not using gommon/config because it will change very soon
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
