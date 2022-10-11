package log

import (
	"io"
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
		o.Logger.Debug(string(p))

	case InfoLevel:
		o.Logger.Info(string(p))

	case ErrorLevel:
		o.Logger.Error(string(p))

	case FatalLevel:
		o.Logger.Fatal(string(p))
	}

	return len(p), nil
}
