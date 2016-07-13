package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dyweb/Ayi/util"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update dependencies configured in .ayi.yml",
	Long:  "update required libraries and runtimes, auto detect composer.json package.json",
	Run: func(cmd *cobra.Command, args []string) {
		hasupdate := viper.IsSet("update")
		if !hasupdate {
			log.Warn("update configuration not found!")
			log.Errorf("TODO: looking for available commands")
			return
		}
		commands := viper.GetStringSlice("update")
		for _, cmd := range commands {
			log.Infof("executing: %s \n", cmd)
			util.RunCommand(cmd)
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)

}
