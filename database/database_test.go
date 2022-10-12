package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := Config{
		Host:     "invalid:host", // trigger DNS lookup error
		Port:     "1234",
		User:     "default",
		Database: "test",
	}

	ctx := context.Background()
	db, err := New(ctx, cfg)
	assert.Error(t, err)
	assert.Nil(t, db)
}

func TestNewStdlib(t *testing.T) {
	cfg := Config{
		Host:     "invalid:host", // trigger DNS lookup error
		Port:     "1234",
		User:     "default",
		Database: "test",
	}

	ctx := context.Background()
	db, err := NewStdlib(ctx, cfg)
	assert.Error(t, err)
	assert.Nil(t, db)
}
