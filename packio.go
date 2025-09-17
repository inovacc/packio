package packio

type SerdeType int

const (
	JSON SerdeType = iota
	YAML
	TOML
)

// Serializer is a unified interface for (de)serializing wrapped data.
// It abstracts over the underlying format (JSON, YAML, TOML, ...).
// Implementations should provide deep-copy semantics via Clone.
type Serializer[T any] interface {
	// Serialize converts the underlying data into bytes using the wrapper's format.
	Serialize() ([]byte, error)
	// Deserialize populates the underlying data from the provided bytes using the wrapper's format.
	Deserialize(v []byte) error
	// Clone returns a new wrapper of the same format.
	// If empty is true, the new wrapper contains the zero value of T; otherwise it contains a deep copy of the data.
	Clone(empty bool) Serializer[T]
	// Get returns the underlying data.
	Get() T
	// Set replaces the underlying data.
	Set(data T)
}

// New creates a new wrapper instance.
// If a format is provided, it will be used; otherwise it defaults to JSON.
// The return type is the unified Serializer[T] interface for flexible usage.
func New[T any](data T, format ...SerdeType) Serializer[T] {
	f := JSON
	if len(format) > 0 {
		f = format[0]
	}

	switch f {
	case YAML:
		return &WithYAML[T]{Data: data}
	case TOML:
		return &WithTOML[T]{Data: data}
	case JSON:
		fallthrough
	default:
		return &WithJSON[T]{Data: data}
	}
}
