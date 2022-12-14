package log

import (
	"bytes"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slog"
)

// TestingT is a subset of the API provided by all *testing.T and
// *testing.B objects.
type TestingT interface {
	// Logf logs the given message without failing the test.
	Logf(string, ...interface{})

	// Errorf logs the given message and marks the test as failed.
	Errorf(string, ...interface{})

	// FailNow marks the test as failed and stops execution of that test.
	FailNow()

	// Helper marks the calling function as a test helper function.
	Helper()
}

// NewTestLogger builds a new Logger that logs all messages to the given
// testing.TB. The logs get only printed if a test fails or if the test
// is run with -v verbose flag.
func NewTestLogger(t TestingT) *Logger {
	t.Helper()

	handler := newTestHandler(t)
	cfg := Config{
		JSONOutput: false,
		CallerInfo: true,
		Level:      DebugLevel,
		Handler:    handler,
	}
	l, err := NewWithConfig(cfg)
	require.NoError(t, err)
	return l
}

type testHandler struct {
	handler slog.Handler
	t       TestingT
}

func newTestHandler(t TestingT) *testHandler {
	writer := &testingWriter{
		t: t,
	}
	return &testHandler{
		t:       t,
		handler: slog.NewTextHandler(writer),
	}
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (t testHandler) Enabled(level slog.Level) bool {
	return t.handler.Enabled(level)
}

// Handle handles the Record.
func (t testHandler) Handle(r slog.Record) error {
	err := t.handler.Handle(r)
	if r.Level >= ErrorLevel {
		t.t.FailNow()
	}
	return err
}

// WithAttrs returns a new Handler whose attributes consist of
// both the receiver's attributes and the arguments.
func (t testHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return t.handler.WithAttrs(attrs)
}

// WithGroup returns a new Handler with the given group appended to
// the receiver's existing groups.
func (t testHandler) WithGroup(name string) slog.Handler {
	return t.handler.WithGroup(name)
}

type testingWriter struct {
	t TestingT
}

func (w testingWriter) Write(p []byte) (int, error) {
	n := len(p)
	p = bytes.TrimRight(p, "\n") // strip trailing newline because t.Log always adds one

	w.t.Logf("%s", p)
	return n, nil
}
