package cmd

import (
	"fmt"
	"os"

	"github.com/dyweb/Ayi/util"
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
		command := viper.GetString("scripts." + scriptName)
		if command == "" {
			log.Errorf("script %s not found!", scriptName)
			listAllScripts()
			os.Exit(1)
		}
		log.Infof("executing: %s \n", command)
		err := util.RunCommand(command)
		if err != nil {
			log.Errorf("%s failed due to: %s", scriptName, err.Error())
			os.Exit(1)
		}
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
