package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestExtendedLogger(t *testing.T) {
	logger, err := New()
	require.NoError(t, err)
	logger.SetLevel(DebugLevel)

	core, observed := observer.New(logger.Level())
	logger.Logger = zap.New(core)

	extendedLogger := NewExtendedLogger(logger)

	extendedLogger.Debugf("test1")
	extendedLogger.Warnf("test1")
	extendedLogger.Errorf("test1")

	all := observed.TakeAll()
	assert.Len(t, all, 3)
}
