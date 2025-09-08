package config

import (
	"github.com/spf13/viper"
	"github.com/vict-devv/tasks/internal/constants"
)

// DatabaseConfig holds the database-related configuration.
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  bool   `mapstructure:"sslmode"`
}

func setDatabaseConfigDefaults() {
	viper.SetDefault("DatabaseConfig.Host", constants.ConfigDatabaseDefaultHost)
	viper.SetDefault("DatabaseConfig.Port", constants.ConfigDatabaseDefaultPort)
	viper.SetDefault("DatabaseConfig.User", constants.ConfigDatabaseDefaultUser)
	viper.SetDefault("DatabaseConfig.Password", constants.ConfigDatabaseDefaultPassword)
	viper.SetDefault("DatabaseConfig.Name", constants.ConfigDatabaseDefaultName)
	viper.SetDefault("DatabaseConfig.SSLMode", false)
}
