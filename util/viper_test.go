package util

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViperGetStringOrFail(t *testing.T) {
	assert := assert.New(t)
	viper.SetConfigFile("../.ayi.example.yml")
	err := viper.ReadInConfig()
	assert.Nil(err)
	s, err := ViperGetStringOrFail("scripts.mie")
	assert.Equal("echo mie", s)
	_, err = ViperGetStringOrFail("scripts.qian")
	assert.NotNil(err)
	ss := viper.GetStringSlice("scripts.qian")
	assert.Equal(2, len(ss))
}
