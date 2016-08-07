package git

import (
	"github.com/dyweb/Ayi/util"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Host struct {
	URL          string
	SupportHTTPS bool
	SupportSSH   bool
	HTTPPort     int
	SSHPort      int
	// TODO: add type, github or gitlab in order to use api client
}

// DefaultHosts include common public git hosts
// TODO: maybe use an string array to construct this is better
var DefaultHosts = [...]Host{
	*NewHost("github.com"),
	*NewHost("gitlab.com"),
	*NewHost("bitbucket.org"),
	*NewHost("coding.net"),
	*NewHost("git.oschina.net"),
}

var hosts []Host
var log = util.Logger

// ReadConfigFile read user defined hosts in .ayi.yml
func ReadConfigFile() {
	log.Debug("Read git section in config file")
	hostsNameMap := make(map[string]bool)
	hostsSlice := cast.ToSlice(viper.Get("git.hosts"))
	for _, h := range hostsSlice {
		m := cast.ToStringMap(h)
		name := cast.ToString(m["name"])
		// TODO: check if attributes exist and give default value
		https := cast.ToBool(m["https"])
		// TODO: more attributes
		h := NewHost(name)
		h.SupportHTTPS = https
		hosts = append(hosts, *h)
		hostsNameMap[name] = true
	}
	// only merge default hosts that are not configed in config files
	for _, defaultHost := range DefaultHosts {
		_, exists := hostsNameMap[defaultHost.URL]
		if !exists {
			hosts = append(hosts, defaultHost)
		}
	}
}

// GetAllHosts return hard coded hosts and user defined hosts
func GetAllHosts() []Host {
	return hosts
}

// NewHost return a new Host object with default config
// https://golang.org/doc/effective_go.html#allocation_new
// TODO: why use poniter instead of return object directly, The Go Programming Language P32 2.3.2 pointer
func NewHost(url string) *Host {
	return &Host{URL: url, SupportHTTPS: true, SupportSSH: true, HTTPPort: 80, SSHPort: 25}
}
