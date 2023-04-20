// Package env provides a type for runtime environments.
package env

import (
	"fmt"
	"strings"
)

// Environment defines a runtime environment.
type Environment string

// Available environments.
const (
	Test        Environment = "test"
	Development Environment = "dev"
	Qa          Environment = "qa"
	Staging     Environment = "staging"
	Production  Environment = "prod"
)

// Parse returns an environment constant from a given string representation of it.
func Parse(envName string) (Environment, error) {
	switch Environment(strings.ToLower(envName)) {
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
