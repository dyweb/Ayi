package cmd

import (
	"fmt"
	"os"

	"github.com/dyweb/Ayi/util/runner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run user defined commands in scripts block",
	Long:  "run user defined commands in .ayi.yml's scripts block. ie: Ayi run build",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Error("Must specify command to run! ie: Ayi run build")
			listAllScripts()
			os.Exit(1)
		}
		hasScripts := viper.IsSet("scripts")
		if !hasScripts {
			log.Error("run configuration not found!")
			os.Exit(1)
		}
		scriptName := args[0]
		_, err := runner.ExecuteCommand(scriptName)
		if err != nil {
			log.Error(err.Error())
			log.Error("script failed")
			os.Exit(1)
		}
		log.Info("script finished")
	},
}

func listAllScripts() {
	scripts := viper.GetStringMapString("scripts")
	log.Info("all avaliable scripts are listed below")
	for name := range scripts {
		fmt.Println(name)
	}
}

func init() {
	RootCmd.AddCommand(runCmd)
}
