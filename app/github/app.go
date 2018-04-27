package github

import (
	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"
	"os"
)

type App struct {
	root *cobra.Command

	log *dlog.Logger
}

func NewApp() (*App, error) {
	a := &App{}
	dlog.NewStructLogger(log, a)
	root := &cobra.Command{
		Use:   "github",
		Short: "github util",
		Long:  "Manage GitHub issues",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	a.root = root
	return a, nil
}

func (a *App) CobraCommand() *cobra.Command {
	return a.root
}
