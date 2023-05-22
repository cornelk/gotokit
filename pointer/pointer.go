// Package pointer helps to create a pointer to a value.
package pointer

// Pointer returns a pointer to the passed value.
func Pointer[T any](value T) *T {
	return &value
}
