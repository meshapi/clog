package clog

import (
	"fmt"
	"strings"
)

// Level is the severity level of a log statement.
type Level uint8

const (
	// LevelPanic level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	LevelPanic Level = iota
	// LevelFatal level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	LevelFatal
	// LevelError level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	LevelError
	// LevelWarn level. Non-critical entries that deserve eyes.
	LevelWarn
	// LevelInfo level. General operational entries about what's going on inside the
	// application.
	LevelInfo
	// LevelDebug level. Usually only enabled when debugging. Very verbose logging.
	LevelDebug
	// LevelTrace level. Designates finer-grained informational events than the Debug.
	LevelTrace
)

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "panic":
		return LevelPanic, nil
	case "fatal":
		return LevelFatal, nil
	case "error":
		return LevelError, nil
	case "warn", "warning":
		return LevelWarn, nil
	case "info":
		return LevelInfo, nil
	case "debug":
		return LevelDebug, nil
	case "trace":
		return LevelTrace, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid clog Level: %q", level)
}
