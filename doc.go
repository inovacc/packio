// Package wrapper provides generic serialization wrappers for your types.
//
// The package exposes a single generic interface, Serializer[T], which defines
// Serialize and Deserialize methods to convert the wrapped value to and from
// bytes, plus Get/Set accessors and Clone for deep-copy semantics.
//
// You typically construct a wrapper with NewWrapper, optionally selecting the
// format (JSON by default):
//
//	w := wrapper.NewWrapper(MyType{})               // JSON by default
//	y := wrapper.NewWrapper(MyType{}, wrapper.YAML) // YAML
//	t := wrapper.NewWrapper(MyType{}, wrapper.TOML) // TOML
//
// For direct format construction you can also use NewYAMLWrapper and
// NewTOMLWrapper helpers.
//
// Note: This package does not provide concurrency control; protect shared
// access with synchronization if you mutate wrappers from multiple goroutines.
package wrapper
