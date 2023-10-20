package app

import (
	"context"
	"time"
)

// Sleep pauses the current goroutine for at least the duration d.
// The function will return earlier in case that the passed context
// is cancelled before.
func Sleep(ctx context.Context, duration time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-time.After(duration):
		return nil
	}
}
