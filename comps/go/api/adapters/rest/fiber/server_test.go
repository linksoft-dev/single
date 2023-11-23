package fiber

import "testing"

func TestConvertBraceToColon(t *testing.T) {
	tests := []struct {
		Name     string
		Input    string
		Expected string
	}{
		{"TestWithParam", "{param}", ":param"},
		{"TestWithPathParam", "path/{param}/value", "path/:param/value"},
		{"TestWithMultipleBraces", "{a}/{b}/{c}", ":a/:b/:c"},
		{"TestWithoutBraces", "no_braces", "no_braces"},
		{"TestEmptyInput", "", ""},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := convertBraceToColon(test.Input)
			if result != test.Expected {
				t.Errorf("Input: %s, Expected: %s, Got: %s", test.Input, test.Expected, result)
			}
		})
	}
}
