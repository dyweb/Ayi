package util

import (
	"github.com/codegangsta/cli"
)

var DummyCommand = cli.Command{
	Name:  "util-dummy",
	Usage: "dummy util command",
	Action: func(c *cli.Context) {
		println("util command is working!")
	},
}
