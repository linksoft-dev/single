package validation

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestIsDateTime(t *testing.T) {
	expectedResults := map[string]bool{
		"2020-12-01":          true,
		"2020-01-12":          true,
		"2020-12-01 10:34:57": true,
		"30/01/2020 22:00:57": true,
		"33/01/2020 22:00:57": false,
	}
	validation := NewValidation(language.Portuguese)
	for date, expected := range expectedResults {
		if validation.IsDateTime(date) != expected {
			t.Errorf("Fail test for string date %s, expected : %t got: %t ", date, !expected, expected)
		}
	}
}

func TestIsIn(t *testing.T) {
	type inputAndOutput struct {
		input      string
		IsInValues []string
		output     bool
	}
	expectedResults := []inputAndOutput{
		{"A", []string{"A", "B", "C"}, true},
	}
	validation := NewValidation(language.Portuguese)
	for _, value := range expectedResults {
		got := validation.IsIn("fieldName", value.input, value.IsInValues...)
		assert.Equal(t, value.output, got, "IsIn of the value '%f' is expected '%f' but got '%f'",
			value.input, value.output, got)
	}
}

func TestIsCreditCardNumber(t *testing.T) {
	type inputAndOutput struct {
		input  string
		output bool
	}
	expectedResults := []inputAndOutput{
		{"5366042384178249", true},
	}
	validation := NewValidation(language.Portuguese)
	for _, value := range expectedResults {
		got := validation.IsCreditCardNumber("fieldName", value.input)
		assert.Equal(t, value.output, got, "IsIn of the value '%f' is expected '%f' but got '%f'",
			value.input, value.output, got)
	}
}
