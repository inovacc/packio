package packio

import (
	"encoding/json"
	"reflect"
	"testing"
)

type User struct {
	Name        string   `json:"name" xml:"name"`
	Description string   `json:"description" xml:"description"`
	Categories  []string `json:"categories" xml:"categories"`
	Price       float64  `json:"price" xml:"price"`
	Features    []string `json:"features" xml:"features"`
	Color       string   `json:"color" xml:"color"`
	Material    string   `json:"material" xml:"material"`
}

func TestWrapper(t *testing.T) {
	tests := []struct {
		name    string
		input   User
		wantErr bool
	}{
		{
			name: "Valid full user data",
			input: User{
				Name:        "Microwave Vertex Marble",
				Description: "Full him bale me within. As far to canoe wad its it.",
				Categories:  []string{"musical instruments", "bicycles and accessories", "books"},
				Price:       46.06,
				Features:    []string{"user-friendly", "compact"},
				Color:       "navy",
				Material:    "granite",
			},
			wantErr: false,
		},
		{
			name: "Empty categories and features",
			input: User{
				Name:        "Simple Product",
				Description: "Basic description",
				Categories:  []string{},
				Price:       19.99,
				Features:    []string{},
				Color:       "red",
				Material:    "plastic",
			},
			wantErr: false,
		},
		{
			name: "Zero price",
			input: User{
				Name:        "Free Item",
				Description: "Free product description",
				Categories:  []string{"free"},
				Price:       0.0,
				Features:    []string{"free"},
				Color:       "white",
				Material:    "paper",
			},
			wantErr: false,
		},
		{
			name: "Special characters in strings",
			input: User{
				Name:        "Product!@#$%^&*()",
				Description: "Description with šĕęćīàł characters 你好",
				Categories:  []string{"category#1", "category@2"},
				Price:       99.99,
				Features:    []string{"feature!1", "feature@2"},
				Color:       "blue-green",
				Material:    "mixed/material",
			},
			wantErr: false,
		},
		{
			name: "Maximum float value",
			input: User{
				Name:        "Expensive Product",
				Description: "Very expensive item",
				Categories:  []string{"luxury"},
				Price:       1.797693134862315e+308, // Max float64
				Features:    []string{"expensive"},
				Color:       "gold",
				Material:    "diamond",
			},
			wantErr: false,
		},
		{
			name: "Negative price",
			input: User{
				Name:        "Invalid Product",
				Description: "Product with negative price",
				Categories:  []string{"test"},
				Price:       -1.0,
				Features:    []string{"test"},
				Color:       "red",
				Material:    "plastic",
			},
			wantErr: false,
		},
		{
			name: "Empty required fields",
			input: User{
				Name:        "",
				Description: "",
				Categories:  nil,
				Price:       0.0,
				Features:    nil,
				Color:       "",
				Material:    "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create wrapper
			p := New(tt.input)

			// Test Serialize (JSON)
			data, err := p.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Validate JSON structure
			var rawMap map[string]interface{}
			if err := json.Unmarshal(data, &rawMap); err != nil {
				t.Errorf("Generated JSON is invalid: %v", err)
				return
			}

			// Test Deserialize
			newWrapper := New(User{})
			if err := newWrapper.Deserialize(data); err != nil {
				t.Errorf("Deserialize() error = %v", err)
				return
			}

			// Compare original and unmarshalled
			result := newWrapper.Get()
			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("Data mismatch after marshal/unmarshal\ngot: %+v\nwant: %+v", result, tt.input)
			}
		})
	}
}

func TestWrapperEdgeCases(t *testing.T) {
	t.Run("Deserialize invalid JSON", func(t *testing.T) {
		w := New(User{})

		err := w.Deserialize([]byte(`{"invalid json`))
		if err == nil {
			t.Error("Expected error for invalid JSON, got nil")
		}
	})

	t.Run("Deserialize empty JSON", func(t *testing.T) {
		w := New(User{})

		err := w.Deserialize([]byte(`{}`))
		if err != nil {
			t.Errorf("Unexpected error for empty JSON: %v", err)
		}
	})

	t.Run("Deserialize with invalid types", func(t *testing.T) {
		w := New(User{})

		err := w.Deserialize([]byte(`{"price": "not a number"}`))
		if err == nil {
			t.Error("Expected error for invalid type conversion, got nil")
		}
	})

	t.Run("Deserialize with null values", func(t *testing.T) {
		w := New(User{})

		err := w.Deserialize([]byte(`{"name": null, "price": null}`))
		if err != nil {
			t.Errorf("Unexpected error for null values: %v", err)
		}
	})
}

