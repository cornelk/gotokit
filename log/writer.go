package log

import (
	"context"
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
		o.Logger.Log(context.TODO(), DebugLevel, string(p))

	case InfoLevel:
		o.Logger.Log(context.TODO(), InfoLevel, string(p))

	case ErrorLevel:
		o.Logger.Log(context.TODO(), ErrorLevel, string(p))

	case FatalLevel:
		o.Logger.Log(context.TODO(), ErrorLevel, string(p))
		os.Exit(1)
	}

	return len(p), nil
}
