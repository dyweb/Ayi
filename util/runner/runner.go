package runner

import (
	"github.com/dyweb/Ayi/util"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var log = util.Logger

// https://github.com/dyweb/Ayi/issues/53
// Command runner behave like nodejs's command runner, execute command in config file
// support the following features
// - pre and post hook
// - built-in command has same behavior as user defined commands
// - support both one command and an array of commands
// - support -n
// - * support run in parallel
// - * support timeout

// TODO: may move it to packages like common
var BuiltInCommands = map[string]bool{
	"test":        true,
	"install":     true,
	"dep-install": true,
}

// LookUpCommands looks for single or array of commands in built-in and scripts blocks
// TODO: test
func LookUpCommands(cmdName string) ([]string, error) {
	commands := make([]string, 1)
	// FIXME: cannot use map[string]bool as type map[string]interface {}
	// TODO: maybe because bool is a primitive type?
	// if HasKey(BuiltInCommands, cmdName) {
	// 	commands = viper.GetStringSlice(cmdName)
	// }
	_, isBuiltIn := BuiltInCommands[cmdName]

	// FIXED: https://github.com/dyweb/Ayi/issues/54
	// TODO: need more test to test the behavior of command
	fullName := cmdName
	if !isBuiltIn {
		fullName = "scripts." + cmdName
	}

	// first try single string
	command, err := util.ViperGetStringOrFail(fullName)
	// TODO: maybe we should allow empty command
	if err == nil && command != "" {
		log.Debugf("command is %s", command)
		commands[0] = command
		return commands, nil
	}

	log.Debug("single string command not found, try array")

	commands = viper.GetStringSlice(fullName)
	if len(commands) == 0 {
		if isBuiltIn {
			return commands, errors.Errorf("%s configuration not found", cmdName)
		}
		return commands, errors.Errorf("command %s not found in scripts block", cmdName)

	}
	return commands, nil
}

// ExecuteCommand look for configured command(s) and execute
// TODO: return value can have name, if my memory is correct
func ExecuteCommand(cmdName string) (int, error) {
	commands, err := LookUpCommands(cmdName)
	if err != nil {
		log.Warnf(err.Error())
		return 0, errors.Wrap(err, "Runner can't find commands")
	}
	success := 0
	for _, cmd := range commands {
		log.Infof("executing: %s \n", cmd)
		err := util.RunCommand(cmd)

		if err != nil {
			_, ok := err.(*util.DryRunError)
			if !ok {
				return success, errors.Errorf("%s failed due to: %s", cmdName, err.Error())
			}
		}
		success++
	}
	// TODO: may need to return dry run error here
	return len(commands), nil
}

// TODO: try to log stdout to file does not work for
// - ls
// - git clone
// func runCommand(cmd string) error {
// 	command, err := util.Command(cmd)
// 	if err != nil {
// 		return errors.Wrap(err, "runner cannot recognize command")
// 	}
// 	// f, _ := os.Create("log.txt")
// 	// defer f.Close()
// 	// multiWriter := io.MultiWriter(os.Stdout, f)
// 	command.Stdin = os.Stdin
// 	command.Stdout = os.Stdout
// 	command.Stderr = os.Stderr
// 	err = command.Run()
// 	if err != nil {
// 		return errors.Wrap(err, "Failure when executing command")
// 	}
// 	return nil
// }