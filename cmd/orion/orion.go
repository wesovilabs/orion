package main

import (
	"github.com/spf13/cobra"
	"github.com/wesovilabs/orion/cmd/orion/help"
	"github.com/wesovilabs/orion/cmd/orion/lint"
	"github.com/wesovilabs/orion/cmd/orion/run"
)

const (
	flagVerbose = "verbose"
	flagConfig  = "config"
)

var (
	verbose bool
	cfgPath string
	cmd     = command()
	helpCmd = help.New()
	runCmd  = run.New()
	lintCmd = lint.New()
)

func main() {
	var cmd = command()
	//common.SetCommonFlags(cmd)
	//setUpConfig(cmd.OutOrStdout())
	cmd.SetHelpCommand(helpCmd)
	cmd.AddCommand(helpCmd, runCmd, lintCmd)
	cmd.Execute()
}

func command() *cobra.Command {
	return &cobra.Command{
		Use: "orion-cli [cmd]",
	}
}
