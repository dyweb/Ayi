package git

import (
	"os"

	"github.com/dyweb/Ayi"
	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"
)

type App struct {
	root *cobra.Command

	log *dlog.Logger
}

func NewApp(r Ayi.Registry) (*App, error) {
	a := &App{}
	dlog.NewStructLogger(log, a)
	root := &cobra.Command{
		Use:   "git",
		Short: "git helper",
		Long:  "git helper long",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	root.AddCommand(a.cloneCommand())
	a.root = root
	return a, nil
}

func (a *App) CobraCommand() *cobra.Command {
	return a.root
}
