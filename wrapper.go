package wrapper

// Serializer is a unified interface for (de)serializing wrapped data.
// It abstracts over the underlying format (JSON, YAML, TOML, ...).
// Implementations should provide deep-copy semantics via Clone.
type Serializer[T any] interface {
	// Serialize converts the underlying data into bytes using the wrapper's format.
	Serialize() ([]byte, error)
	// Deserialize populates the underlying data from the provided bytes using the wrapper's format.
	Deserialize([]byte) error
	// Clone returns a new wrapper of the same format.
	// If empty is true, the new wrapper contains the zero value of T; otherwise it contains a deep copy of the data.
	Clone(empty bool) Serializer[T]
	// Get returns the underlying data.
	Get() T
	// Set replaces the underlying data.
	Set(data T)
}

// NewWrapper creates a new JSON wrapper instance.
// Returning the concrete type keeps backward compatibility; callers can still accept the Serializer[T] interface.
func NewWrapper[T any](data T) *WithJSON[T] {
	return &WithJSON[T]{Data: data}
}
