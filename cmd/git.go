package cmd

import (
	"github.com/dyweb/Ayi/app/git"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "wrapper for git commands",
	Long:  "git expand url and e ... TODO",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Need to call ReadConfigFile manually, can't put it in init func
		git.ReadConfigFile()
		log.Debug(git.GetAllHosts())
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("I am the git command !")

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
