package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/go-errors/errors"
	"github.com/dyweb/Ayi/lib/config"
	"github.com/dyweb/Ayi/lib/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "Ayi"
	app.Usage = "Help you solve all the messy commands"
	app.Commands = []cli.Command{
		// TODO: Move the command to different packages, not a good idea to have a big file here
		// roast is not just for fun, it's for eating
		{
			Name: "roast",
			Aliases:[]string{"r", "rst"},
			Usage: "roast mie",
			Action: func(c *cli.Context) {
				println("roast mie")
			},
		},
		// config host
		{
			Name: "hosts",
			Aliases:[]string{"host"},
			// TODO: this should show help?
			// TODO: update the usage
			Usage:"config/show  host",
			Subcommands:[]cli.Command{
				{
					Name: "list",
					Aliases:[]string{"l"},
					Usage:"Show all current hosts in /etc/hosts, ipv4 only",
					Action:func(c *cli.Context) {
						config.PrintHosts(config.ParseHosts())
					},
				},
				{
					Name: "add",
					Aliases:[]string{"a"},
					Usage:"Add ip and domain to host",
					// TODO: share the flags with remove
					Flags:[]cli.Flag{
						cli.StringFlag{
							Name: "ip",
							Value: "127.0.0.1",
							Usage: "ip address to add to hosts file",
						},
						cli.StringFlag{
							Name: "name",
							Value: "localhost",
							Usage: "domain name to add to hosts file",
						},
					},
					Action:func(c *cli.Context) {
						//						println(c.String("ip"))
						//						println(c.String("name"))
						//						println(c.String("aaa") == "")
						added, err := config.AddDomainToIp(c.String("name"), c.String("ip"))
						if !added {
							println("Fail adding host: " + err.Error())
						}
						// no output if no error
					},
				},
				{
					Name: "remove",
					Aliases:[]string{"rm"},
					Usage: "Remove domain from host file",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "name",
							Usage: "domain name to remove from hosts file",
						},
					},
					Action:func(c *cli.Context) {
						if c.String("name") == "" {
							println("must speicify domain name to remove from host file")
						}
						removed, err := config.RemoveDomain(c.String("name"))
						if !removed {
							// TODO: only print the stack in debug mode
							println(err.(*errors.Error).ErrorStack())
						}
					},
				},
			},
		},
		{
			Name:"serve",
			Usage:"serve static files",
			Action:func(c *cli.Context) {
				server.Run("front/public_html", 8888)
			},
		},
	}
	app.Run(os.Args)
}