package github

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/oauth2"

	"github.com/dyweb/Ayi"
	dlog "github.com/dyweb/gommon/log"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

type App struct {
	root *cobra.Command
	r    Ayi.Registry
	c    *github.Client

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
	root.AddCommand(a.labelCommand())
	a.root = root
	return a, nil
}

func (a *App) CobraCommand() *cobra.Command {
	return a.root
}

// TODO: move it to gommon as stringutil
func maskSensitive(s string) string {
	l := len(s)
	if l < 2 {
		return "*"
	}
	s = s[:2]
	for i := 0; i < l-2; i++ {
		s += "*"
	}
	return s
}

// TODO: move it to gommon requests
// NOTE: base on https://github.com/golang/oauth2/blob/master/oauth2.go#L314
// FIXME: should have a way to not use oauth2 package ...
func oauthHttpClient(token string) *http.Client {
	return oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
}

func (a *App) createClient() {
	if !a.r.HasHomeConfig() {
		a.log.Fatal("~/.ayi.yml not found, required for github token")
		return
	}
	token := a.r.HomeConfig().GitHub.Token
	if token == "" {
		a.log.Fatal("github token is empty")
		return
	}
	a.log.Debugf("you token is %s", maskSensitive(token))
	a.c = github.NewClient(oauthHttpClient(token))
}
