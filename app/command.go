package app

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/dyweb/Ayi/app/git"
)

// GitCommands wrap common git operations with config defined in Ayi
var GitCommands = cli.Command{
	Name:    "git",
	Aliases: []string{"g"},
	Usage:   "git command wrapper",
	Subcommands: []cli.Command{
		{
			Name:    "status",
			Aliases: []string{"s"},
			Usage:   "git status",
			Action: func(c *cli.Context) {
				out, err, e := git.Status{}.Execute()
				if e != nil {
					println(err)
					log.Fatal(e)
				}
				println(out)
			},
		},
	},
}
