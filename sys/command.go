package sys

import (
	"github.com/codegangsta/cli"
	"github.com/dyweb/Ayi/sys/net"
	"github.com/go-errors/errors"
)

// HostCommand for list and modify host file
var HostCommand = cli.Command{
	Name:    "hosts",
	Aliases: []string{"host"},
	Usage:   "config/show  host",
	Subcommands: []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Show all current hosts in /etc/hosts, ipv4 only",
			Action: func(c *cli.Context) {
				net.PrintHosts(net.ParseHosts())
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add ip and domain to host",
			// TODO: share the flags with remove
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ip",
					Value: "127.0.0.1",
					Usage: "ip address to add to hosts file",
				},
				cli.StringFlag{
					Name:  "name",
					Value: "localhost",
					Usage: "domain name to add to hosts file",
				},
			},
			Action: func(c *cli.Context) {
				//						println(c.String("ip"))
				//						println(c.String("name"))
				//						println(c.String("aaa") == "")
				added, err := net.AddDomainToIP(c.String("name"), c.String("ip"))
				if !added {
					println("Fail adding host: " + err.Error())
				}
				// no output if no error
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "Remove domain from host file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "domain name to remove from hosts file",
				},
			},
			Action: func(c *cli.Context) {
				if c.String("name") == "" {
					println("must speicify domain name to remove from host file")
				}
				removed, err := net.RemoveDomain(c.String("name"))
				if !removed {
					// TODO: only print the stack in debug mode
					println(err.(*errors.Error).ErrorStack())
				}
			},
		},
	},
}
