package app

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/dyweb/Ayi/app/git"
	"github.com/dyweb/Ayi/app/mail"
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

// MailCommands wrap mail service
var MailCommands = cli.Command{
	Name:    "mail",
	Aliases: []string{"m"},
	Usage:   "send mail to all web stuff",
	Subcommands: []cli.Command{
		{
			Name:    "send",
			Aliases: []string{"s"},
			Usage:   "git status",
			Action:  mail.SendMailToWebStuff,
		},
	},
}
