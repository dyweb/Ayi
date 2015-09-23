package main

import (
	"os"
	"github.com/codegangsta/cli"
//	"github.com/dyweb/Ayi/lib/configs"
)

func main() {
	//	fmt.Println("Hello Mie!")
	//	configs.PrintHosts(configs.ParseHosts())
	app := cli.NewApp()
	app.Name = "Ayi"
	app.Usage = "Help you solve all the messy commands"
	app.Action = func(c *cli.Context) {
		println("Hello friend, do you want a roast mie?")
	}
	app.Run(os.Args)
}