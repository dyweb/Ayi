package util

import (
	"github.com/codegangsta/cli"
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
	Action: func(c *cli.Context) {
		ServeStatic(".", 8888)
	},
}
