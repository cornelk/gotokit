package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReaderPrefixes(t *testing.T) {
	type Database struct {
		Host string `env:"HOST"`
	}

	type myConfig struct {
		Database Database `envPrefix:"DATABASE_"`
	}

	t.Setenv("DATABASE_HOST", "defaulthost")
	t.Setenv("TESTAPP_DATABASE_HOST", "localhost")

	var cfg myConfig
	opts := Options{
		Prefixes: []string{"", "testapp"},
	}
	require.NoError(t, Read(&cfg, opts))
	assert.Equal(t, "localhost", cfg.Database.Host)

	opts.Prefixes = []string{"testapp", ""}
	require.NoError(t, Read(&cfg, opts))
	assert.Equal(t, "defaulthost", cfg.Database.Host)

	require.Error(t, Read(cfg, opts)) // not passing a pointer fails
}

func TestReaderFuncMap(t *testing.T) {
	type foo struct {
		name string
	}

	type Database struct {
		Host string `env:"HOST"`
	}

	type myConfig struct {
		Database Database `envPrefix:"DATABASE_"`
		Foo      foo      `env:"FOO" envDefault:"admin"`
	}

	var called int
	fooFunc := func(v string) (interface{}, error) {
		assert.Equal(t, "admin", v)
		called++
		return foo{name: v}, nil
	}

	var cfg myConfig
	opts := Options{
		FuncMap: map[reflect.Type]ParserFunc{
			reflect.TypeOf(foo{}): fooFunc,
		},
	}

	require.NoError(t, Read(&cfg, opts))
	assert.Equal(t, "admin", cfg.Foo.name)
	assert.Equal(t, 1, called)
}
