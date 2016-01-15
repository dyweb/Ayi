package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/dyweb/Ayi/app"
	"github.com/dyweb/Ayi/sys"
	"github.com/dyweb/Ayi/util"
)

func main() {
	application := cli.NewApp()
	application.Name = "Ayi"
	application.Usage = "Let Ayi do it for you"
	application.Commands = []cli.Command{
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
		util.ServeStaticCommand,
		sys.HostCommands,
		app.GitCommands,
	}
	application.Run(os.Args)
}
