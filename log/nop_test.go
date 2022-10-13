package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestNewNop(t *testing.T) {
	logger := NewNop()

	core, observed := observer.New(logger.Level())
	logger.Logger = zap.New(core)

	logger.Error("test")

	all := observed.TakeAll()
	assert.Len(t, all, 0)
}
