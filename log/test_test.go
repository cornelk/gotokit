package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestLogger(t *testing.T) {
	logger := NewTestLogger(t)
	assert.True(t, logger.Core().Enabled(DebugLevel))
}
