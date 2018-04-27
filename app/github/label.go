package github

import (
	"context"
	"os"
	"strings"

	"github.com/dyweb/gommon/errors"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

func (a *App) labelCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "label",
		Short: "manage labels",
		Long:  "Manage GitHub issue labels",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			a.createClient()
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	list := &cobra.Command{
		Use:   "list",
		Short: "list labels",
		Long:  "List labels of a repository",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				a.log.Fatal("must pass at least one repository")
				return
			}
			ownerRepo := args[0]
			labels, err := a.listLabels(ownerRepo)
			if err != nil {
				a.log.Fatalf("failed to list labels for %s: %s", ownerRepo, err)
				return
			}
			a.log.Infof("labels %s", labels)
		},
	}
	root.AddCommand(list)
	return root
}

func (a *App) listLabels(ownerRepo string) ([]*github.Label, error) {
	segments := strings.Split(ownerRepo, "/")
	if len(segments) < 2 {
		return nil, errors.Errorf("invalid repo %s only has %d segments", ownerRepo, len(segments))
	}
	owner, repo := segments[0], segments[1]
	labels, _, err := a.c.Issues.ListLabels(context.Background(), owner, repo, nil)
	if err != nil {
		return nil, err
	}
	return labels, nil
}
