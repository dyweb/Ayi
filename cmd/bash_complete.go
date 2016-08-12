package cmd

import "github.com/spf13/cobra"

var bashCompleteCmd = &cobra.Command{
	Use:   "gen-bash",
	Short: "generate bash completion",
	Long:  "generate bash completion for Ayi's command",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: log
		// TODO: location
		// TODO: Add to ~/.bashrc
		// https://github.com/dyweb/Ayi/issues/45
		// TODO: when use on windows, bash_complition must be sourced
		RootCmd.GenBashCompletionFile("scripts/ayi_completion.sh")
	},
}

func init() {
	RootCmd.AddCommand(bashCompleteCmd)
}
