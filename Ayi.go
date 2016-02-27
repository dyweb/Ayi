package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/dyweb/Ayi/app"
	"github.com/dyweb/Ayi/sys"
	"github.com/dyweb/Ayi/util"
	"github.com/spf13/viper"
)

const (
	configName = ".ayi"
	appName    = "Ayi"
)

func main() {
	// read the config from config file
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)                        // name of config file (without extension)
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", appName))  // path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", appName)) // call multiple times to add many search paths
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// alloc the cli
	application := cli.NewApp()
	application.Name = appName
	application.Usage = "Let Ayi do it for you"
	application.Commands = []cli.Command{
		{
			// Ayi like roast mie
			Name:    "mie",
			Aliases: []string{"arrowrowe"},
			Usage:   "roast mie",
			Action: func(c *cli.Context) {
				println("roast mie")
			},
		},
		util.DummyCommand,
		util.ServeStaticCommand,
		sys.HostCommands,
		app.GitCommands,
	}
	application.Run(os.Args)
}
