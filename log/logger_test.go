package log

import (
	"bytes"
	"context"
	"errors"
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
	cfg.JSONOutput = false

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)
	named := logger.Named("test")
	assert.Equal(t, DebugLevel, named.level.Level())
}

func TestLoggerFatal(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)
	var buf bytes.Buffer

	cfg.CallerInfo = false
	cfg.JSONOutput = false
	cfg.Output = &buf
	cfg.TimeFormat = "-"

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)
	exited := false
	fatalExitFunc = func() {
		exited = true
	}

	logger.Fatal("something bad happened", Err(errors.New("network error")))

	assert.True(t, exited)
	output := buf.String()
	assert.Equal(t, "FATAL   something bad happened {\"error\":\"network error\"}\n", output)
}

func TestLoggerTrace(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)
	var buf bytes.Buffer

	cfg.CallerInfo = false
	cfg.Level = TraceLevel
	cfg.Output = &buf
	cfg.TimeFormat = "-"

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)
	exited := false
	fatalExitFunc = func() {
		exited = true
	}

	logger.Trace("something happened")

	assert.False(t, exited)
	output := buf.String()
	assert.Equal(t, "TRACE   something happened\n", output)
}

func TestLoggerCaller(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)
	var buf bytes.Buffer

	cfg.CallerInfo = true
	cfg.Level = TraceLevel
	cfg.Output = &buf
	cfg.TimeFormat = "-"

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)

	logger.Trace("something happened")

	output := buf.String()
	assert.Contains(t, output, "TRACE")
	assert.Contains(t, output, "logger_test.go")
	assert.Contains(t, output, "something happened\n")
}

func TestLoggerSlog(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)

	slogger := logger.Slog()
	assert.NotNil(t, slogger)
}
