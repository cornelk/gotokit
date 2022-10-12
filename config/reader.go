// Package config contains configuration management helper.
package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
)

// Read reads the environment variables for the given prefix and unmarshals it into the config object.
func Read(prefix string, config interface{}) error {
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
