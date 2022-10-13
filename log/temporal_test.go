package log

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeyValuesToFields(t *testing.T) {
	fields := keyValuesToFields("key1", "value1")
	require.Len(t, fields, 1)

	fields = keyValuesToFields("key1", "value1", "key2")
	require.Len(t, fields, 2)
}
