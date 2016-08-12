package cmd

import "github.com/spf13/cobra"

// https://github.com/dyweb/Ayi/issues/45
var bashCompleteCmd = &cobra.Command{
	Use:   "bash-gen",
	Short: "generate bash completion",
	Long:  "generate bash completion for Ayi's command",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: log
		// TODO: allow specify generated file location
		// TODO: add to ~/.bashrc
		// TODO: when use on windows, bash_complition must be sourced
		RootCmd.GenBashCompletionFile("scripts/ayi_completion.sh")
	},
}

func init() {
	RootCmd.AddCommand(bashCompleteCmd)
}
