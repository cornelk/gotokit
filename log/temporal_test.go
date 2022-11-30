package log

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cornelk/gotokit/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemporalLogger(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)

	var buf bytes.Buffer
	cfg.Output = &buf
	cfg.Level = DebugLevel

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)

	temporalLogger := NewTemporalLogger(logger)

	temporalLogger.Debug("test1")
	temporalLogger.Info("test2")
	temporalLogger.Warn("test1")
	temporalLogger.Error("test1")

	s := buf.String()
	all := strings.Split(s, "\n")
	assert.Len(t, all, 5)
}

func TestKeyValuesToFields(t *testing.T) {
	fields := keyValuesToFields("key1", "value1")
	require.Len(t, fields, 1)

	fields = keyValuesToFields("key1", "value1", "key2")
	require.Len(t, fields, 2)
}
