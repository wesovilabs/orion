package main

import (
	"github.com/spf13/cobra"
	"github.com/wesovilabs/orion/cmd/orion/help"
	"github.com/wesovilabs/orion/cmd/orion/lint"
	"github.com/wesovilabs/orion/cmd/orion/run"
)

// const flagConfig  = "config"

var (
	// cfgPath string.
	helpCmd = help.New()
	runCmd  = run.New()
	lintCmd = lint.New()
)

func main() {
	cmd := command()
	// common.SetCommonFlags(cmd)
	// setUpConfig(cmd.OutOrStdout())
	cmd.SetHelpCommand(helpCmd)
	cmd.AddCommand(helpCmd, runCmd, lintCmd)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

func command() *cobra.Command {
	return &cobra.Command{
		Use: "orion [cmd]",
	}
}
