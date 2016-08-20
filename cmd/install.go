package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dyweb/Ayi/util"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "build and install binary",
	Long:  "build and install binary following commands defined in install block in .ayi.yml",
	Run: func(cmd *cobra.Command, args []string) {
		hasInstall := viper.IsSet("install")
		if !hasInstall {
			log.Warn("Install configuration not found!")
			return
		}
		commands := viper.GetStringSlice("install")
		for _, cmd := range commands {
			log.Infof("executing: %s \n", cmd)
			util.RunCommand(cmd)
		}
	},
}

func init() {
	RootCmd.AddCommand(installCmd)

}
