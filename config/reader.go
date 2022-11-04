// Package config contains configuration management helper.
package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
)

// Read reads the environment variables for the given prefix and unmarshals it into the config object.
// To support both prefixed and non prefixed envs at the same time it is recommended to call the function
// first without prefix and a second time with the prefix set. Only environment variables that exist will
// set a field in the config. This way, an environment variable set without a prefix can be overwritten
// by an environment variable with a prefix.
func Read(prefix string, config any) error {
	if prefix != "" {
		if !strings.HasSuffix(prefix, "_") {
			prefix += "_"
		}
		prefix = strings.ToUpper(prefix)
	}

	opts := env.Options{Prefix: prefix}
	if err := env.Parse(config, opts); err != nil {
		return fmt.Errorf("reading config from env: %w", err)
	}
	return nil
}
