// Package jsonutils offers additional JSON processing helpers.
package jsonutils

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// ValidateRemarshal verifies that the json source and the given object match when being marshalled.
// This allows to detect changes in struct fields that are returned from APIs and not .
func ValidateRemarshal(source []byte, object any) error {
	var input map[string]interface{}
	if err := json.Unmarshal(source, &input); err != nil {
		return fmt.Errorf("unmarshalling source: %w", err)
	}

	b, err := json.Marshal(object)
	if err != nil {
		return fmt.Errorf("marshalling object: %w", err)
	}
	var output map[string]interface{}
	if err := json.Unmarshal(b, &output); err != nil {
		return fmt.Errorf("unmarshalling object: %w", err)
	}

	if diff := cmp.Diff(input, output); diff != "" {
		return fmt.Errorf("remarshal mismatch (-want +got):\n%s", diff)
	}
	return nil
}
