package github

import (
	"github.com/dyweb/Ayi"
	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"
	"os"
)

type App struct {
	root *cobra.Command
	r    Ayi.Registry

	log *dlog.Logger
}

func NewApp(r Ayi.Registry) (*App, error) {
	a := &App{
		r: r,
	}
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
	label := &cobra.Command{
		Use:   "label",
		Short: "manage label",
		Long:  "Manage GitHub issue labels",
		Run: func(cmd *cobra.Command, args []string) {
			a.r.HasHomeConfig()
			a.r.HomeConfig()
			a.log.Infof("you token is %s", a.r.HomeConfig().GitHub.Token)
		},
	}
	root.AddCommand(label)
	a.root = root
	return a, nil
}

func (a *App) CobraCommand() *cobra.Command {
	return a.root
}
