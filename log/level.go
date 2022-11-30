package log

import (
	"sync/atomic"

	"golang.org/x/exp/slog"
)

// Log levels.
const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = slog.DebugLevel

	// InfoLevel is the default logging priority.
	InfoLevel = slog.InfoLevel

	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = slog.WarnLevel

	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = slog.ErrorLevel

	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = slog.ErrorLevel + 1
)

// Level is a logging priority. Higher levels are more important.
type Level = slog.Level

var defaultLevel = uintptr(slog.InfoLevel)

// DefaultLevel returns the current default level for all loggers
// newly created with New().
func DefaultLevel() Level {
	return Level(atomic.LoadUintptr(&defaultLevel))
}

// SetDefaultLevel sets the default level for all newly created loggers.
func SetDefaultLevel(level Level) {
	atomic.StoreUintptr(&defaultLevel, uintptr(level))
}
