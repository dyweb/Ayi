// Package cmd defines and implements command-line commands and flags
// used by Ayi. Commands and flags are implemented using Cobra.
// it is generated by cobra https://github.com/spf13/cobra/tree/master/cobra
// and modified following https://github.com/spf13/hugo/blob/master/commands/hugo.go
package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dyweb/Ayi/util"
)

// Flags that are to be added to commands.
var cfgFile string
var (
	version bool
	verbose bool
	dryRun  bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "Ayi",
	Short: "Ayi makes your life easier",
	Long:  `Ayi is a collection of small applications and tools that speed up your develop process`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			versionCmd.Run(cmd, args)
			return
		}

		// FIXME: print the help here
		// FIXME: On Windows, it works in cmd, but does not work in Git Bash
		color.Green("Use 'Ayi help' to see all commands")
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		// TODO: use logger
		// https://github.com/spf13/cobra/issues/304
		// error message is printed twice for command not found
		// fmt.Println(err)
		os.Exit(-1)
	}
}

func loadDefaultSettings() {
	viper.SetDefault("Verbose", false)
	viper.SetDefault("DryRun", false)
}

func bindRootCmdFlagsToViper() {
	viper.BindPFlag("Verbose", RootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("DryRun", RootCmd.PersistentFlags().Lookup("dry-run"))
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ayi.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "n", false, "show commands to execute")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolVar(&version, "version", false, "show current version")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// https://github.com/spf13/viper#working-with-flags
	bindRootCmdFlagsToViper()

	// NOTE: code here does work, but the problem is, you can't get flag?
	var dummyCmd = &cobra.Command{
		Use:   "dummy",
		Short: "dummy is foo",
		Long:  `dummy is foo bar`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Hi I am dummy")
		},
	}
	RootCmd.AddCommand(dummyCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if verbose {
		util.UseVerboseLog()
	}

	viper.AutomaticEnv() // read in environment variables that match
	// TODO: what's the order of merge, and how does it handle array
	util.ViperReadAndMerge("$HOME/.ayi.yml")
	util.ViperReadAndMerge(".ayi.yml")
	util.ViperReadAndMerge(".ayi.local.yml")
	if cfgFile != "" { // enable ability to specify config file via flag
		util.ViperReadAndMerge(cfgFile)
	}

	// Set default value for viper
	loadDefaultSettings()

	// FIXME: this is just test if dynamic registering command is possible for cobra
	// TODO: code does not work here
	// var dummyCmd = &cobra.Command{
	// 	Use:   "dummy",
	// 	Short: "dummy is foo",
	// 	Long:  `dummy is foo bar`,
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		log.Info("Hi I am dummy")
	// 	},
	// }
	// RootCmd.AddCommand(dummyCmd)
}
