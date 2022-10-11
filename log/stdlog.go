package log

import "log"

// StdLogger creates a new standard library logger that outputs using the given logger.
func StdLogger(logger *Logger) *log.Logger {
	w := NewWriter(logger, InfoLevel)
	return log.New(w, "", 0)
}
