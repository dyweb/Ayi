package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/dyweb/Ayi/lib/config"
)

func main() {
	//	fmt.Println("Hello Mie!")
	//	config.PrintHosts(config.ParseHosts())
	app := cli.NewApp()
	app.Name = "Ayi"
	app.Usage = "Help you solve all the messy commands"
	app.Commands = []cli.Command{
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
			},
		},
	}
	app.Run(os.Args)
}