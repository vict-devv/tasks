// Package logger provides logging functionalities for the application.
package logger

// Logger defines the interface for logging within the application.
type Logger interface {
	// Debug logs a message at debug level.
	Debug(pkg, msg string)
	// Debugf logs a formatted message at debug level.
	Debugf(pkg, format string, args ...any)
	// Info logs a message at info level.
	Info(pkg, msg string)
	// Infof logs a formatted message at info level.
	Infof(pkg, format string, args ...any)
	// Warn logs a message at warn level.
	Warn(pkg, msg string)
	// Warnf logs a formatted message at warn level.
	Warnf(pkg, format string, args ...any)
	// Error logs a message at error level.
	Error(pkg, msg string)
	// Errorf logs a formatted message at error level.
	ErrorF(pkg, format string, args ...any)
	// Fatal logs a message at fatal level and then exits the application.
	Fatal(pkg, msg string)
	// Fatalf logs a formatted message at fatal level and then exits the application.
	Fatalf(pkg, format string, args ...any)
}

// HandlerType represents the type of log handler to use in CustomLogger.
type HandlerType int

const (
	// TextHandler represents a text log handler.
	TextHandler HandlerType = iota
	// JSONHandler represents a JSON log handler.
	JSONHandler
)
