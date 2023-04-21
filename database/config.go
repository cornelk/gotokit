package database

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/jackc/pgx/v5/tracelog"
)

// LoggerContract defines the logger interface used by the database implementation.
type LoggerContract interface {
	Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any)
	Level() tracelog.LogLevel
}

// Config contains the database configuration.
type Config struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`

	Logger LoggerContract
}

// Validate checks that all mandatory configuration values are set.
func (cfg *Config) Validate() error {
	var errs []error

	port, err := strconv.Atoi(cfg.Port)
	if err == nil {
		if port == 0 || port > math.MaxUint16 {
			errs = append(errs, errors.New("invalid port set"))
		}
	} else {
		errs = append(errs, fmt.Errorf("parsing port: %w", err))
	}

	if cfg.Database == "" {
		errs = append(errs, errors.New("database is not set"))
	}
	if cfg.User == "" {
		errs = append(errs, errors.New("user is not set"))
	}

	return errors.Join(errs...)
}
