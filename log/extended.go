package log

import (
	"context"
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
func (l *ExtendedLogger) Debugf(format string, v ...interface{}) {
	if l.logger.Level() > DebugLevel {
		return
	}

	msg := fmt.Sprintf(format, v...)
	l.logger.Log(context.TODO(), DebugLevel, msg)
}

// Warnf logs a message at WarnLevel.
func (l *ExtendedLogger) Warnf(format string, v ...interface{}) {
	if l.logger.Level() > WarnLevel {
		return
	}

	msg := fmt.Sprintf(format, v...)
	l.logger.Log(context.TODO(), WarnLevel, msg)
}

// Errorf logs a message at ErrorLevel.
func (l *ExtendedLogger) Errorf(format string, v ...interface{}) {
	if l.logger.Level() > ErrorLevel {
		return
	}

	msg := fmt.Sprintf(format, v...)
	l.logger.Log(context.TODO(), ErrorLevel, msg)
}
