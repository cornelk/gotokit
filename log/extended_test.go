package log

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cornelk/gotokit/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtendedLogger(t *testing.T) {
	cfg, err := ConfigForEnv(env.Development)
	require.NoError(t, err)

	var buf bytes.Buffer
	cfg.Output = &buf
	cfg.Level = DebugLevel

	logger, err := NewWithConfig(cfg)
	require.NoError(t, err)

	extendedLogger := NewExtendedLogger(logger)

	extendedLogger.Debugf("test1")
	extendedLogger.Warnf("test1")
	extendedLogger.Errorf("test1")

	s := buf.String()
	all := strings.Split(s, "\n")
	assert.Len(t, all, 4)
}
