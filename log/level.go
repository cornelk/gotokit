package log

import (
	"sync/atomic"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log levels.
const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zap.DebugLevel

	// InfoLevel is the default logging priority.
	InfoLevel = zap.InfoLevel

	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zap.WarnLevel

	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zap.ErrorLevel

	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zap.FatalLevel
)

// Level is a logging priority. Higher levels are more important.
type Level = zapcore.Level

var defaultLevel = uintptr(zap.InfoLevel)

// DefaultLevel returns the current default level for all loggers
// newly created with New().
func DefaultLevel() Level {
	return Level(atomic.LoadUintptr(&defaultLevel))
}

// SetDefaultLevel sets the default level for all newly created loggers.
func SetDefaultLevel(level Level) {
	atomic.StoreUintptr(&defaultLevel, uintptr(level))
}
