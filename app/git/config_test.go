package git

// TODO: use sutie http://godoc.org/github.com/stretchr/testify/suite for test
// - [x] read config in setup
// - [x] test cast
// - [x] test if config is readed
// - [ ] test if default host is merged
// - [x] test if default value for host exists

import (
	"testing"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GitConfigTestSuite struct {
	suite.Suite
}

func (suite *GitConfigTestSuite) SetupTest() {
	assert := assert.New(suite.T())
	viper.SetConfigFile("../../.ayi.example.yml")
	err := viper.ReadInConfig()
	assert.Nil(err)
	assert.Equal(true, viper.Get("debug"))
}

func (suite *GitConfigTestSuite) TestCast() {
	assert := assert.New(suite.T())
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

func (suite *GitConfigTestSuite) TestReadConfig() {
	assert := assert.New(suite.T())
	ReadConfigFile()
	hosts := GetAllHosts()
	hostsMap := GetAllHostsMap()
	assert.NotEqual(0, len(hosts))
	assert.Equal(len(hosts), len(hostsMap))

	// Skip host without name
	defaultHostsLength := len(DefaultHosts)
	// ayi.example.yml has three hosts and one does not have name
	assert.Equal(2+defaultHostsLength, len(hosts))

	// Default value is working
	assert.Equal(hostsMap["github.com"].SSHPort, DefaultSSHPort)
	assert.Equal(hostsMap["github.com"].HTTPPort, DefaultHTTPPort)

	// Config value is working
	assert.Equal(hostsMap["git.saber.io"].SSHPort, 10086)
}

func TestGitConfigTestSuite(t *testing.T) {
	suite.Run(t, new(GitConfigTestSuite))
}
