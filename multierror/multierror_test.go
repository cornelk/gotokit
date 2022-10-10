package multierror

import (
	"errors"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

func TestDefaultLibMultiError(t *testing.T) {
	var result error
	result = multierror.Append(result, errors.New("error1"))
	result = multierror.Append(result, errors.New("error2"))

	expected := `2 errors occurred:
	* error1
	* error2

`
	assert.Equal(t, expected, result.Error())
}

func TestDefaultLibSingleError(t *testing.T) {
	var result error
	result = multierror.Append(result, errors.New("error1"))

	expected := `1 error occurred:
	* error1

`
	assert.Equal(t, expected, result.Error())
}

func TestMultiError(t *testing.T) {
	result := New()
	result = Append(result, errors.New("error1"))
	result = Append(result, errors.New("error2"))

	expected := `	* error1
	* error2
`
	assert.Equal(t, expected, result.Error())
}

func TestSingleError(t *testing.T) {
	result := New()
	result = Append(result, errors.New("error1"))

	expected := `error1`
	assert.Equal(t, expected, result.Error())
}
