package git

import (
	"testing"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	assert := assert.New(t)
	// NOTE: the path is relative to the test file, not project root
	viper.SetConfigFile("../../.ayi.example.yml")
	err := viper.ReadInConfig()
	assert.Nil(err)
	assert.Equal(true, viper.Get("debug"))
	// NOTE: cast is needed
	hostsRaw := viper.Get("git.hosts")
	//fmt.Println(hostsRaw)
	hostsSlice := cast.ToSlice(hostsRaw)
	//fmt.Println(hostsSlice)

	for _, host := range hostsSlice {
		hostMap := cast.ToStringMap(host)
		name := cast.ToString(hostMap["name"])
		https := cast.ToBool(hostMap["https"])
		if name == "git.saber.io" {
			assert.Equal(false, https)
		}
	}
}
