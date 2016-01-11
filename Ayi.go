package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/dyweb/Ayi/sys"
	"github.com/dyweb/Ayi/util"
)

func main() {
	app := cli.NewApp()
	app.Name = "Ayi"
	app.Usage = "Let Ayi do it for you"
	app.Commands = []cli.Command{
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
		sys.HostCommand,
	}
	app.Run(os.Args)
}
