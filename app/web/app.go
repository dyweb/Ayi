package web

import (
	"os"

	"github.com/dyweb/Ayi"
	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"
)

type App struct {
	root *cobra.Command

	// config
	port int

	log *dlog.Logger
}

func NewApp(r Ayi.Registry) (*App, error) {
	a := &App{}
	dlog.NewStructLogger(log, a)
	root := &cobra.Command{
		Use:   "web",
		Short: "web helper",
		Long:  "web helper long",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	root.PersistentFlags().IntVarP(&a.port, "port", "p", 3000, "port listen to")
	root.AddCommand(
		a.staticCommand(),
		a.sshdCommand(),
	)
	a.root = root
	return a, nil
}

func (a *App) CobraCommand() *cobra.Command {
	return a.root
}
