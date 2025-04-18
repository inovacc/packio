# JSON Wrapper Package [![Test](https://github.com/dyammarcano/wrapper/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/dyammarcano/wrapper/actions/workflows/test.yml)

A generic Go package that provides a simple way to add JSON marshaling and unmarshaling capabilities to any struct type.
This wrapper is particularly useful when you need to add JSON functionality to existing structs without modifying them
directly.

## Features

- ✅ Generic implementation using Go 1.21+ type parameters
- ✅ Simple and intuitive API
- ✅ Thread-safe for concurrent reading
- ✅ Built-in `Get` and `Set` methods
- ✅ Zero external dependencies

---

## Installation

```sh
go get github.com/inovacc/wrapper
```

---

## Usage

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/inovacc/wrapper"
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	person := gofakeit.Person()

	// Wrap the person
	wrapped := wrapper.NewWrapper(person)

	// Marshal to JSON
	jsonData, err := wrapped.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON: %s\n", string(jsonData))

	// Unmarshal from JSON
	newWrapped := wrapper.NewWrapper(gofakeit.PersonInfo{})
	if err = newWrapped.UnmarshalJSON(jsonData); err != nil {
		panic(err)
	}

	result := newWrapped.Get() // is type PersonInfo
	fmt.Printf("Unmarshaled: %+v\n", result)
}
```

---

### Advanced Usage

#### Custom Types

```go
type CustomType struct {
	Field1 string
	Field2 int
}

wrapped := wrapper.NewWrapper(CustomType{
	Field1: "value",
	Field2: 42,
})
```

#### Updating Data

```go
wrapped := wrapper.NewWrapper(User{})
newUser := User{Name: "Updated Name"}
wrapped.Set(newUser)
```

#### Error Handling

```go
wrapped := wrapper.NewWrapper(User{})
if err := wrapped.UnmarshalJSON([]byte(`invalid json`)); err != nil {
	fmt.Printf("Error unmarshaling JSON: %v\n", err)
}
```

---

## Interface

### `Wrapper` Interface

```go
type Wrapper interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}
```

### `WithJSON[T]` Type

```go
type WithJSON[T any] struct {
	Data T
}
```

---

## Methods

- `NewWrapper[T any](data T) *WithJSON[T]`: Creates a new wrapper
- `MarshalJSON() ([]byte, error)`: Converts data to JSON
- `UnmarshalJSON(data []byte) error`: Parses JSON into data
- `Get() T`: Retrieves the stored value
- `Set(data T)`: Updates the value

---

## Best Practices

1. Always check for errors on (un)marshaling
2. Use `json` tags for correct serialization
3. Always use `NewWrapper` to initialize
4. Explicitly define types for clarity

---

## Thread Safety

- Safe for concurrent **read** operations
- Use sync primitives (e.g., `sync.Mutex`) for **concurrent writes**

---

## Contributing

PRs are welcome! Please open issues or submit a Pull Request if you have improvements or fixes.

---

## License

Licensed under the [MIT License](./LICENSE)
