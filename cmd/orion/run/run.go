package run

import (
	"fmt"
	"os"

	"github.com/wesovilabs-tools/orion/cmd/orion/common"
	"github.com/wesovilabs-tools/orion/executor"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/spf13/cobra"
	"github.com/wesovilabs-tools/orion/internal/config"
	"github.com/wesovilabs-tools/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
)

const (
	flagInputPath     = "input"
	flagVariablesPath = "vars"
	flagVerbose       = "verbose"
	defInputPath      = "feature.hcl"
)

var (
	use       = "run"
	helpShort = `Execute feature.`
	helpLong  = `
		Execute the scenarios in the provided input file.
	`
	example = `
		# Execute the scenario in file feature.hcl
		orion run --feature feature.hcl
	`
	inputPath     string
	variablesPath string
	verbose       string
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     use,
		Short:   helpShort,
		Long:    helpLong,
		Example: example,
		Run:     run,
		PreRun:  common.PreRun,
	}
	cmd.PersistentFlags().StringVar(
		&inputPath,
		flagInputPath,
		defInputPath,
		fmt.Sprintf("path of the input file (default is %s)", defInputPath))
	cmd.PersistentFlags().StringVar(
		&variablesPath,
		flagVariablesPath,
		"",
		"path of the variables file")
	cmd.PersistentFlags().StringVar(
		&verbose,
		flagVerbose,
		"INFO",
		fmt.Sprintf("log level. Supported values are:  'TRACE', 'DEBUG','INFO','WARN','ERROR'"))
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	ct.ChangeColor(ct.Green, false, ct.None, false)

	verboseMode := cmd.Flag(flagVerbose)
	if verboseMode.Value.String() != "" {
		config.Get().Logger.SetLevel(verboseMode.Value.String())
	}
	common.SetUpConfig(os.Stdout)

	inputPath := cmd.Flag(flagInputPath)
	exec := executor.New()
	if err := exec.SetUp(inputPath.Value.String()); err != nil {
		common.PrintError(cmd, err)
		os.Exit(err.ExitStatus())
	}
	variablesPath := cmd.Flag(flagVariablesPath)
	variables := make(map[string]cty.Value)
	var err errors.Error
	if variablesPath.Value.String() != "" {
		variables, err = executor.ParseVariables(variablesPath.Value.String())
		if err != nil {
			common.PrintError(cmd, err)
			os.Exit(err.ExitStatus())
		}
	}

	if err := exec.Run(variables); err != nil {
		common.PrintError(cmd, err)
		os.Exit(err.ExitStatus())
	}
}
