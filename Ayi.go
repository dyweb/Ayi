package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/dyweb/Ayi/util"
)

func main() {
	app := cli.NewApp()
	app.Name = "Ayi"
	app.Usage = "Let Ayi do it for you"
	app.Commands = []cli.Command{
		// TODO: Move the command to different packages, not a good idea to have a big file here
		{
			// Ayi like roast mie
			Name:    "mie",
			Aliases: []string{"arrowrowe"},
			Usage:   "roast mie",
			Action: func(c *cli.Context) {
				println("roast mie")
			},
		},
		util.DummyCommand,
	}
	app.Run(os.Args)
}
