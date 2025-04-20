package wrapper

import (
	"encoding/json"
)

// Wrapper is an interface that provides JSON marshaling capabilities
type Wrapper interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

// WithJSON adds JSON marshaling capabilities to any struct
type WithJSON[T any] struct {
	Data T
}

// NewWrapper creates a new instance of withJSON
func NewWrapper[T any](data T) *WithJSON[T] {
	return &WithJSON[T]{
		Data: data,
	}
}

// MarshalJSON implements the json.Marshaler interface
func (w *WithJSON[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.Data)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (w *WithJSON[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &w.Data)
}

// Get returns the underlying data
func (w *WithJSON[T]) Get() T {
	return w.Data
}

// Set updates the underlying data
func (w *WithJSON[T]) Set(data T) {
	w.Data = data
}

// Clone returns a new empty wrapper of the same type
func (w *WithJSON[T]) Clone() *WithJSON[T] {
	var zero T
	return NewWrapper(zero)
}