func TestCloneFull(t *testing.T) {
	wUser := New(User{
		Name:        "Microwave Vertex Marble",
		Description: "Full him bale me within. As far to canoe wad its it.",
		Categories:  []string{"musical instruments", "bicycles and accessories", "books"},
		Price:       46.06,
		Features:    []string{"user-friendly", "compact"},
		Color:       "navy",
		Material:    "granite",
	})

	clone := wUser.Clone(true)
	expected := clone.Get().Name

	if expected != "" {
		t.Error("Cloned wrapper empty")
	}
}

func TestCloneEmpty(t *testing.T) {
	wUser := New(User{
		Name:        "Microwave Vertex Marble",
		Description: "Full him bale me within. As far to canoe wad its it.",
		Categories:  []string{"musical instruments", "bicycles and accessories", "books"},
		Price:       46.06,
		Features:    []string{"user-friendly", "compact"},
		Color:       "navy",
		Material:    "granite",
	})

	clone := wUser.Clone(false)
	expected := clone.Get().Name

	if expected != wUser.Get().Name {
		t.Error("Cloned wrapper need to be empty")
	}
}

func TestCloneDeepCopy(t *testing.T) {
	original := New(User{
		Categories: []string{"cat1", "cat2"},
		Features:   []string{"feat1", "feat2"},
	})

	clone := original.Clone(false)

	// Modify original's slices
	originalUser := original.Get()
	originalUser.Categories[0] = "modified"
	originalUser.Features[0] = "modified"

	// Check if clone's slices were affected
	clonedUser := clone.Get()
	if clonedUser.Categories[0] == "modified" || clonedUser.Features[0] == "modified" {
		t.Error("Clone did not perform deep copy of slices")
	}
}

func TestEmptyWrapperWithSet(t *testing.T) {
	// Create an empty wrapper
	emptyWrapper := New(User{})

	// Prepare test data
	testUser := User{
		Name:        "Test User",
		Description: "Test Description",
		Categories:  []string{"test1", "test2"},
		Price:       29.99,
		Features:    []string{"feature1", "feature2"},
		Color:       "blue",
		Material:    "metal",
	}

	// Set the data to an empty wrapper
	emptyWrapper.Set(testUser)

	// Verify the data was set correctly
	result := emptyWrapper.Get()
	if !reflect.DeepEqual(result, testUser) {
		t.Errorf("Set() failed to update empty wrapper\ngot: %+v\nwant: %+v", result, testUser)
	}

	// Test serialization of the updated wrapper
	data, err := emptyWrapper.Serialize()
	if err != nil {
		t.Errorf("Serialize() error after Set(): %v", err)
	}

	// Verify JSON structure
	var rawMap map[string]interface{}
	if err := json.Unmarshal(data, &rawMap); err != nil {
		t.Errorf("Generated JSON is invalid after Set(): %v", err)
	}

	// Create another empty wrapper and deserialize the data
	verifyWrapper := New(User{})
	if err := verifyWrapper.Deserialize(data); err != nil {
		t.Errorf("Deserialize() error after Set(): %v", err)
	}

	// Compare the results
	if !reflect.DeepEqual(verifyWrapper.Get(), testUser) {
		t.Errorf("Data mismatch after Set() and marshal/unmarshal\ngot: %+v\nwant: %+v",
			verifyWrapper.Get(), testUser)
	}
}

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
	w := New(u, YAML)

	b, err := w.Serialize()
	if err != nil {
		t.Fatalf("Serialize error: %v", err)
	}

	other := New(User{}, YAML)
	if err := other.Deserialize(b); err != nil {
		t.Fatalf("Deserialize error: %v", err)
	}

	if !reflect.DeepEqual(u, other.Get()) {
		t.Errorf("YAML round-trip mismatch\nwant: %+v\n got: %+v", u, other.Get())
	}
}

func TestYAMLCloneDeepCopy(t *testing.T) {
	orig := New(User{Categories: []string{"a", "b"}}, YAML)
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
	w := New(u, TOML)

	b, err := w.Serialize()
	if err != nil {
		t.Fatalf("Serialize error: %v", err)
	}

	other := New(User{}, TOML)
	if err := other.Deserialize(b); err != nil {
		t.Fatalf("Deserialize error: %v", err)
	}

	if !reflect.DeepEqual(u, other.Get()) {
		t.Errorf("TOML round-trip mismatch\nwant: %+v\n got: %+v", u, other.Get())
	}
}

func TestTOMLCloneDeepCopy(t *testing.T) {
	orig := New(User{Features: []string{"f1", "f2"}}, TOML)
	clone := orig.Clone(false)
	origData := orig.Get()
	origData.Features[0] = "y"

	cloned := clone.Get()
	if cloned.Features[0] == "y" {
		t.Error("TOML clone did not deep copy slices")
	}
}
