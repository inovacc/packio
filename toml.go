package wrapper

import (
	"github.com/pelletier/go-toml/v2"
)

// WithTOML adds TOML marshaling capabilities to any struct
// It mirrors the API of WithJSON but uses TOML under the hood.
type WithTOML[T any] struct {
	Data T
}

// NewTOMLWrapper creates a new instance of WithTOML
func NewTOMLWrapper[T any](data T) *WithTOML[T] { return &WithTOML[T]{Data: data} }

// Serialize implements Serializer using TOML format.
func (w *WithTOML[T]) Serialize() ([]byte, error) { return toml.Marshal(w.Data) }

// Deserialize implements Serializer using TOML format.
func (w *WithTOML[T]) Deserialize(b []byte) error { return toml.Unmarshal(b, &w.Data) }

// MarshalTOMLBytes Backward-compat helpers for code that may still call them.
func (w *WithTOML[T]) MarshalTOMLBytes() ([]byte, error) { return toml.Marshal(w.Data) }
func (w *WithTOML[T]) UnmarshalTOMLBytes(b []byte) error { return toml.Unmarshal(b, &w.Data) }

// Get returns the underlying data
func (w *WithTOML[T]) Get() T { return w.Data }

// Set updates the underlying data
func (w *WithTOML[T]) Set(data T) { w.Data = data }

// Clone returns a new wrapper with a deep copy of the data
func (w *WithTOML[T]) Clone(empty bool) Serializer[T] {
	if empty {
		var zero T
		return NewTOMLWrapper(zero)
	}

	b, err := toml.Marshal(w.Data)
	if err != nil {
		var zero T
		return NewTOMLWrapper(zero)
	}

	var newData T
	if err := toml.Unmarshal(b, &newData); err != nil {
		var zero T
		return NewTOMLWrapper(zero)
	}

	return NewTOMLWrapper(newData)
}
