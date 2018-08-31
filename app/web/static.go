package web

import (
	"fmt"
	"net/http"

	iconfig "github.com/dyweb/go.ice/ice/config"
	ihttp "github.com/dyweb/go.ice/ice/transport/http"
	"github.com/spf13/cobra"
)

func (a *App) staticCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "static",
		Short: "serve static content",
		Long:  "Serve static content",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: go.ice should have a static http server to be used out of box
			cfg := iconfig.HttpServerConfig{
				Addr:          fmt.Sprintf(":%d", a.port),
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
	return root
}
