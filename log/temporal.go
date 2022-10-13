package log

import (
	"fmt"

	"go.uber.org/zap"
)

// TemporalLogger implements a logger compatible with the temporal package.
type TemporalLogger struct {
	logger *Logger
}

// NewTemporalLogger returns a new temporal logger based on the given logger.
func NewTemporalLogger(logger *Logger) *TemporalLogger {
	return &TemporalLogger{
		logger: logger,
	}
}

// Debug logs a message at DebugLevel.
func (l *TemporalLogger) Debug(msg string, keyValues ...any) {
	if l.logger.Level() > DebugLevel {
		return
	}

	fields := keyValuesToFields(keyValues...)
	l.logger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel.
func (l *TemporalLogger) Info(msg string, keyValues ...any) {
	if l.logger.Level() > InfoLevel {
		return
	}

	fields := keyValuesToFields(keyValues...)
	l.logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel.
func (l *TemporalLogger) Warn(msg string, keyValues ...any) {
	if l.logger.Level() > WarnLevel {
		return
	}

	fields := keyValuesToFields(keyValues...)
	l.logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel.
func (l *TemporalLogger) Error(msg string, keyValues ...any) {
	if l.logger.Level() > ErrorLevel {
		return
	}

	fields := keyValuesToFields(keyValues...)
	l.logger.Error(msg, fields...)
}

func keyValuesToFields(keyValues ...any) []Field {
	l := len(keyValues)
	fields := make([]Field, 0, (len(keyValues)/2)+1)

	for i := 0; i < l; i++ {
		k := keyValues[i]
		var key string
		switch t := k.(type) {
		case string:
			key = t
		default:
			key = fmt.Sprint(t)
		}

		i++
		var field Field
		if i < l {
			v := keyValues[i]
			field = zap.Any(key, v)
		} else {
			field = zap.String(key, "")
		}

		fields = append(fields, field)
	}
	return fields
}
