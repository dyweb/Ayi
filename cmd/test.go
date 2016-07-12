package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dyweb/Ayi/util"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run test configured in .ayi.yml",
	Long:  "run user defined test commands in .ayi.yml",
	Run: func(cmd *cobra.Command, args []string) {
		hasTest := viper.IsSet("test")
		if !hasTest {
			log.Error("Test not found!")
			return
		}
		commands := viper.GetStringSlice("test")
		for _, cmd := range commands {
			log.Infof("executing: %s \n", cmd)
			util.RunCommand(cmd)
		}
	},
}

func init() {
	RootCmd.AddCommand(testCmd)

}
