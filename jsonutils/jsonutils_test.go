package jsonutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompareRawJSON(t *testing.T) {
	input1 := []byte(`{ "a": 1, "b": 2 }`)
	input2 := []byte(`{ "b": 2,   "a": 1 }`)
	assert.NoError(t, CompareRawJSON(input1, input2))

	input3 := []byte(`{ "a": 1, "b": 2, "c": 3 }`)
	err := CompareRawJSON(input1, input3)
	require.Error(t, err)
	assert.ErrorContains(t, err, "mismatch")
}

func TestValidateRemarshal(t *testing.T) {
	type test struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	input1 := []byte(`{ "a": 1, "b": 2 }`)
	object := test{
		A: 1,
		B: 2,
	}
	assert.NoError(t, ValidateRemarshal(input1, object))

	input2 := []byte(`{ "a": 1, "b": 2, "c": 3 }`)
	err := ValidateRemarshal(input2, object)
	require.Error(t, err)
	assert.ErrorContains(t, err, "mismatch")
}
