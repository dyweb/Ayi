package util

import (
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

// DummyCommand for test if the package is working
var DummyCommand = cli.Command{
	Name:  "util-dummy",
	Usage: "dummy util command",
	Action: func(c *cli.Context) {
		println("util command is working!")
	},
}

// ServeStaticCommand serve static file in current folder
var ServeStaticCommand = cli.Command{
	Name:  "static",
	Usage: "serve static files",
	Flags: []cli.Flag{
		// TODO: add flag for folder
		cli.IntFlag{
			Name:  "port",
			Value: 8888,
			Usage: "the port to listen on",
		},
	},
	Action: func(c *cli.Context) {
		viper.SetDefault("base", ".")
		viper.Set("port", c.Int("port"))
		ServeStatic()
	},
}
