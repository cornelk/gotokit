package log

import (
	"context"
	"errors"
	"log/slog"
)

var _ slog.Handler = &FanOutHandler{}

// FanOutHandler implements a fan-out to multiple log handlers.
type FanOutHandler struct {
	handlers []slog.Handler
}

// NewFanOutHandler creates a new fan-out log handler.
func NewFanOutHandler(handlers ...slog.Handler) *FanOutHandler {
	return &FanOutHandler{
		handlers: handlers,
	}
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (h *FanOutHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

// Handle handles the Record.
func (h *FanOutHandler) Handle(ctx context.Context, r slog.Record) error {
	var errs []error

	for _, handler := range h.handlers {
		if err := handler.Handle(ctx, r); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

// WithAttrs returns a new Handler whose attributes consist of
// both the receiver's attributes and the arguments.
// nolint: ireturn
func (h *FanOutHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handlers := make([]slog.Handler, 0, len(h.handlers))

	for _, handler := range h.handlers {
		handlers = append(handlers, handler.WithAttrs(attrs))
	}

	return NewFanOutHandler(handlers...)
}

// WithGroup returns a new Handler with the given group appended to
// the receiver's existing groups.
// nolint: ireturn
func (h *FanOutHandler) WithGroup(name string) slog.Handler {
	handlers := make([]slog.Handler, 0, len(h.handlers))

	for _, handler := range h.handlers {
		handlers = append(handlers, handler.WithGroup(name))
	}

	return NewFanOutHandler(handlers...)
}
