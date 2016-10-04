package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/dyweb/Ayi/util/runner"
)

var depInstallCmd = &cobra.Command{
	Use:   "dep-install",
	Short: "install dependencies configured in .ayi.yml",
	Long:  "install required libraries and runtimes, auto detect composer.json package.json",
	Run: func(cmd *cobra.Command, args []string) {
		count, err := runner.ExecuteCommand("dep-install")
		if err != nil {
			log.Error(err.Error())
			log.Error("install failed")
			os.Exit(1)
		}
		log.Infof("All %d dependency install commands finished", count)
	},
}

func init() {
	RootCmd.AddCommand(depInstallCmd)

}
