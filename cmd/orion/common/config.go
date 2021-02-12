package common

import (
	"io"

	"github.com/spf13/viper"
	"github.com/wesovilabs-tools/orion/internal/config"
	"github.com/wesovilabs-tools/orion/internal/logger"
)

var cfg = &config.Config{
	Logger: logger.Default(),
}

func SetUpConfig(stdOut io.Writer) {
	viper.AutomaticEnv()
	logger.SetUp(cfg.Logger, stdOut)
}
