package database

import (
	"errors"
	"testing"

	"github.com/cornelk/gotokit/multierror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigValidate(t *testing.T) {
	cfg := Config{
		Host: "invalid:host",
		Port: "a",
	}

	err := cfg.Validate()
	require.Error(t, err)

	var e *multierror.Error
	require.True(t, errors.As(err, &e))
	assert.Equal(t, 3, e.Len())
}
