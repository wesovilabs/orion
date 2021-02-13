package common

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wesovilabs/orion/internal/config"
	"github.com/wesovilabs/orion/internal/logger"
)

const (
	flagVerbose = "verbose"
	flagConfig  = "config"
)

var (
	verbose bool
	cfgPath string
)

func SetCommonFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVar(
		&verbose,
		flagVerbose,
		false,
		"it displays info about time & durations execution (default is false)",
	)
	cmd.PersistentFlags().StringVar(
		&cfgPath,
		flagConfig,
		config.DefaultPath,
		fmt.Sprintf("config file (default is %s)", config.DefaultPath))
}

func PreRun(cmd *cobra.Command, args []string) {
	viper.AutomaticEnv()
	// cfg := config.Load(cfgPath)
	cfg := &config.Config{
		Logger: logger.Default(),
	}
	logger.SetUp(cfg.Logger, cmd.OutOrStdout())
}
