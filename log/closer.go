package log

import (
	"errors"
	"io"
	"os"
)

// CloserLogOnError calls the closer function and if an error gets returned it logs an error.
// This function is useful when using patterns like defer resp.Body.Close() which now become:
// defer logger.CloserLogOnError(resp.Body, "closing body").
func (l *Logger) CloserLogOnError(closer io.Closer, msg string) {
	err := closer.Close()
	if err == nil || errors.Is(err, os.ErrClosed) {
		return
	}

	l.Error(msg, Err(err))
}
