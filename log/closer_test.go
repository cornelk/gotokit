package log

import (
	"bytes"
	"errors"
	"strings"
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

func TestLoggerCloserLogOnError(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)
	var buf bytes.Buffer
	cfg.Output = &buf
	cfg.TimeFormat = "-"

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)

	closer := testCloser{}
	msg := "closing failed"

	logger.CloserLogOnError(closer, msg)
	output := buf.String()
	assert.False(t, strings.Contains(output, "ERROR"))
	assert.False(t, strings.Contains(output, msg))

	errMsg := "failure"
	closer.err = errors.New(errMsg)
	logger.CloserLogOnError(closer, msg)
	output = buf.String()
	assert.True(t, strings.Contains(output, "ERROR"))
	assert.True(t, strings.Contains(output, msg))
	assert.True(t, strings.Contains(output, errMsg))
}
