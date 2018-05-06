package git

import (
	"fmt"
	"path/filepath"

	"github.com/dyweb/gommon/errors"
	"github.com/spf13/cobra"
)

func (a *App) cloneCommand() *cobra.Command {
	var useSSH bool
	root := &cobra.Command{
		Use:   "clone",
		Short: "auto deduct clone url",
		Long:  "Detect repository from url and convert to a cloneable url",
		Example: `
ayi git clone dyweb/gommon
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				a.log.Fatal("must provide at least one url")
				return
			}
			err := a.clone(args, useSSH)
			if err != nil {
				a.log.Fatal(err)
			}
		},
	}
	root.Flags().BoolVar(&useSSH, "ssh", true, "use ssh instead of http")
	return root
}

// TODO: considering clone one repo and move the loop to caller?
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
