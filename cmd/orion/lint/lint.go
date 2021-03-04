package lint

import (
	"fmt"
	"os"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/spf13/cobra"
	"github.com/wesovilabs/orion/cmd/orion/common"
	"github.com/wesovilabs/orion/executor"
)

const (
	flagInputPath = "input"
	defInputPath  = "feature.hcl"
)

var (
	use       = "lint"
	helpShort = `static analysis of feature definition.`
	helpLong  = `
		Verify the content of the input file
	`
	example = `
		# Execute the scenario in file feature.hcl
		orion lint --input feature.hcl
	`
	inputPath string
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
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	ct.ChangeColor(ct.Green, false, ct.None, false)
	inputPath := cmd.Flag(flagInputPath)
	exec := executor.New()
	if err := exec.SetUp(inputPath.Value.String()); err != nil {
		common.PrintError(cmd, err)
		os.Exit(err.ExitStatus())
	}
}
