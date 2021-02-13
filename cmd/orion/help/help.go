/*
Copyright 2021 The orion Authors.
*/
package help

import (
	"strings"

	"github.com/spf13/cobra"
)

var (
	use       = "help [command]"
	helpShort = `Help about any command.`
	helpLong  = `
		Help provides help for any command in the application.
		Simply type orion help [path to command] for full details.
	`
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Short:   helpShort,
		Long:    helpLong,
		Example: "help run",
		Run:     run,
	}
}

// RunHelp checks given arguments and executes command.
func run(cmd *cobra.Command, args []string) {
	foundCmd, _, err := cmd.Root().Find(args)
	switch {
	case foundCmd == nil:
		cmd.Printf("Unknown help topic %#q.\n", args)
		if usageErr := cmd.Root().Usage(); usageErr != nil {
			panic(usageErr)
		}
		return
	case err != nil:
		cmd.Println(err)
		argsString := strings.Join(args, " ")
		matchedMsgIsPrinted := false
		for _, foundCmd := range foundCmd.Commands() {
			if strings.Contains(foundCmd.Short, argsString) {
				if !matchedMsgIsPrinted {
					cmd.Printf("Matchers of string '%s' in short descriptions of commands: \n", argsString)
					matchedMsgIsPrinted = true
				}
				cmd.Printf("  %-14s %s\n", foundCmd.Name(), foundCmd.Short)
			}
		}
		if !matchedMsgIsPrinted {
			if err := cmd.Root().Usage(); err != nil {
				panic(err)
			}
		}
		return
	default:
		if len(args) == 0 {
			foundCmd = cmd
		}
		helpFunc := foundCmd.HelpFunc()
		helpFunc(foundCmd, args)
	}
}
