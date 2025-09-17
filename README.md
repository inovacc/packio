# packio — Generic Serialization Wrappers (JSON, YAML, TOML)

A generic Go package that provides a simple way to add marshaling and unmarshaling capabilities to any struct type for multiple formats.
This wrapper is particularly useful when you need to add serialization functionality to existing structs without modifying them directly.

## Features

- ✅ Generic implementation using Go type parameters (Go 1.23+)
- ✅ Simple and intuitive API
- ✅ Built-in `Get`, `Set`, and `Clone` methods
- ✅ JSON support (std lib)
- ✅ YAML support (`gopkg.in/yaml.v3`)
- ✅ TOML support (`github.com/pelletier/go-toml/v2`)

---

## Installation

```sh
go get github.com/inovacc/packio
```

---

## Usage

You can construct a wrapper by passing an optional format selector:

```go
w := packio.New(MyObj{}, packio.JSON) // or packio.YAML / packio.TOML
// If you omit the second argument, it defaults to JSON:
defaultJSON := packio.New(MyObj{})
```

### JSON (built-in)

```go
package main

import (
  "fmt"
  "github.com/inovacc/packio"
)

type Person struct {
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
}

func main() {
  p := Person{FirstName: "Ada", LastName: "Lovelace"}

  // Wrap the value
  wrapped := packio.New(p)

  // Marshal to JSON
  jsonData, err := wrapped.Serialize()
  if err != nil {
    panic(err)
  }
  fmt.Printf("JSON: %s\n", string(jsonData))

  // Unmarshal back
  other := packio.New(Person{})
  if err := other.Deserialize(jsonData); err != nil {
    panic(err)
  }
  fmt.Printf("Unmarshaled: %+v\n", other.Get())
}
```

### YAML

```go
p := Person{FirstName: "Ada", LastName: "Lovelace"}
yw := packio.New(p, packio.YAML)
b, _ := yw.Serialize()
// ... later
otherY := packio.New(Person{}, packio.YAML)
_ = otherY.Deserialize(b)
```

### TOML

```go
p := Person{FirstName: "Ada", LastName: "Lovelace"}
tw := packio.New(p, packio.TOML)
b, _ := tw.Serialize()
// ... later
otherT := packio.New(Person{}, packio.TOML)
_ = otherT.Deserialize(b)
```

---

## Advanced Usage

#### Custom Types

```go
type CustomType struct {
  Field1 string
  Field2 int
}

wrapped := packio.New(CustomType{
  Field1: "value",
  Field2: 42,
})
```

#### Updating Data

```go
wrapped := packio.New(Person{})
wrapped.Set(Person{FirstName: "Grace", LastName: "Hopper"})
```

#### Cloning

```go
src := packio.New(Person{FirstName: "Linus", LastName: "Torvalds"})
fullCopy := src.Clone(false) // deep copy of data
emptyCopy := src.Clone(true) // zero-value data

// YAML/TOML wrappers also support Clone
_ = packio.New(src.Get(), packio.YAML).Clone(false)
_ = packio.New(src.Get(), packio.TOML).Clone(true)
```

#### Error Handling

```go
wrapped := packio.New(Person{})
if err := wrapped.Deserialize([]byte("invalid")); err != nil {
    fmt.Printf("Error deserializing: %v\n", err)
}
```

---

## API

### `Serializer[T any]` Interface (Unified)

```go
type Serializer[T any] interface {
  Serialize() ([]byte, error)
  Deserialize(data []byte) error
  Clone(empty bool) Serializer[T]
  Get() T
  Set(data T)
}
```

Program to this interface in your functions to enforce use of the serializer (any format), for example:

```go
func Save[T any](w packio.Serializer[T]) ([]byte, error) {
    return w.Serialize()
}
```

### Format-specific types

WithYAML[T] and WithTOML[T] are concrete types that implement Serializer[T] using YAML and TOML respectively. Prefer depending on the Serializer[T] interface in your APIs; instantiate the concrete
types when you need a specific format.

These allow you to accept interfaces in your APIs, ensuring callers pass the appropriate wrapper and enabling features like Clone and format-specific marshal/unmarshal.

### `WithJSON[T]` Type

```go
type WithJSON[T any] struct {
    Data T
}
```

### Additional types

- `WithYAML[T]`: `Serialize() ([]byte, error)`, `Deserialize([]byte) error`, `Get()`, `Set(T)`, `Clone(empty bool) Serializer[T]`.
- `WithTOML[T]`: `Serialize() ([]byte, error)`, `Deserialize([]byte) error`, `Get()`, `Set(T)`, `Clone(empty bool) Serializer[T]`.

---

## Notes on Concurrency

This package does not include synchronization primitives. If you share a wrapper instance across goroutines and perform concurrent writes, protect access with your own sync (e.g., `sync.Mutex`,
`sync.RWMutex`).

---

## Contributing

PRs are welcome! Please open issues or submit a Pull Request if you have improvements or fixes.

---

## License

Licensed under the MIT License.
