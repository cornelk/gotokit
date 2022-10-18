// Package jsonutils offers additional JSON processing helpers.
package jsonutils

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// CompareRawJSON verifies that both given json raw messages match whilst ignoring
// the key order and any whitespace differences. The parameters can be of type
// json.RawMessage as this is an alias for []byte.
// Any difference will be returned as error.
func CompareRawJSON(expected, actual []byte) error {
	var expectedMap map[string]any
	if err := json.Unmarshal(expected, &expectedMap); err != nil {
		return fmt.Errorf("unmarshalling expected parameter: %w", err)
	}

	var actualMap map[string]any
	if err := json.Unmarshal(actual, &actualMap); err != nil {
		return fmt.Errorf("unmarshalling actual parameter: %w", err)
	}

	if diff := cmp.Diff(expectedMap, actualMap); diff != "" {
		return fmt.Errorf("mismatch (-want +got):\n%s", diff)
	}
	return nil
}

// ValidateRemarshal verifies that the json source and the given object match when being marshalled.
// This allows to detect changes in struct fields that are returned from APIs and not handled in the struct type.
// Any difference will be returned as error.
func ValidateRemarshal(source []byte, object any, options ...cmp.Option) error {
	var sourceMap map[string]any
	if err := json.Unmarshal(source, &sourceMap); err != nil {
		return fmt.Errorf("unmarshalling source: %w", err)
	}

	b, err := json.Marshal(object)
	if err != nil {
		return fmt.Errorf("marshalling object: %w", err)
	}
	var objectMap map[string]any
	if err := json.Unmarshal(b, &objectMap); err != nil {
		return fmt.Errorf("unmarshalling object: %w", err)
	}

	if diff := cmp.Diff(sourceMap, objectMap, options...); diff != "" {
		return fmt.Errorf("mismatch (-want +got):\n%s", diff)
	}
	return nil
}
