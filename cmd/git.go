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
		git.ReadConfigFile()
		log.Info(git.GetAllHosts())
	},
}

func init() {
	RootCmd.AddCommand(gitCmd)
}
