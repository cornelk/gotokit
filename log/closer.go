package log

import (
	"context"
	"errors"
	"io"
	"os"
)

// Closer calls the closer function and if an error gets returned it logs an error.
// This function is useful when using patterns like defer resp.Body.Close() which now become:
// defer logger.Closer(resp.Body, "closing body").
func (l *Logger) Closer(closer io.Closer, msg string) {
	err := closer.Close()
	if err == nil || errors.Is(err, os.ErrClosed) {
		return
	}

	l.Error(msg, Err(err))
}

// closerCtx is the interface that wraps the extended Close method.
type closerCtx interface {
	Close(ctx context.Context) error
}

// CloserCtx calls the closer function and if an error gets returned it logs an error.
func (l *Logger) CloserCtx(ctx context.Context, closer closerCtx, msg string) {
	err := closer.Close(ctx)
	if err == nil || errors.Is(err, os.ErrClosed) {
		return
	}

	l.ErrorContext(ctx, msg, Err(err))
}
