package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	type Database struct {
		Host string `env:"HOST"`
	}

	type Config struct {
		Database Database `envPrefix:"DATABASE_"`
	}

	t.Setenv("DATABASE_HOST", "defaulthost")
	t.Setenv("TESTAPP_DATABASE_HOST", "localhost")

	var cfg Config
	require.NoError(t, Read(&cfg, "", "testapp"))
	assert.Equal(t, "localhost", cfg.Database.Host)

	require.NoError(t, Read(&cfg, "testapp", ""))
	assert.Equal(t, "defaulthost", cfg.Database.Host)

	require.Error(t, Read(cfg, "")) // not passing a pointer fails
}
