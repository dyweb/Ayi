package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/dyweb/Ayi/util/runner"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run test configured in .ayi.yml",
	Long:  "run user defined test commands in .ayi.yml",
	Run: func(cmd *cobra.Command, args []string) {
		count, err := runner.ExecuteCommand("test")
		if err != nil {
			log.Error(err.Error())
			log.Error("test failed")
			os.Exit(1)
		}
		log.Infof("All %d test commands have passed", count)
	},
}

func init() {
	RootCmd.AddCommand(testCmd)

}
