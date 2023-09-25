package clog

import (
	"github.com/sirupsen/logrus"
)

type logrusEntryWrapper struct {
	entry *logrus.Entry
}

type LogrusWrapper struct {
	logger *logrus.Logger
}

func NewLogrusWrapper(logger *logrus.Logger) LogrusWrapper {
	return LogrusWrapper{logger: logger}
}

// Trace logs a message at level Trace on the standard logger.
func (l logrusEntryWrapper) Trace(args ...interface{}) {
	l.entry.Trace(args...)
}

// Trace logs a message at level Trace on the standard logger.
func (l LogrusWrapper) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

// Traceln logs a message at level Traceln on the standard logger.
func (l logrusEntryWrapper) Traceln(args ...interface{}) {
	l.entry.Traceln(args...)
}

// Traceln logs a message at level Traceln on the standard logger.
func (l LogrusWrapper) Traceln(args ...interface{}) {
	l.logger.Traceln(args...)
}

// Tracef logs a message at level Tracef on the standard logger.
func (l logrusEntryWrapper) Tracef(format string, args ...interface{}) {
	l.entry.Tracef(format, args...)
}

// Tracef logs a message at level Tracef on the standard logger.
func (l LogrusWrapper) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

// Debug logs a message at level Debug on the standard logger.
func (l logrusEntryWrapper) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

// Debug logs a message at level Debug on the standard logger.
func (l LogrusWrapper) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Debugln logs a message at level Debugln on the standard logger.
func (l logrusEntryWrapper) Debugln(args ...interface{}) {
	l.entry.Debugln(args...)
}

// Debugln logs a message at level Debugln on the standard logger.
func (l LogrusWrapper) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

// Debugf logs a message at level Debugf on the standard logger.
func (l logrusEntryWrapper) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// Debugf logs a message at level Debugf on the standard logger.
func (l LogrusWrapper) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func (l logrusEntryWrapper) Info(args ...interface{}) {
	l.entry.Info(args...)
}

// Info logs a message at level Info on the standard logger.
func (l LogrusWrapper) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Infoln logs a message at level Infoln on the standard logger.
func (l logrusEntryWrapper) Infoln(args ...interface{}) {
	l.entry.Infoln(args...)
}

// Infoln logs a message at level Infoln on the standard logger.
func (l LogrusWrapper) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

// Infof logs a message at level Infof on the standard logger.
func (l logrusEntryWrapper) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

// Infof logs a message at level Infof on the standard logger.
func (l LogrusWrapper) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l logrusEntryWrapper) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l LogrusWrapper) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Warnln logs a message at level Warnln on the standard logger.
func (l logrusEntryWrapper) Warnln(args ...interface{}) {
	l.entry.Warnln(args...)
}

// Warnln logs a message at level Warnln on the standard logger.
func (l LogrusWrapper) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

// Warnf logs a message at level Warnf on the standard logger.
func (l logrusEntryWrapper) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Warnf logs a message at level Warnf on the standard logger.
func (l LogrusWrapper) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func (l logrusEntryWrapper) Error(args ...interface{}) {
	l.entry.Error(args...)
}

// Error logs a message at level Error on the standard logger.
func (l LogrusWrapper) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Errorln logs a message at level Errorln on the standard logger.
func (l logrusEntryWrapper) Errorln(args ...interface{}) {
	l.entry.Errorln(args...)
}

// Errorln logs a message at level Errorln on the standard logger.
func (l LogrusWrapper) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

// Errorf logs a message at level Errorf on the standard logger.
func (l logrusEntryWrapper) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// Errorf logs a message at level Errorf on the standard logger.
func (l LogrusWrapper) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (l logrusEntryWrapper) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (l LogrusWrapper) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Fatalln logs a message at level Fatalln on the standard logger.
func (l logrusEntryWrapper) Fatalln(args ...interface{}) {
	l.entry.Fatalln(args...)
}

// Fatalln logs a message at level Fatalln on the standard logger.
func (l LogrusWrapper) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

// Fatalf logs a message at level Fatalf on the standard logger.
func (l logrusEntryWrapper) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

// Fatalf logs a message at level Fatalf on the standard logger.
func (l LogrusWrapper) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Panic logs a message at level Panic on the standard logger.
func (l logrusEntryWrapper) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

// Panic logs a message at level Panic on the standard logger.
func (l LogrusWrapper) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func (l logrusEntryWrapper) Panicln(args ...interface{}) {
	l.entry.Panicln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func (l LogrusWrapper) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (l logrusEntryWrapper) Panicf(format string, args ...interface{}) {
	l.entry.Panicf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (l LogrusWrapper) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func levelToLogrusLevel(level Level) logrus.Level {
	switch level {
	case LevelPanic:
		return logrus.PanicLevel
	case LevelFatal:
		return logrus.FatalLevel
	case LevelError:
		return logrus.ErrorLevel
	case LevelWarn:
		return logrus.WarnLevel
	case LevelInfo:
		return logrus.InfoLevel
	case LevelDebug:
		return logrus.DebugLevel
	case LevelTrace:
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}

// SetLevel sets the logger level.
func (l logrusEntryWrapper) SetLevel(level Level) {
	l.entry.Logger.Level = levelToLogrusLevel(level)
}

// SetLevel sets the logger level.
func (l LogrusWrapper) SetLevel(level Level) {
	l.logger.Level = levelToLogrusLevel(level)
}

// WithField adds field to the logger.
func (l logrusEntryWrapper) WithField(key string, value interface{}) Logger {
	return logrusEntryWrapper{entry: l.entry.WithField(key, value)}
}

// WithField adds field to the logger.
func (l LogrusWrapper) WithField(key string, value interface{}) Logger {
	return &logrusEntryWrapper{entry: l.logger.WithField(key, value)}
}

// WithFields adds fields to the logger.
func (l logrusEntryWrapper) WithFields(fields Fields) Logger {
	return logrusEntryWrapper{entry: l.entry.WithFields(logrus.Fields(fields))}
}

// WithFields adds fields to the logger.
func (l LogrusWrapper) WithFields(fields Fields) Logger {
	return &logrusEntryWrapper{entry: l.logger.WithFields(logrus.Fields(fields))}
}

// SetFormatter sets a custom logrus formatter.
func (l logrusEntryWrapper) SetFormatter(formatter logrus.Formatter) {
	l.entry.Logger.SetFormatter(formatter)
}

// SetFormatter sets a custom logrus formatter.
func (l LogrusWrapper) SetFormatter(formatter logrus.Formatter) {
	l.logger.SetFormatter(formatter)
}
