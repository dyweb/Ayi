package git

//import "github.com/spf13/viper"

type Host struct {
	URL          string
	SupportHTTPS bool
	SupportSSH   bool
	HTTPPort     int
	SSHPort      int
}

var DefaultHosts = [...]Host{
	Host{URL: "github.com", SupportHTTPS: true, SupportSSH: true},
	Host{URL: "gitlab.com", SupportHTTPS: true, SupportSSH: true},
	Host{URL: "bitbucket.org", SupportHTTPS: true, SupportSSH: true},
	Host{URL: "coding.net", SupportHTTPS: true, SupportSSH: true},
	Host{URL: "git.oschina.net", SupportHTTPS: true, SupportSSH: true},
}

func init() {
	// TODO: read from viper
	//hosts := viper.Get("git.hosts")
	//for _, host := range in hosts {
	//
	//}
}
