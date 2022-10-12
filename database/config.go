package database

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/cornelk/gotokit/multierror"
)

// Config contains the database configuration.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string

	Logger Logger
}

// Validate checks that all mandatory configuration values are set.
func (cfg *Config) Validate() error {
	errs := multierror.New()

	port, err := strconv.Atoi(cfg.Port)
	if err == nil {
		if port == 0 || port > math.MaxUint16 {
			errs = multierror.Append(errs, errors.New("invalid port set"))
		}
	} else {
		errs = multierror.Append(errs, fmt.Errorf("parsing port: %w", err))
	}

	if cfg.Database == "" {
		errs = multierror.Append(errs, errors.New("database is not set"))
	}
	if cfg.User == "" {
		errs = multierror.Append(errs, errors.New("user is not set"))
	}

	return errs.ErrorOrNil()
}
