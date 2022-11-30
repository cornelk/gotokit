package log

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/cornelk/gotokit/env"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseLogger(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)

	var buf bytes.Buffer
	cfg.Output = &buf
	cfg.Level = DebugLevel

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)

	ctx := context.Background()
	dbLogger := NewDatabaseLogger(logger)
	assert.Equal(t, tracelog.LogLevelDebug, dbLogger.Level())

	dbLogger.Log(ctx, tracelog.LogLevelTrace, "test1", nil)
	dbLogger.Log(ctx, tracelog.LogLevelInfo, "test2", nil)
	dbLogger.Log(ctx, tracelog.LogLevelWarn, "test3", nil)
	dbLogger.Log(ctx, tracelog.LogLevelError, "test4", nil)

	s := buf.String()
	all := strings.Split(s, "\n")
	assert.Len(t, all, 5)
}
