package obj

import (
	"testing"
)

func TestToStringAsArray(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		sep      string
		quote    bool
		expected string
	}{
		{"EmptyInput", nil, ",", false, ""},
		{"StringArrayInput", []interface{}{"apple", "orange", "banana"}, ",", false, "apple,orange,banana"},
		{"MixedTypeArrayInput", []interface{}{"apple", 42, "banana"}, ",", false, "apple,banana"},
		{"SpaceAsSeparator", []interface{}{"apple", "orange", "banana"}, " ", false, "apple orange banana"},
		{"SpaceAsSeparato and quote", []interface{}{"apple", "orange", "banana"}, " ", true, "'apple' 'orange' 'banana'"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToStringAsArray(tc.input, tc.sep, tc.quote)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
