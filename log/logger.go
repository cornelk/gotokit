// Package log provides logging functionality.
package log

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/cornelk/gotokit/env"
	"golang.org/x/exp/slog"
)

// Logger provides fast, leveled, structured logging. All methods are safe
// for concurrent use.
type Logger struct {
	*slog.Logger
	level *slog.LevelVar
}

// New returns a new Logger instance.
func New() (*Logger, error) {
	var (
		cfg Config
		err error
	)
	level := DefaultLevel()
	if level == DebugLevel {
		cfg, err = ConfigForEnv(env.Development)
	} else {
		cfg, err = ConfigForEnv(env.Production)
	}
	if err != nil {
		return nil, fmt.Errorf("getting config for env: %w", err)
	}

	return NewWithConfig(cfg)
}

// NewWithConfig creates a new logger for the given config.
// If no level is set in the config, it will use the default level of
// this package.
func NewWithConfig(cfg Config) (*Logger, error) {
	level := &slog.LevelVar{}
	level.Set(cfg.Level)

	opts := &slog.HandlerOptions{
		AddSource: cfg.CallerInfo,
		Level:     level,
	}

	var output io.Writer
	if cfg.Output == nil {
		output = os.Stdout
	} else {
		output = cfg.Output
	}

	handler := cfg.Handler
	if handler == nil {
		if cfg.JSONOutput {
			handler = slog.NewJSONHandler(output, opts)
		} else {
			opts.ReplaceAttr = ReplaceLevelName
			consoleOpts := &ConsoleHandlerOptions{
				SlogOptions: opts,
				TimeFormat:  cfg.TimeFormat,
			}
			if cfg.TimeFormat == "" {
				consoleOpts.TimeFormat = DefaultTimeFormat
			}
			handler = NewConsoleHandler(output, consoleOpts)
		}
	}

	l := slog.New(handler)
	logger := &Logger{
		Logger: l,
		level:  level,
	}
	return logger, nil
}

// Named adds a new path segment to the logger's name. Segments are joined by
// periods. By default, Loggers are unnamed.
func (l *Logger) Named(name string) *Logger {
	newLogger := l.Logger.WithGroup(name)
	return &Logger{
		Logger: newLogger,
		level:  l.level,
	}
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func (l *Logger) With(fields ...any) *Logger {
	newLogger := l.Logger.With(fields...)
	return &Logger{
		Logger: newLogger,
		level:  l.level,
	}
}

// Level returns the minimum enabled log level.
func (l *Logger) Level() Level {
	return l.level.Level()
}

// SetLevel alters the logging level.
func (l *Logger) SetLevel(level Level) {
	l.level.Set(level)
}

// Trace logs at TraceLevel.
func (l *Logger) Trace(msg string, args ...any) {
	l.Log(context.Background(), TraceLevel, msg, args...)
}

// TraceContext logs at TraceLevel with the given context.
func (l *Logger) TraceContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, TraceLevel, msg, args...)
}

// Fatal logs at FatalLevel.
func (l *Logger) Fatal(msg string, args ...any) {
	l.Log(context.Background(), FatalLevel, msg, args...)
	fatalExitFunc()
}

// FatalContext logs at FatalLevel with the given context.
func (l *Logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, FatalLevel, msg, args...)
	fatalExitFunc()
}

// fatalExitFunc defines the function to call when exiting due to a fatal log error.
// This is used in unit tests.
var fatalExitFunc = fatalExit

func fatalExit() {
	os.Exit(1)
}
