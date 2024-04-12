package config

import (
	"reflect"

	"github.com/caarlos0/env/v10"
)

// ParserFunc defines the signature of a function that can be used within `CustomParsers`.
type ParserFunc = env.ParserFunc

// Options for the config reader.
type Options struct {
	Prefixes []string                    // Prefixes define a prefix for each key.
	FuncMap  map[reflect.Type]ParserFunc // Custom parse functions for different types.
}
