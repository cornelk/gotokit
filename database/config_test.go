package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigValidate(t *testing.T) {
	cfg := Config{
		Host: "invalid:host",
		Port: "a",
	}

	err := cfg.Validate()
	require.Error(t, err)
}
