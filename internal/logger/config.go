package logger

import (
	"github.com/sirupsen/logrus"
)

const timestampFormat = "15:04:05.000000"

var defaultConfig = &Config{
	Level: logrus.TraceLevel.String(),
}

// Config contains configuration for logrus.
type Config struct {
	Level string `yaml:"level"`
}

func (c *Config) level() logrus.Level {
	lvl, err := logrus.ParseLevel(c.Level)
	if err == nil {
		return lvl
	}
	logrus.Warningf("unexpected error parsing log level: '%s'", err)
	return logrus.InfoLevel
}

func (c *Config) SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		c.Level = logrus.InfoLevel.String()
	}
	c.Level = lvl.String()
}

func (c *Config) formatter() logrus.Formatter {
	return &Formatter{
		TimestampFormat: timestampFormat,
		ColorDisabled:   false,
	}
}
