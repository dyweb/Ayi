package github

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dyweb/Ayi/util/configutil"
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/fsutil"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

const (
	labelDelimiter = "/"
)

type LabelConfig struct {
	Name  string
	Desc  string
	Color string
	Sub   []LabelConfig
}

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
			printLabels(a.listLabels(args))
		},
	}
	var output string
	save := &cobra.Command{
		Use:   "save",
		Short: "save labels",
		Long:  "Save labels of a repository",
		Run: func(cmd *cobra.Command, args []string) {
			labels := a.listLabels(args)
			b, err := configutil.SaveYAML(convertLabels(labels))
			if err != nil {
				log.Fatal(err)
				return
			}
			// FIXME: should store ownerRepo in App struct instead of using args[0] directly
			// TODO: might add repo to labels config instead of just a comment
			prefix := fmt.Sprintf("# saved from %s to %s\n", args[0], output)
			b = append([]byte(prefix), b...)
			if err := fsutil.WriteFile(output, b); err != nil {
				log.Fatal(err)
				return
			} else {
				fmt.Fprintln(os.Stdout, output)
			}
		},
	}
	save.Flags().StringVar(&output, "output", "labels.yml", "saved file")
	root.AddCommand(list, save)
	return root
}

func (a *App) listLabels(args []string) []*github.Label {
	if len(args) < 1 {
		a.log.Fatal("must pass at least one repository")
		return nil
	}
	ownerRepo := args[0]
	labels, err := ListLabels(a.c, ownerRepo)
	if err != nil {
		// FIXME: the logger registry seems to be incorrect, this fatal log does not have color
		a.log.Fatalf("failed to list labels for %s: %s", ownerRepo, err)
		return nil
	}
	return labels
}

func ListLabels(c *github.Client, ownerRepo string) ([]*github.Label, error) {
	segments := strings.Split(ownerRepo, "/")
	if len(segments) < 2 {
		return nil, errors.Errorf("invalid repo %s only has %d segments", ownerRepo, len(segments))
	}
	owner, repo := segments[0], segments[1]
	labels, _, err := c.Issues.ListLabels(context.Background(), owner, repo, nil)
	if err != nil {
		return nil, err
	}
	return labels, nil
}

func printLabels(labels []*github.Label) {
	for _, label := range labels {
		if label.Description != nil {
			fmt.Printf("%s desc %s color %s\n", *label.Name, *label.Description, *label.Color)
		} else {
			fmt.Printf("%s color %s\n", *label.Name, *label.Color)
		}
	}
}

// TODO: keep the hierarchy
func convertLabels(labels []*github.Label) []LabelConfig {
	cfgs := make([]LabelConfig, 0, len(labels))
	for _, label := range labels {
		l := LabelConfig{
			Name:  *label.Name,
			Color: *label.Color,
		}
		if label.Description != nil {
			l.Desc = *label.Description
		}
		cfgs = append(cfgs, l)
	}
	return cfgs
}

func ReadLabelConfig(file string) ([]LabelConfig, error) {
	labels := make([]LabelConfig, 0, 10)
	err := configutil.LoadYAMLFile(file, &labels)
	return labels, err
}

func FlattenLabelConfigs(cfgs []LabelConfig) []github.Label {
	labels := make([]github.Label, 0, len(cfgs))
	for _, cfg := range cfgs {
		labels = append(labels, flattenLabelConfig("", cfg)...)
	}
	return labels
}

func flattenLabelConfig(prefix string, cfg LabelConfig) []github.Label {
	if len(cfg.Sub) == 0 {
		label := github.Label{}
		name := prefix + cfg.Name
		label.Name = &name
		if cfg.Desc != "" {
			label.Description = &cfg.Desc
		}
		// TODO: should allow using parent color
		if cfg.Color != "" {
			label.Color = &cfg.Color
		}
		return []github.Label{label}
	} else {
		labels := make([]github.Label, 0, len(cfg.Sub))
		for _, subCfg := range cfg.Sub {
			labels = append(labels, flattenLabelConfig(prefix+cfg.Name+labelDelimiter, subCfg)...)
		}
		return labels
	}
}
