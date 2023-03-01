package log

import (
	"context"
	"fmt"

	"golang.org/x/exp/slog"
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
	l.logger.Log(context.TODO(), DebugLevel, msg, fields...)
}

// Info logs a message at InfoLevel.
func (l *TemporalLogger) Info(msg string, keyValues ...any) {
	if l.logger.Level() > InfoLevel {
		return
	}

	fields := keyValuesToFields(keyValues...)
	l.logger.Log(context.TODO(), InfoLevel, msg, fields...)
}

// Warn logs a message at WarnLevel.
func (l *TemporalLogger) Warn(msg string, keyValues ...any) {
	if l.logger.Level() > WarnLevel {
		return
	}

	fields := keyValuesToFields(keyValues...)
	l.logger.Log(context.TODO(), WarnLevel, msg, fields...)
}

// Error logs a message at ErrorLevel.
func (l *TemporalLogger) Error(msg string, keyValues ...any) {
	if l.logger.Level() > ErrorLevel {
		return
	}

	fields := keyValuesToFields(keyValues...)
	l.logger.Log(context.TODO(), ErrorLevel, msg, fields...)
}

func keyValuesToFields(keyValues ...any) []any {
	l := len(keyValues)
	fields := make([]any, 0, (len(keyValues)/2)+1)

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
			field = slog.Any(key, v)
		} else {
			field = slog.String(key, "")
		}

		fields = append(fields, field)
	}
	return fields
}
