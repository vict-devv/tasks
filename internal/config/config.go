// Package config provides configuration management for the application.
package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/vict-devv/tasks/internal/constants"
	"github.com/vict-devv/tasks/internal/logger"
)

// Config holds the entire configuration for the application.
type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

// New initializes and returns a new Config instance.
func New() *Config {
	logger := logger.NewStandardLogger()
	logger.Debug(constants.LogPackageConfig, "Loading configuration")

	cfg := &Config{}
	setConfigDefault()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./internal/config")

	if err := viper.ReadInConfig(); err != nil {
		logger.
			ErrorF(constants.LogPackageConfig, "Failed to read config file: %v", err)
	}

	setAutomaticEnv()

	if err := viper.Unmarshal(cfg); err != nil {
		panic("Failed to unmarshal config: " + err.Error())
	}

	return cfg
}

func setConfigDefault() {
	setServerConfigDefaults()
	setDatabaseConfigDefaults()
}

func setAutomaticEnv() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(constants.ConfigEnvPrefix)
	viper.AutomaticEnv()
}
