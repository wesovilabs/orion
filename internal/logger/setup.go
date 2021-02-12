package logger

import (
	"io"

	log "github.com/sirupsen/logrus"
)

// SetUp configure logger
func SetUp(cfg *Config, stdOut io.Writer) {
	log.SetFormatter(cfg.formatter())
	log.SetLevel(cfg.level())
	log.SetOutput(stdOut)

}

// Default configure logger with default values
func Default() *Config {
	return defaultConfig
}
