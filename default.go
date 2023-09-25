package clog

import (
	"context"
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	defaultLogger Logger = NewLogrusWrapper(logrus.New())
	lock          sync.Mutex
)

// SetDefault sets the default logger. Using this package should be avoided as much as possible.
func SetDefault(logger Logger) error {
	if logger == nil {
		return errors.New("cannot set the logger to a nil value")
	}

	lock.Lock()
	defer lock.Unlock()

	defaultLogger = logger
	return nil
}

// Default returns the default logger.
func Default() Logger {
	return defaultLogger
}

func Panic(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Panic(values...)
}

func Panicln(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Panicln(values...)
}

func Panicf(ctx context.Context, pattern string, args ...interface{}) {
	LoggerFromContext(ctx).Panicf(pattern, args...)
}

func Fatal(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Fatal(values...)
}

func Fatalln(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Fatalln(values...)
}

func Fatalf(ctx context.Context, pattern string, args ...interface{}) {
	LoggerFromContext(ctx).Fatalf(pattern, args...)
}

func Error(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Error(values...)
}

func Errorln(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Errorln(values...)
}

func Errorf(ctx context.Context, pattern string, args ...interface{}) {
	LoggerFromContext(ctx).Errorf(pattern, args...)
}

func Warn(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Warn(values...)
}

func Warnln(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Warnln(values...)
}

func Warnf(ctx context.Context, pattern string, args ...interface{}) {
	LoggerFromContext(ctx).Warnf(pattern, args...)
}

func Info(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Info(values...)
}

func Infoln(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Infoln(values...)
}

func Infof(ctx context.Context, pattern string, args ...interface{}) {
	LoggerFromContext(ctx).Infof(pattern, args...)
}

func Debug(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Debug(values...)
}

func Debugln(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Debugln(values...)
}

func Debugf(ctx context.Context, pattern string, args ...interface{}) {
	LoggerFromContext(ctx).Debugf(pattern, args...)
}

func Trace(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Trace(values...)
}

func Traceln(ctx context.Context, values ...interface{}) {
	LoggerFromContext(ctx).Traceln(values...)
}

func Tracef(ctx context.Context, pattern string, args ...interface{}) {
	LoggerFromContext(ctx).Tracef(pattern, args...)
}

func WithField(ctx context.Context, key string, value interface{}) Logger {
	return LoggerFromContext(ctx).WithField(key, value)
}

func WithFields(ctx context.Context, fields Fields) Logger {
	return LoggerFromContext(ctx).WithFields(fields)
}
