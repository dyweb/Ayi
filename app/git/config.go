package git

import (
	"github.com/dyweb/Ayi/util"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// Host represent a remote git host server like github.com, gitlab.com
type Host struct {
	URL          string
	SupportHTTPS bool
	SupportSSH   bool
	HTTPPort     int
	SSHPort      int
	SSHURL       string
	// TODO: add type, github or gitlab in order to use api client
	// TODO: add access token
}

const (
	// DefaultSSHPort 22
	DefaultSSHPort = 22
	// DefaultHTTPPort 80
	DefaultHTTPPort = 80
)

// DefaultHosts include common public git hosts
// TODO: maybe use an string array to construct this is better
var DefaultHosts = [...]Host{
	*NewHost("github.com"),
	*NewHost("gitlab.com"),
	*NewHost("bitbucket.org"),
	*NewHost("coding.net", "git.coding.net"),
	*NewHost("git.oschina.net"),
}

// host array keep the order of hosts
var hosts []Host

// host map is more convenient
var hostsMap map[string]Host

// ReadConfigFile read user defined hosts in .ayi.yml
func ReadConfigFile() {
	log.Debug("Read git section in config file")
	hostsMap = make(map[string]Host)
	hostsSlice := cast.ToSlice(viper.Get("git.hosts"))

	for _, h := range hostsSlice {
		m := cast.ToStringMap(h)
		_, exists := m["name"]
		if !exists {
			log.Warn("Skipp host without name")
			continue
		}

		// TODO: more attributes, the following is not working
		// - http port
		// - support ssh
		name := cast.ToString(m["name"])
		https := cast.ToBool(util.GetWithDefault(m, "https", true))
		port := cast.ToInt(util.GetWithDefault(m, "port", DefaultSSHPort))

		h := NewHost(name)
		h.SupportHTTPS = https
		h.SSHPort = port

		hosts = append(hosts, *h)
		// TODO: may add order to host
		hostsMap[name] = *h
	}

	// only merge default hosts that are not configed in config files
	for _, defaultHost := range DefaultHosts {
		_, exists := hostsMap[defaultHost.URL]
		if !exists {
			hosts = append(hosts, defaultHost)
			hostsMap[defaultHost.URL] = defaultHost
		}
	}
}

// GetAllHosts return an array of hosts
func GetAllHosts() []Host {
	return hosts
}

// GetAllHostsMap return a map with host name(url) as key
func GetAllHostsMap() map[string]Host {
	return hostsMap
}

// NewHost return a new Host object with default config
// https://golang.org/doc/effective_go.html#allocation_new
// TODO: why use poniter instead of return object directly, The Go Programming Language P32 2.3.2 pointer
func NewHost(urls ...string) *Host {
	h := Host{SupportHTTPS: true, SupportSSH: true, HTTPPort: DefaultHTTPPort, SSHPort: DefaultSSHPort}
	if len(urls) > 0 {
		h.URL = urls[0]
		h.SSHURL = h.URL
	}
	if len(urls) > 1 {
		h.SSHURL = urls[1]
	}
	return &h
}
