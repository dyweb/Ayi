package cmd

import (
	"github.com/dyweb/Ayi/app/git"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "wrapper for git commands",
	Long:  "git expand url and e ... TODO",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("I am the git command !")
		log.Info(git.DefaultHosts)
		// Need to call ReadConfigFile manually
		// TODO: how to avoid call this in every subcommand
		git.ReadConfigFile()
		log.Info(git.GetAllHosts())
	},
}

var gitCloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone a repo to workspace",
	Long:  "clone a repository to your workspace, short and browser urls are supported",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("I am the git clone command !")
	},
}

func init() {
	gitCmd.AddCommand(gitCloneCmd)
	RootCmd.AddCommand(gitCmd)
}
