// Package constants defines application-wide constant values.
package constants

// Configuration-related constants.
const (
	ConfigEnvPrefix = "TASKS"
	// Server configuration defaults.
	ConfigServerDefaultHost = "localhost"
	ConfigServerDefaultPort = 8080
	// Database configuration defaults.
	ConfigDatabaseDefaultHost     = "localhost"
	ConfigDatabaseDefaultPort     = 5432
	ConfigDatabaseDefaultUser     = "admin"
	ConfigDatabaseDefaultPassword = "123456"
	ConfigDatabaseDefaultName     = "tasks_db"
)

// Logging-related constants.
const (
	LogAttributePackage = "package"
	// Packages names for logging.
	LogPackageMain   = "main"
	LogPackageConfig = "config"
)
