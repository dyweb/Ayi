package cmd

import (
	"os"

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
			err := util.RunCommand(cmd)
			if err != nil {
				log.Errorf("Test failed due to: %s", err.Error())
				os.Exit(1)
			}
		}
		log.Infof("All %d test commands have passed", len(commands))
	},
}

func init() {
	RootCmd.AddCommand(testCmd)

}
