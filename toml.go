package packio

import (
	"github.com/pelletier/go-toml/v2"
)

// WithTOML adds TOML marshaling capabilities to any struct
// It mirrors the API of WithJSON but uses TOML under the hood.
type WithTOML[T any] struct {
	Data T
}

// Serialize implements Serializer using TOML format.
func (w *WithTOML[T]) Serialize() ([]byte, error) { return toml.Marshal(w.Data) }

// Deserialize implements Serializer using TOML format.
func (w *WithTOML[T]) Deserialize(b []byte) error { return toml.Unmarshal(b, &w.Data) }

// Get returns the underlying data
func (w *WithTOML[T]) Get() T { return w.Data }

// Set updates the underlying data
func (w *WithTOML[T]) Set(data T) { w.Data = data }

// Clone returns a new wrapper with a deep copy of the data
func (w *WithTOML[T]) Clone(empty bool) Serializer[T] {
	if empty {
		var zero T
		return New(zero, TOML)
	}

	b, err := toml.Marshal(w.Data)
	if err != nil {
		var zero T
		return New(zero, TOML)
	}

	var newData T
	if err := toml.Unmarshal(b, &newData); err != nil {
		var zero T
		return New(zero, TOML)
	}

	return New(newData, TOML)
}
