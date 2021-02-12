/*
Copyright 2021 The orion Authors.
*/
package help

import (
	"fmt"
	"io"
	"strings"

	ct "github.com/daviddengcn/go-colortext"
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

// RunHelp checks given arguments and executes command
func run(cmd *cobra.Command, args []string) {

	foundCmd, _, err := cmd.Root().Find(args)
	if foundCmd == nil {
		cmd.Printf("Unknown help topic %#q.\n", args)
		cmd.Root().Usage()
	} else if err != nil {
		// print error message at first, since it can contain suggestions
		cmd.Println(err)

		argsString := strings.Join(args, " ")
		var matchedMsgIsPrinted = false
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
			// if nothing is found, just print usage
			cmd.Root().Usage()
		}
	} else {
		if len(args) == 0 {
			// help message for help command :)
			foundCmd = cmd
		}
		helpFunc := foundCmd.HelpFunc()
		helpFunc(foundCmd, args)
	}
}

func printService(out io.Writer, name, link string) {
	ct.ChangeColor(ct.Green, false, ct.None, false)
	fmt.Fprint(out, name)
	ct.ResetColor()
	fmt.Fprint(out, " is running at ")
	ct.ChangeColor(ct.Yellow, false, ct.None, false)
	fmt.Fprint(out, link)
	ct.ResetColor()
	fmt.Fprintln(out, "")
}
