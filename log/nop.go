package log

import "go.uber.org/zap"

// NewNop creates a no-op logger which never writes logs to the output.
// Useful for tests.
func NewNop() *Logger {
	logger := zap.NewNop()
	return &Logger{
		Logger: logger,
		level:  zap.NewAtomicLevelAt(ErrorLevel),
	}
}
