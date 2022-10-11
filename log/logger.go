// Package log provides logging functionality.
package log

import (
	"fmt"

	"github.com/cornelk/gotokit/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger provides fast, leveled, structured logging. All methods are safe
// for concurrent use.
// This logger library wrapper has an optimized handling of Debug() calls
// for loggers that have debug logging disabled.
type Logger struct {
	*zap.Logger
	level zap.AtomicLevel
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
	if cfg.Level == (zap.AtomicLevel{}) {
		level := DefaultLevel()
		cfg.Level = zap.NewAtomicLevelAt(level)
	}

	l, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("building logger: %w", err)
	}

	logger := &Logger{
		Logger: l,
		level:  cfg.Level,
	}
	return logger, nil
}

// Named adds a new path segment to the logger's name. Segments are joined by
// periods. By default, Loggers are unnamed.
func (l *Logger) Named(name string) *Logger {
	newLogger := l.Logger.Named(name)
	return &Logger{
		Logger: newLogger,
		level:  l.level,
	}
}

// With creates a child logger and adds structured context to it. Fields added
// to the child don't affect the parent, and vice versa.
func (l *Logger) With(fields ...Field) *Logger {
	newLogger := l.Logger.With(fields...)
	return &Logger{
		Logger: newLogger,
		level:  l.level,
	}
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// This function optimizes the check for the Debug level of the logger
// and avoid an unnecessary call to time.Now() as well as an allocation
// of zapcore.Entry in the case that the debug level is not set.
func (l *Logger) Debug(msg string, fields ...Field) {
	if l.level.Level() != DebugLevel {
		return
	}

	l.Logger.Debug(msg, fields...)
}

// Check returns a CheckedEntry if logging a message at the specified level
// is enabled. It's a completely optional optimization; in high-performance
// applications, Check can help avoid allocating a slice to hold fields.
func (l *Logger) Check(level zapcore.Level, msg string) *zapcore.CheckedEntry {
	if level == zap.DebugLevel && l.level.Level() != DebugLevel {
		return nil
	}

	return l.Logger.Check(level, msg)
}
