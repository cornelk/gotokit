package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	env, err := Parse("prod")
	require.NoError(t, err)
	assert.Equal(t, Production, env)

	_, err = Parse("invalid")
	require.Error(t, err)
}
