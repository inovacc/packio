package wrapper

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create wrapper
			p := NewWrapper(tt.input)

			// Test MarshalJSON
			data, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Validate JSON structure
			var rawMap map[string]interface{}
			if err := json.Unmarshal(data, &rawMap); err != nil {
				t.Errorf("Generated JSON is invalid: %v", err)
				return
			}

			// Test UnmarshalJSON
			newWrapper := NewWrapper(User{})
			if err := newWrapper.UnmarshalJSON(data); err != nil {
				t.Errorf("UnmarshalJSON() error = %v", err)
				return
			}

			// Compare original and unmarshaled results
			result := newWrapper.Get()
			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("Data mismatch after marshal/unmarshal\ngot: %+v\nwant: %+v", result, tt.input)
			}
		})
	}
}

func TestWrapperEdgeCases(t *testing.T) {
	t.Run("Unmarshal invalid JSON", func(t *testing.T) {
		w := NewWrapper(User{})
		err := w.UnmarshalJSON([]byte(`{"invalid json`))
		if err == nil {
			t.Error("Expected error for invalid JSON, got nil")
		}
	})

	t.Run("Unmarshal empty JSON", func(t *testing.T) {
		w := NewWrapper(User{})
		err := w.UnmarshalJSON([]byte(`{}`))
		if err != nil {
			t.Errorf("Unexpected error for empty JSON: %v", err)
		}
	})
}
