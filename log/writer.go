package log

import (
	"io"
	"os"
)

var _ io.Writer = &Writer{}

// Writer implements an io.Writer compatible log writer.
type Writer struct {
	*Logger
	level Level
}

// NewWriter creates a new io.Writer that writes all messages to the given
// logger using the given log level.
func NewWriter(logger *Logger, level Level) *Writer {
	return &Writer{
		Logger: logger,
		level:  level,
	}
}

// Write implements the io.Writer interface.
func (o Writer) Write(p []byte) (int, error) {
	switch o.level {
	case DebugLevel:
		o.Logger.Log(nil, DebugLevel, string(p))

	case InfoLevel:
		o.Logger.Log(nil, InfoLevel, string(p))

	case ErrorLevel:
		o.Logger.Log(nil, ErrorLevel, string(p))

	case FatalLevel:
		o.Logger.Log(nil, ErrorLevel, string(p))
		os.Exit(1)
	}

	return len(p), nil
}
