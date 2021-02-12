package config

import (
	"path/filepath"
	"sync"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/wesovilabs/orion/internal/errors"
	"github.com/wesovilabs/orion/internal/logger"
)

const (
	baseDir = ".orion"
)

var (
	rootDir, _  = homedir.Dir()
	DefaultPath = filepath.Join(rootDir, baseDir, "orion.yml")
	v           = viper.New()
	cfg         *Config
	once        sync.Once
)

type Config struct {
	Logger *logger.Config `yaml:"logger"`
}

/**
func Load(path string) *Config {
	once.Do(func() {
		v.SetConfigFile(path)
		if err := read(); err == nil {
			log.Debugf("configuration file in '%s' was read successfully!", path)
		} else {
			if path != DefaultPath {
				log.Errorf("it failed reading config file: %s", err)
				os.Exit(1)
			}
			cfg.Logger = logger.Default()
		}

	})
	return cfg
}**/

func Get() *Config {
	once.Do(func() {
		cfg = new(Config)
		cfg.Logger = logger.Default()
	})
	return cfg
}

func read() errors.Error {
	cfg = &Config{}
	if err := v.ReadInConfig(); err != nil {
		return errors.InvalidArguments(err.Error())
	}
	if err := v.Unmarshal(cfg); err != nil {
		return errors.InvalidArguments(err.Error())
	}
	return nil
}
