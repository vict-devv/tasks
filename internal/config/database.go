package config

import "github.com/spf13/viper"

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
	viper.SetDefault("DatabaseConfig.Host", "localhost")
	viper.SetDefault("DatabaseConfig.Port", 5432)
	viper.SetDefault("DatabaseConfig.User", "admin")
	viper.SetDefault("DatabaseConfig.Password", "123456")
	viper.SetDefault("DatabaseConfig.Name", "tasks_db")
	viper.SetDefault("DatabaseConfig.SSLMode", false)
}
