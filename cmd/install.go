package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/dyweb/Ayi/util/runner"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "build and install binary",
	Long:  "build and install binary following commands defined in install block in .ayi.yml",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := runner.ExecuteCommand("install")
		if err != nil {
			log.Error(err.Error())
			log.Error("install failed")
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(installCmd)

}
