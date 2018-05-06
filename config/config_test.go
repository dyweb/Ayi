package config

import (
	"testing"

	"github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := asst.New(t)

	//var cfg *AyiConfig // NOTE: this will cause panic
	cfg := &AyiConfig{}
	testutil.ReadYAMLToStrict(t, "example.ayi.yml", cfg)
	assert.Equal("I shall not be committed", cfg.GitHub.Token)
}
