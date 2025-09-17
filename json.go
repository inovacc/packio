package packio

import "encoding/json"

// WithJSON adds JSON marshaling capabilities to any struct
type WithJSON[T any] struct {
	Data T
}

// Serialize implements Serializer using JSON format.
func (w *WithJSON[T]) Serialize() ([]byte, error) { return json.Marshal(w.Data) }

// Deserialize implements Serializer using JSON format.
func (w *WithJSON[T]) Deserialize(data []byte) error { return json.Unmarshal(data, &w.Data) }

// Get returns the underlying data
func (w *WithJSON[T]) Get() T { return w.Data }

// Set updates the underlying data
func (w *WithJSON[T]) Set(data T) { w.Data = data }

// Clone returns a new wrapper with a deep copy of the data
func (w *WithJSON[T]) Clone(empty bool) Serializer[T] {
	if empty {
		var zero T
		return New(zero)
	}

	// Marshal the original data
	data, err := json.Marshal(w.Data)
	if err != nil {
		var zero T
		return New(zero)
	}

	// Create a new instance
	var newData T

	// Unmarshal into the new instance
	if err := json.Unmarshal(data, &newData); err != nil {
		var zero T
		return New(zero)
	}

	return New(newData)
}
