package database

import (
	"context"
	"log/slog"

	"github.com/cornelk/gotokit/log"
	"github.com/jackc/pgx/v5/tracelog"
)

// Logger implements a database logger.
type Logger struct {
	logger *log.Logger
}

// NewLogger returns a new database logger based on the given logger.
func NewLogger(logger *log.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Log a message from the database package.
func (l *Logger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	if level == tracelog.LogLevelNone {
		return
	}

	fields := make([]log.Field, 0, len(data))
	for k, v := range data {
		fields = append(fields, slog.Any(k, v))
	}

	switch level {
	case tracelog.LogLevelTrace, tracelog.LogLevelDebug:
		l.logger.Log(ctx, log.DebugLevel, msg, fields...)

	case tracelog.LogLevelInfo:
		l.logger.Log(ctx, log.InfoLevel, msg, fields...)

	case tracelog.LogLevelWarn:
		l.logger.Log(ctx, log.WarnLevel, msg, fields...)

	case tracelog.LogLevelError:
		l.logger.Log(ctx, log.ErrorLevel, msg, fields...)
	}
}

// Level returns the minimum enabled log level.
func (l *Logger) Level() tracelog.LogLevel {
	level := l.logger.Level()

	switch level {
	case log.DebugLevel:
		return tracelog.LogLevelDebug

	case log.InfoLevel:
		return tracelog.LogLevelInfo

	case log.WarnLevel:
		return tracelog.LogLevelWarn

	case log.ErrorLevel, log.FatalLevel:
		return tracelog.LogLevelError

	default:
		return tracelog.LogLevelNone
	}
}
