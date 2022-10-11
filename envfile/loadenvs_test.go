package envfile

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Setenv("test", "")

	tmpdir := t.TempDir()
	file := filepath.Join(tmpdir, ".env")
	require.NoError(t, os.WriteFile(file, []byte("TEST=1\n"), 0644))

	require.NoError(t, os.Chdir(tmpdir))
	Load()
	assert.Equal(t, "1", os.Getenv("TEST"))
}
