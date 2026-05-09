package db

import (
	"testing"
)

// TestNewValue tests that NewValue correctly detects the type of a raw string.
func TestNewValue(t *testing.T) {
	tests := []struct {
		raw      string
		wantType ValueType
	}{
		// Integers
		{"42", TypeInteger},
		{"-1", TypeInteger},
		{"0", TypeInteger},
		// Floats
		{"3.14", TypeFloat},
		{"-0.5", TypeFloat},
		// Booleans
		{"true", TypeBoolean},
		{"false", TypeBoolean},
		// JSON objects and arrays
		{`{"key":"value"}`, TypeJSON},
		{`[1,2,3]`, TypeJSON},
		// Plain strings
		{"hello", TypeString},
		{"", TypeString},
		{"hello world", TypeString},
	}
	for _, test := range tests {
		got := NewValue(test.raw)
		if got.Type != test.wantType {
			t.Errorf("NewValue(%q).Type = %q; want %q", test.raw, got.Type, test.wantType)
		}
		if got.Raw != test.raw {
			t.Errorf("NewValue(%q).Raw = %q; want %q", test.raw, got.Raw, test.raw)
		}
	}
}
