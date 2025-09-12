package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/vict-devv/tasks/internal/constants"
)

// StandardLogger is a simple wrapper around slog.Logger for standard logging needs.
type StandardLogger struct {
	*slog.Logger
}

// Debug implements Logger.
// It invokes the Degug method of embedded slog.Logger, passing the package name as an attribute.
func (s *StandardLogger) Debug(pkg, msg string) {
	s.Logger.With(constants.LogAttributePackage, pkg).Debug(msg)
}

// Debugf implements Logger.
// Same as Debug but for formatted messages.
func (s *StandardLogger) Debugf(pkg, format string, args ...any) {
	s.Logger.With(constants.LogAttributePackage, pkg).Debug(fmt.Sprintf(format, args...))
}

// Error implements Logger.
// It invokes the Error method of embedded slog.Logger, passing the package name as an attribute.
func (s *StandardLogger) Error(pkg, msg string) {
	s.Logger.With(constants.LogAttributePackage, pkg).Error(msg)
}

// ErrorF implements Logger.
// Same as Error but for formatted messages.
func (s *StandardLogger) ErrorF(pkg, format string, args ...any) {
	s.Logger.With(constants.LogAttributePackage, pkg).Error(fmt.Sprintf(format, args...))
}

// Info implements Logger.
// It invokes the Info method of embedded slog.Logger, passing the package name as an attribute.
func (s *StandardLogger) Info(pkg, msg string) {
	s.Logger.With(constants.LogAttributePackage, pkg).Info(msg)
}

// Infof implements Logger.
// Same as Info but for formatted messages.
func (s *StandardLogger) Infof(pkg, format string, args ...any) {
	s.Logger.With(constants.LogAttributePackage, pkg).Info(fmt.Sprintf(format, args...))
}

// Warn implements Logger.
// It invokes the Warn method of embedded slog.Logger, passing the package name as an attribute.
func (s *StandardLogger) Warn(pkg, msg string) {
	s.Logger.With(constants.LogAttributePackage, pkg).Warn(msg)
}

// Warnf implements Logger.
// Same as Warn but for formatted messages.
func (s *StandardLogger) Warnf(pkg, format string, args ...any) {
	s.Logger.With(constants.LogAttributePackage, pkg).Warn(fmt.Sprintf(format, args...))
}

// Fatal implements Logger.
// It logs the message at error level and then exits the application.
func (s *StandardLogger) Fatal(pkg, msg string) {
	s.Error(pkg, msg)
	os.Exit(1)
}

// Fatalf implements Logger.
// It logs the formatted message at error level and then exits the application.
func (s *StandardLogger) Fatalf(pkg, format string, args ...any) {
	s.ErrorF(pkg, format, args...)
	os.Exit(1)
}

// NewStandardLogger creates a new instance of StandardLogger with default settings.
// It uses a text handler that outputs to standard output.
func NewStandardLogger() Logger {
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	return &StandardLogger{slog.New(h)}
}
