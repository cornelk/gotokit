// Package config contains configuration management helper.
package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v10"
)

// Read reads the environment variables for the given prefix and unmarshals it into the config object.
// To support both prefixed and non prefixed envs at the same time it is recommended to call the function
// with an empty first prefix and a second set prefix. Only environment variables that exist will
// set a field in the config. This way, an environment variable set without a prefix can be overwritten
// by an environment variable with a prefix.
func Read(config any, opts Options) error {
	// fall back to a default prefix if none are provided
	if len(opts.Prefixes) == 0 {
		opts.Prefixes = []string{""}
	}

	for _, prefix := range opts.Prefixes {
		if prefix != "" {
			if !strings.HasSuffix(prefix, "_") {
				prefix += "_"
			}
			prefix = strings.ToUpper(prefix)
		}

		envOpts := env.Options{
			Prefix:  prefix,
			FuncMap: opts.FuncMap,
		}
		if err := env.ParseWithOptions(config, envOpts); err != nil {
			return fmt.Errorf("reading config from env: %w", err)
		}
	}
	return nil
}
