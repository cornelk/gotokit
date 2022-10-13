package log

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestDatabaseLogger(t *testing.T) {
	logger, err := New()
	require.NoError(t, err)
	logger.SetLevel(DebugLevel)

	core, observed := observer.New(logger.Level())
	logger.Logger = zap.New(core)

	ctx := context.Background()
	dbLogger := NewDatabaseLogger(logger)
	assert.Equal(t, tracelog.LogLevelDebug, dbLogger.Level())

	dbLogger.Log(ctx, tracelog.LogLevelTrace, "test1", nil)
	dbLogger.Log(ctx, tracelog.LogLevelInfo, "test2", nil)
	dbLogger.Log(ctx, tracelog.LogLevelWarn, "test3", nil)
	dbLogger.Log(ctx, tracelog.LogLevelError, "test4", nil)

	all := observed.TakeAll()
	assert.Len(t, all, 4)
}
