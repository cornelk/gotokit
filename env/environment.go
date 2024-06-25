// Package env provides a type for runtime environments.
package env

import (
	"errors"
	"fmt"
	"strings"
)

// Environment defines a runtime environment.
type Environment string

// Available commonly used environments.
const (
	Local       Environment = "local"
	Test        Environment = "test"
	Development Environment = "dev"
	Qa          Environment = "qa"
	Staging     Environment = "staging"
	Production  Environment = "prod"
)

// ErrNotAllowed describes an error for action
// which is not allowed in a given environment.
var ErrNotAllowed = errors.New("action not allowed for the environment")

// Parse returns an environment constant from a given string representation of it.
func Parse(envName string) (Environment, error) {
	switch Environment(strings.ToLower(envName)) {
	case Local:
		return Local, nil

	case Test:
		return Test, nil

	case Development:
		return Development, nil

	case Qa:
		return Qa, nil

	case Staging:
		return Staging, nil

	case Production, "production":
		return Production, nil

	default:
		return "", fmt.Errorf("unknown environment '%s'", envName)
	}
}

// Validate checks that the environment value is valid.
func (env Environment) Validate() error {
	_, err := Parse(string(env))
	return err
}
