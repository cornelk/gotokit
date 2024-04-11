package config

import (
	"reflect"

	"github.com/caarlos0/env/v10"
)

// Options for the config reader.
type Options struct {
	Prefixes []string                        // Prefixes define a prefix for each key.
	FuncMap  map[reflect.Type]env.ParserFunc // Custom parse functions for different types.
}
