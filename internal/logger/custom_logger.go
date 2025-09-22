package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/vict-devv/tasks/internal/constants"
)

// CustomLogger is a more configurable logger that allows specifying the handler type,
// handler options, and output destination.
type CustomLogger struct {
	handler        slog.Handler
	handlerOptions *slog.HandlerOptions
	output         io.WriteCloser
	logger         *slog.Logger
}

// Debug implements Logger.
// It invokes the Degug method of embedded slog.Logger, passing the package name as an attribute.
func (c *CustomLogger) Debug(pkg, msg string) {
	c.logger.With(constants.LogAttributePackage, pkg).Debug(msg)
}

// Debugf implements Logger.
// Same as Debug but for formatted messages.
func (c *CustomLogger) Debugf(pkg, format string, args ...any) {
	c.logger.With(constants.LogAttributePackage, pkg).Debug(fmt.Sprintf(format, args...))
}

// Error implements Logger.
// It invokes the Error method of embedded slog.Logger, passing the package name as an attribute.
func (c *CustomLogger) Error(pkg, msg string) {
	c.logger.With(constants.LogAttributePackage, pkg).Error(msg)
}

// ErrorF implements Logger.
// Same as Error but for formatted messages.
func (c *CustomLogger) ErrorF(pkg, format string, args ...any) {
	c.logger.With(constants.LogAttributePackage, pkg).Error(fmt.Sprintf(format, args...))
}

// Info implements Logger.
// It invokes the Info method of embedded slog.Logger, passing the package name as an attribute.
func (c *CustomLogger) Info(pkg, msg string) {
	c.logger.With(constants.LogAttributePackage, pkg).Info(msg)
}

// Infof implements Logger.
// Same as Info but for formatted messages.
func (c *CustomLogger) Infof(pkg, format string, args ...any) {
	c.logger.With(constants.LogAttributePackage, pkg).Info(fmt.Sprintf(format, args...))
}

// Warn implements Logger.
// It invokes the Warn method of embedded slog.Logger, passing the package name as an attribute.
func (c *CustomLogger) Warn(pkg, msg string) {
	c.logger.With(constants.LogAttributePackage, pkg).Warn(msg)
}

// Warnf implements Logger.
// Same as Warn but for formatted messages.
func (c *CustomLogger) Warnf(pkg, format string, args ...any) {
	c.logger.With(constants.LogAttributePackage, pkg).Warn(fmt.Sprintf(format, args...))
}

// Fatal implements Logger.
// It logs the message at error level and then exits the application.
func (c *CustomLogger) Fatal(pkg, msg string) {
	c.Error(pkg, msg)
	os.Exit(1)
}

// Fatalf implements Logger.
// It logs the formatted message at error level and then exits the application.
func (c *CustomLogger) Fatalf(pkg, format string, args ...any) {
	c.ErrorF(pkg, format, args...)
	os.Exit(1)
}

// NewCustomLogger creates a new instance of CustomLogger with the specified handler type,
// handler options, and output destination.
func NewCustomLogger(
	handlerType HandlerType,
	handlerOpts *slog.HandlerOptions,
	output io.WriteCloser,
) Logger {
	var handler slog.Handler

	switch handlerType {
	case TextHandler:
		handler = slog.NewTextHandler(output, handlerOpts)
	case JSONHandler:
		handler = slog.NewJSONHandler(output, handlerOpts)
	default:
		handler = slog.NewTextHandler(output, handlerOpts)
	}

	return &CustomLogger{
		handler:        handler,
		handlerOptions: handlerOpts,
		output:         output,
		logger:         slog.New(handler),
	}
}
