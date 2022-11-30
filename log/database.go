package log

import (
	"context"

	"github.com/jackc/pgx/v5/tracelog"
	"golang.org/x/exp/slog"
)

// DatabaseLogger implements a logger compatible with the database package.
type DatabaseLogger struct {
	logger *Logger
}

// NewDatabaseLogger returns a new database logger based on the given logger.
func NewDatabaseLogger(logger *Logger) *DatabaseLogger {
	return &DatabaseLogger{
		logger: logger,
	}
}

// Log a message from the database package.
func (l *DatabaseLogger) Log(_ context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	if level == tracelog.LogLevelNone {
		return
	}

	fields := make([]any, 0, len(data))
	for k, v := range data {
		fields = append(fields, slog.Any(k, v))
	}

	switch level {
	case tracelog.LogLevelTrace, tracelog.LogLevelDebug:
		l.logger.LogDepth(1, DebugLevel, msg, fields...)

	case tracelog.LogLevelInfo:
		l.logger.LogDepth(1, InfoLevel, msg, fields...)

	case tracelog.LogLevelWarn:
		l.logger.LogDepth(1, WarnLevel, msg, fields...)

	case tracelog.LogLevelError:
		l.logger.LogDepth(1, ErrorLevel, msg, fields...)
	}
}

// Level returns the minimum enabled log level.
func (l *DatabaseLogger) Level() tracelog.LogLevel {
	level := l.logger.Level()

	switch level {
	case DebugLevel:
		return tracelog.LogLevelDebug

	case InfoLevel:
		return tracelog.LogLevelInfo

	case WarnLevel:
		return tracelog.LogLevelWarn

	case ErrorLevel, FatalLevel:
		return tracelog.LogLevelError

	default:
		return tracelog.LogLevelNone
	}
}
