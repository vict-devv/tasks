package config

import (
	"github.com/spf13/viper"
	"github.com/vict-devv/tasks/internal/constants"
)

// Server holds the server-related configuration.
type Server struct {
	Host string `mapstructure:"host"`
	Port uint   `mapstructure:"port"`
}

func setServerConfigDefaults() {
	viper.SetDefault("ServerConfig.Host", constants.ConfigServerDefaultHost)
	viper.SetDefault("ServerConfig.Port", constants.ConfigServerDefaultPort)
}
