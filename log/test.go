package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

// TestingT is a subset of the API provided by all *testing.T and
// *testing.B objects.
type TestingT interface {
	zaptest.TestingT

	Helper()
}

// NewTestLogger builds a new Logger that logs all messages to the given
// testing.TB. The logs get only printed if a test fails or if the test
// is run with -v verbose flag.
func NewTestLogger(t TestingT) *Logger {
	t.Helper()
	logger := zaptest.NewLogger(t)
	return &Logger{
		Logger: logger,
		level:  zap.NewAtomicLevelAt(DebugLevel),
	}
}
