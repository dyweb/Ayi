package util

import (
	"os"

	"github.com/spf13/viper"
)

// https://github.com/dyweb/Ayi/issues/53
// Command runner behave like nodejs's command runner, execute command in config file
// support the following features
// - pre and post hook
// - built-in command has same behavior as user defined commands
// - support both one command and an array of commands
// - support -n
// - * support run in parallel
// - * support timeout

type Runner struct {
}

// TODO: may move it to packages like common
// TODO: may use map
// var BuiltInCommands = [...]string{
// 	"test",
// 	"install",
// 	"dep-install",
// }

var BuiltInCommands = map[string]bool{
	"test":        true,
	"install":     true,
	"dep-install": true,
}

func (runner *Runner) execute(cmdName string) error {
	var commands []string
	// FIXME: cannot use map[string]bool as type map[string]interface {}
	// if HasKey(BuiltInCommands, cmdName) {
	// 	commands = viper.GetStringSlice(cmdName)
	// }
	_, isBuiltIn := BuiltInCommands[cmdName]
	// TODO: support both single string and array
	// TODO: may put runner in a nested package add method called lookup command
	if isBuiltIn {
		commands = viper.GetStringSlice(cmdName)
	} else {
		commands = viper.GetStringSlice("scripts." + cmdName)
	}
	for _, cmd := range commands {
		log.Infof("executing: %s \n", cmd)
		err := RunCommand(cmd)
		if err != nil {
			log.Errorf("%s failed due to: %s", cmdName, err.Error())
			os.Exit(1)
		}
	}
	return nil
}
