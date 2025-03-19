package reverse

import (
	"testing"
)

// TestReverse tests the Reverse function.
func TestReverse(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},   // edge case: empty string
		{"a", "a"}, // edge case: single character
		{"12345", "54321"},
		{"Go", "oG"},
	}

	for _, c := range cases {
		result := Reverse(c.input)
		if result != c.expected {
			t.Errorf("Reverse(%q) == %q, expected %q", c.input, result, c.expected)
		}
	}
}
