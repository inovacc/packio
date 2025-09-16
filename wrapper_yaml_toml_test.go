package wrapper

import (
	"reflect"
	"testing"
)

func TestYAMLWrapperRoundTrip(t *testing.T) {
	u := User{
		Name:        "Ada",
		Description: "Pioneer",
		Categories:  []string{"math", "computing"},
		Price:       3.14,
		Features:    []string{"analytical engine"},
		Color:       "blue",
		Material:    "paper",
	}
	w := NewYAMLWrapper(u)
	b, err := w.Serialize()
	if err != nil {
		t.Fatalf("Serialize error: %v", err)
	}
	other := NewYAMLWrapper(User{})
	if err := other.Deserialize(b); err != nil {
		t.Fatalf("Deserialize error: %v", err)
	}
	if !reflect.DeepEqual(u, other.Get()) {
		t.Errorf("YAML round-trip mismatch\nwant: %+v\n got: %+v", u, other.Get())
	}
}

func TestYAMLCloneDeepCopy(t *testing.T) {
	orig := NewYAMLWrapper(User{Categories: []string{"a", "b"}})
	clone := orig.Clone(false)
	origData := orig.Get()
	origData.Categories[0] = "x"
	cloned := clone.Get()
	if cloned.Categories[0] == "x" {
		t.Error("YAML clone did not deep copy slices")
	}
}

func TestTOMLWrapperRoundTrip(t *testing.T) {
	u := User{
		Name:        "Grace",
		Description: "COBOL",
		Categories:  []string{"lang"},
		Price:       1.23,
		Features:    []string{"navy"},
		Color:       "white",
		Material:    "paper",
	}
	w := NewTOMLWrapper(u)
	b, err := w.Serialize()
	if err != nil {
		t.Fatalf("Serialize error: %v", err)
	}
	other := NewTOMLWrapper(User{})
	if err := other.Deserialize(b); err != nil {
		t.Fatalf("Deserialize error: %v", err)
	}
	if !reflect.DeepEqual(u, other.Get()) {
		t.Errorf("TOML round-trip mismatch\nwant: %+v\n got: %+v", u, other.Get())
	}
}

func TestTOMLCloneDeepCopy(t *testing.T) {
	orig := NewTOMLWrapper(User{Features: []string{"f1", "f2"}})
	clone := orig.Clone(false)
	origData := orig.Get()
	origData.Features[0] = "y"
	cloned := clone.Get()
	if cloned.Features[0] == "y" {
		t.Error("TOML clone did not deep copy slices")
	}
}
