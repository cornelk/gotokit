package pointer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	i := 1
	p := Pointer(i)

	require.IsType(t, &i, p)
}
