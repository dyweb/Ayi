package git

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dyweb/gommon/errors"
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
		Use:   "git",
		Short: "git helper",
		Long:  "git helper long",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	var useSsh bool
	clone := &cobra.Command{
		Use:   "clone",
		Short: "auto deduct clone url",
		Long:  "Detect repository from url and convert to a cloneable url",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				a.log.Fatal("must provide at least one url")
				return
			}
			err := a.clone(args, useSsh)
			if err != nil {
				a.log.Fatal(err)
			}
		},
	}
	clone.Flags().BoolVar(&useSsh, "ssh", true, "use ssh instead of http")
	root.AddCommand(clone)
	a.root = root
	return a, nil
}

func (a *App) CobraCommand() *cobra.Command {
	return a.root
}

func (a *App) clone(urls []string, ssh bool) error {
	merr := errors.NewMultiErr()
	for _, u := range urls {
		repo, err := UrlToRepo(u)
		if err != nil {
			a.log.Warnf("invalid repo url %s", err)
			merr.Append(errors.Wrap(err, "invalid repo url"))
			continue
		}
		var (
			src string
			dst string
		)
		if ssh {
			src = repo.SshUrl()
		} else {
			src = repo.HttpUrl()
		}
		if DefaultWorkspace() != "" {
			dst = filepath.Join(DefaultWorkspace(), repo.Host, repo.Owner, repo.Repository)
		} else {
			dst = ""
		}
		// TODO: clone result and keep track of cloned repo
		cmd := fmt.Sprintf("clone %s %s", src, dst)
		a.log.Infof("execute git %s", cmd)
		if err := RunCommand(cmd); err != nil {
			merr.Append(err)
			a.log.Error(err)
		} else {
			a.log.Infof("success git %s", cmd)
		}
	}
	return merr.ErrorOrNil()
}
