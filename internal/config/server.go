package config

import "github.com/spf13/viper"

// ServerConfig holds the server-related configuration.
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port uint   `mapstructure:"port"`
}

func setServerConfigDefaults() {
	viper.SetDefault("ServerConfig.Host", "localhost")
	viper.SetDefault("ServerConfig.Port", 8080)
}
