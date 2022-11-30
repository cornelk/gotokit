package log

import (
	"fmt"
	"io"

	"github.com/cornelk/gotokit/env"
	"golang.org/x/exp/slog"
)

// Config represents configuration for a logger.
type Config struct {
	JSONOutput bool

	// CallerInfo adds a ("source", "file:line") attribute to the output
	// indicating the source code position of the log statement.
	CallerInfo bool

	Level Level

	Output io.Writer

	// Handler handles log records produced by a Logger..
	Handler slog.Handler
}

// ConfigForEnv returns the default config for the given environment.
// The returned config can be adjusted and used to create a logger with
// custom config using the NewWithConfig() function.
func ConfigForEnv(environment env.Environment) (Config, error) {
	cfg := Config{
		Level: DefaultLevel(),
	}

	switch environment {
	case env.Test, env.Development:
		cfg.JSONOutput = false
		cfg.CallerInfo = true

	case env.Staging, env.Production:
		cfg.JSONOutput = true
		cfg.CallerInfo = false

	default:
		return Config{}, fmt.Errorf("invalid environment specified '%v'", environment)
	}

	return cfg, nil
}
