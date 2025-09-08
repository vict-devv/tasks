// Package config provides configuration management for the application.
package config

import "github.com/spf13/viper"

// Config holds the entire configuration for the application.
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

// New initializes and returns a new Config instance.
func New() *Config {
	cfg := &Config{}
	setConfigDefault()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./internal/config")

	// TODO: implement Logger and logging the ReadInConfig error
	_ = viper.ReadInConfig()

	if err := viper.Unmarshal(cfg); err != nil {
		panic("Failed to unmarshal config: " + err.Error())
	}

	return cfg
}

func setConfigDefault() {
	setServerConfigDefaults()
	setDatabaseConfigDefaults()
}
