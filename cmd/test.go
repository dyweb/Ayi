package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dyweb/Ayi/util"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run test configed in .ayi.yml",
	Long:  "run user defined test commands in .ayi.yml",
	Run: func(cmd *cobra.Command, args []string) {
		hasTest := viper.IsSet("test")
		if !hasTest {
			// TODO: add color
			fmt.Println("Test not found!")
			return
		}
		commands := viper.GetStringSlice("test")
		for _, cmd := range commands {
			// TODO: color or put it in RunCommand, may need a log library
			fmt.Printf("executing: %s \n", cmd)
			util.RunCommand(cmd)
		}
	},
}

func init() {
	RootCmd.AddCommand(testCmd)

}
