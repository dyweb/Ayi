package web

import (
	"fmt"
	"net/http"
	"os"

	iconfig "github.com/at15/go.ice/ice/config"
	ihttp "github.com/at15/go.ice/ice/transport/http"
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
	root := &cobra.Command{
		Use:   "web",
		Short: "web helper",
		Long:  "web helper long",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	var port int
	root.PersistentFlags().IntVar(&port, "port", 3000, "port listen to")
	static := &cobra.Command{
		Use:   "static",
		Short: "serve static content",
		Long:  "Serve static content",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: go.ice could have a static http server to be used out of box
			cfg := iconfig.HttpServerConfig{
				Addr:          fmt.Sprintf(":%d", port),
				EnableTracing: false,
			}
			h := http.FileServer(http.Dir("."))
			srv, err := ihttp.NewServer(cfg, h, nil)
			if err != nil {
				a.log.Fatal("failed to create http server")
				return
			}
			if err := srv.Run(); err != nil {
				a.log.Fatal(err)
				return
			}
		},
	}
	root.AddCommand(static)
	a.root = root
	return a, nil
}

func (a *App) CobraCommand() *cobra.Command {
	return a.root
}
