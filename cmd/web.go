package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dyweb/Ayi/app/web"
)

var (
	serverPort int
	serverRoot string
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "web start differen web servers",
	Long:  "web can start static server, the main Ayi server etc",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: start the default ayi server
	},
}

var webStaticCmd = &cobra.Command{
	Use:   "static",
	Short: "serve static content in current/specified folder",
	// FIXME: half of the long usage is not implemented ...
	Long: "serve static file like python's SimpleHTTPServer, support highlight and markdown render inspired by https://github.com/at15/doc-viewer",
	Run: func(cmd *cobra.Command, args []string) {
		if serverRoot == "" {
			serverRoot, _ = os.Getwd()
		}
		log.Infof("Use static server for %s on port %d", serverRoot, serverPort)
		server := web.NewStaticServer(serverRoot, serverPort)
		server.Run()
	},
}

func bindWebCmdFlagsToViper() {
	// TODO: is viper case sensitive?
	viper.BindPFlag("port", webCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("root", webCmd.PersistentFlags().Lookup("root"))
}

func init() {
	webCmd.AddCommand(webStaticCmd)
	RootCmd.AddCommand(webCmd)

	webCmd.PersistentFlags().IntVarP(&serverPort, "port", "p", 3000, "port to listen on")
	webCmd.PersistentFlags().StringVar(&serverRoot, "root", "", "server root folder")

	bindWebCmdFlagsToViper()
}
