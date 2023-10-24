package app

import (
	"context"
	"fmt"
	"time"
)

// Sleep pauses the current goroutine for at least the duration d.
// The function will return earlier in case that the passed context
// is cancelled before.
func Sleep(ctx context.Context, duration time.Duration) error {
	select {
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("context done: %w", err)
		}

	case <-time.After(duration):
	}

	return nil
}
