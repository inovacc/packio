package wrapper

import (
	"gopkg.in/yaml.v3"
)

// WithYAML adds YAML marshaling capabilities to any struct
// It mirrors the API of WithJSON but uses YAML under the hood.
type WithYAML[T any] struct {
	Data T
}

// NewYAMLWrapper creates a new instance of WithYAML
func NewYAMLWrapper[T any](data T) *WithYAML[T] { return &WithYAML[T]{Data: data} }

// Serialize implements Serializer using YAML format.
func (w *WithYAML[T]) Serialize() ([]byte, error) { return yaml.Marshal(w.Data) }

// Deserialize implements Serializer using YAML format.
func (w *WithYAML[T]) Deserialize(b []byte) error { return yaml.Unmarshal(b, &w.Data) }

// Get returns the underlying data
func (w *WithYAML[T]) Get() T { return w.Data }

// Set updates the underlying data
func (w *WithYAML[T]) Set(data T) { w.Data = data }

// Clone returns a new wrapper with a deep copy of the data
func (w *WithYAML[T]) Clone(empty bool) Serializer[T] {
	if empty {
		var zero T
		return NewYAMLWrapper(zero)
	}

	b, err := yaml.Marshal(w.Data)
	if err != nil {
		var zero T
		return NewYAMLWrapper(zero)
	}

	var newData T
	if err := yaml.Unmarshal(b, &newData); err != nil {
		var zero T
		return NewYAMLWrapper(zero)
	}

	return NewYAMLWrapper(newData)
}
