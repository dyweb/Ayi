package git

import (
	"os"

	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"
)

type App struct {
	root *cobra.Command

	log *dlog.Logger
}

func NewApp() (*App, error) {
	a := &App{}
	dlog.NewStructLogger(log, a)
	a.root = &cobra.Command{
		Use:   "git",
		Short: "git helper",
		Long:  "git helper",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	a.root.AddCommand(a.cloneCommand())
	return a, nil
}

func (a *App) CobraCommand() *cobra.Command {
	return a.root
}

func (a *App) cloneCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "clone",
		Short: "auto deduct clone url",
		Long:  "Detect repository from url and convert to a cloneable url",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				a.log.Fatal("must provide at least one url")
			}
			// TODO: clone one by one
			a.log.Info(args)
		},
	}
}
