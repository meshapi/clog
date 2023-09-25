package clog

import "context"

type activeLoggerContextType struct{}

var activeLoggerContextKey = activeLoggerContextType{}

// ContextWithLogger creates a new context with the given logger.
func ContextWithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, activeLoggerContextKey, logger)
}

// LoggerFromContext returns active logger for the given ctx or the default logger if none is specified in the context.
func LoggerFromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(activeLoggerContextKey).(Logger); ok {
		return logger
	}

	return defaultLogger
}
