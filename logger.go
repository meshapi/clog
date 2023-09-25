package clog

// Fields is a map of keys to values that captures the fields a logger is carrying.
type Fields map[string]interface{}

// Logger describes ability to log application events.
type Logger interface {
	// Trace logs a message at level Trace on the standard logger.
	Trace(...interface{})
	// Traceln logs a message at level Trace on the standard logger.
	Traceln(...interface{})
	// Tracef logs a message at level Trace on the standard logger.
	Tracef(string, ...interface{})

	// Debug logs a message at level Debug on the standard logger.
	Debug(...interface{})
	// Debugln logs a message at level Debug on the standard logger.
	Debugln(...interface{})
	// Debugf logs a message at level Debug on the standard logger.
	Debugf(string, ...interface{})

	// Info logs a message at level Info on the standard logger.
	Info(...interface{})
	// Infoln logs a message at level Info on the standard logger.
	Infoln(...interface{})
	// Infof logs a message at level Info on the standard logger.
	Infof(string, ...interface{})

	// Warn logs a message at level Warn on the standard logger.
	Warn(...interface{})
	// Warnln logs a message at level Warn on the standard logger.
	Warnln(...interface{})
	// Warnf logs a message at level Warn on the standard logger.
	Warnf(string, ...interface{})

	// Error logs a message at level Error on the standard logger.
	Error(...interface{})
	// Errorln logs a message at level Error on the standard logger.
	Errorln(...interface{})
	// Errorf logs a message at level Error on the standard logger.
	Errorf(string, ...interface{})

	// Fatal logs a message at level Fatal on the standard logger and exit immediately.
	Fatal(...interface{})
	// Fatalln logs a message at level Fatal on the standard logger and exit immediately.
	Fatalln(...interface{})
	// Fatalf logs a message at level Fatal on the standard logger and exit immediately.
	Fatalf(string, ...interface{})

	// Panic logs a message at level Panic on the standard logger.
	Panic(...interface{})
	// Panicln logs a message at level Panic on the standard logger.
	Panicln(...interface{})
	// Panicf logs a message at level Panic on the standard logger.
	Panicf(string, ...interface{})

	// SetLevel sets the log level.
	SetLevel(Level)

	// WithField creates a new logger with the given key and value pair added, existing values get overwritten.
	WithField(key string, value interface{}) Logger

	// WithFields creates a new logger with with the given table of keys and values, existing values get overwritten.
	WithFields(Fields) Logger
}
