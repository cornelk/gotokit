package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestTemporalLogger(t *testing.T) {
	logger, err := New()
	require.NoError(t, err)
	logger.SetLevel(DebugLevel)

	core, observed := observer.New(logger.Level())
	logger.Logger = zap.New(core)

	temporalLogger := NewTemporalLogger(logger)

	temporalLogger.Debug("test1")
	temporalLogger.Info("test2")
	temporalLogger.Warn("test1")
	temporalLogger.Error("test1")

	all := observed.TakeAll()
	assert.Len(t, all, 4)
}

func TestKeyValuesToFields(t *testing.T) {
	fields := keyValuesToFields("key1", "value1")
	require.Len(t, fields, 1)

	fields = keyValuesToFields("key1", "value1", "key2")
	require.Len(t, fields, 2)
}
