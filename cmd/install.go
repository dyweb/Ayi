package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dyweb/Ayi/util"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install dependencies configured in .ayi.yml",
	Long:  "install required libraries and runtimes, auto detect composer.json package.json",
	Run: func(cmd *cobra.Command, args []string) {
		hasInstall := viper.IsSet("install")
		if !hasInstall {
			// TODO: add color
			fmt.Println("Install configuration not found!")
			// TODO: try lookup composer.json package.json
			fmt.Println("TODO: looking for available commands")
			return
		}
		commands := viper.GetStringSlice("install")
		for _, cmd := range commands {
			// TODO: color or put it in RunCommand, may need a log library
			fmt.Printf("executing: %s \n", cmd)
			util.RunCommand(cmd)
		}
	},
}

func init() {
	RootCmd.AddCommand(installCmd)

}
