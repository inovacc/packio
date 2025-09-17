// Package packio provides generic serialization wrappers for your types.
//
// The package exposes a single generic interface, Serializer[T], which defines
// Serialize and Deserialize methods to convert the wrapped value to and from
// bytes, plus Get/Set accessors and Clone for deep-copy semantics.
//
// You typically construct a wrapper with New, optionally selecting the
// format (JSON by default):
//
//	w := packio.New(MyType{})               // JSON by default
//	y := packio.New(MyType{}, packio.YAML) // YAML
//	t := packio.New(MyType{}, packio.TOML) // TOML
//
// Note: This package does not provide concurrency control; protect shared
// access with synchronization if you mutate wrappers from multiple goroutines.
package packio
