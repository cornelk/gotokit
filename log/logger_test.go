package log

import (
	"context"
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

	assert.True(t, logger.Enabled(context.TODO(), DebugLevel))
}

func TestNewWithConfig(t *testing.T) {
	prev := DefaultLevel()
	SetDefaultLevel(DebugLevel)
	defer SetDefaultLevel(prev)

	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)
	named := logger.Named("test")
	assert.Equal(t, DebugLevel, named.level.Level())
}
