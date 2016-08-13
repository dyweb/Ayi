package cmd

import (
	"github.com/spf13/cobra"

	"github.com/dyweb/Ayi/app/web"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "web start differen web servers",
	Long:  "web can start static server, the main Ayi server etc",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: specify port etc
		server := web.NewStaticServer()
		server.Run()
	},
}

func init() {
	RootCmd.AddCommand(webCmd)

}
