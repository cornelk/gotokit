package log

import (
	"fmt"
)

// ExtendedLogger implements a logger compatible with multiple common packages.
type ExtendedLogger struct {
	logger *Logger
}

// NewExtendedLogger returns a new extended logger based on the given logger.
func NewExtendedLogger(logger *Logger) *ExtendedLogger {
	return &ExtendedLogger{
		logger: logger,
	}
}

// Debugf logs a message at DebugLevel.
func (l *ExtendedLogger) Debugf(format string, v ...any) {
	if l.logger.Level() > DebugLevel {
		return
	}

	msg := fmt.Sprintf(format, v...)
	l.logger.Log(nil, DebugLevel, msg)
}

// Warnf logs a message at WarnLevel.
func (l *ExtendedLogger) Warnf(format string, v ...any) {
	if l.logger.Level() > WarnLevel {
		return
	}

	msg := fmt.Sprintf(format, v...)
	l.logger.Log(nil, WarnLevel, msg)
}

// Errorf logs a message at ErrorLevel.
func (l *ExtendedLogger) Errorf(format string, v ...any) {
	if l.logger.Level() > ErrorLevel {
		return
	}

	msg := fmt.Sprintf(format, v...)
	l.logger.Log(nil, ErrorLevel, msg)
}
