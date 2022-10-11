package log

import (
	"testing"

	"github.com/cornelk/gotokit/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	prev := DefaultLevel()
	SetDefaultLevel(DebugLevel)
	defer SetDefaultLevel(prev)

	logger, err := New()
	require.NoError(t, err)

	assert.True(t, logger.Core().Enabled(DebugLevel))
}

func TestNewWithConfig(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)
	named := logger.Named("test")
	assert.Equal(t, DebugLevel, named.level.Level())
}
