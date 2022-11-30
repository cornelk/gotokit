package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestLogger(t *testing.T) {
	logger := NewTestLogger(t)
	assert.Equal(t, DebugLevel, logger.level.Level())
}
