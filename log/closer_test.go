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

type testCloser struct {
	err error
}

func (t testCloser) Close() error {
	return t.err
}

func TestLoggerCloser(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)
	var buf bytes.Buffer
	cfg.Output = &buf
	cfg.TimeFormat = "-"

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)

	closer := testCloser{}
	msg := "closing failed"

	logger.Closer(closer, msg)
	output := buf.String()
	assert.NotContains(t, output, "ERROR")
	assert.NotContains(t, output, msg)

	errMsg := "failure"
	closer.err = errors.New(errMsg)
	logger.Closer(closer, msg)
	output = buf.String()
	assert.Contains(t, output, "ERROR")
	assert.Contains(t, output, msg)
	assert.Contains(t, output, errMsg)
}

type testCloserCtx struct {
	err error
}

func (t testCloserCtx) Close(context.Context) error {
	return t.err
}

func TestLoggerCloserCtx(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)
	var buf bytes.Buffer
	cfg.Output = &buf
	cfg.TimeFormat = "-"

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)

	ctx := context.Background()
	closer := testCloserCtx{}
	msg := "closing failed"

	logger.CloserCtx(ctx, closer, msg)
	output := buf.String()
	assert.NotContains(t, output, "ERROR")
	assert.NotContains(t, output, msg)

	errMsg := "failure"
	closer.err = errors.New(errMsg)
	logger.CloserCtx(ctx, closer, msg)
	output = buf.String()
	assert.Contains(t, output, "ERROR")
	assert.Contains(t, output, msg)
	assert.Contains(t, output, errMsg)
}
