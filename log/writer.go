package log

import (
	"io"
	"os"
)

type writer struct {
	*Logger
	level Level
}

// NewWriter creates a new io.Writer that writes all messages to the given
// logger using the given log level.
func NewWriter(logger *Logger, level Level) io.Writer {
	return &writer{
		Logger: logger,
		level:  level,
	}
}

// Write implements the io.Writer interface.
func (o writer) Write(p []byte) (int, error) {
	switch o.level {
	case DebugLevel:
		o.Logger.LogDepth(1, DebugLevel, string(p))

	case InfoLevel:
		o.Logger.LogDepth(1, InfoLevel, string(p))

	case ErrorLevel:
		o.Logger.LogDepth(1, ErrorLevel, string(p))

	case FatalLevel:
		o.Logger.LogDepth(1, ErrorLevel, string(p))
		os.Exit(1)
	}

	return len(p), nil
}
