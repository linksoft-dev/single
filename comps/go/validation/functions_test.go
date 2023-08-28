package validation

import (
	"testing"
)

// TestIsValidEan tests the IsValidEan function
func TestIsValidEan(t *testing.T) {
	expectedResults := map[string]bool{
		"40700719670720": true,  //ean14
		"5012345678900":  true,  //ean13
		"012345678905":   true,  //ean12
		"78912342":       true,  //ean8
		"":               true,  //EMPTY ean
		"40700719670721": false, //INVALID ean14
		"5012345678901":  false, //INVALID ean13
		"012345678906":   false, //INVALID ean12
		"78912343":       false, //INVALID ean8
		"SEM GTIM":       false, //INVALID ean
		"1":              false, //INVALID ean
	}
	for ean, expected := range expectedResults {
		if IsValidEan(ean) != expected {
			t.Errorf("Fail test for EAN %s, expected : %t got: %t ", ean, !expected, expected)
		}
	}
}

func TestIsCnpjRelated(t *testing.T) {
	type testCnpjs struct {
		cnpj1    string
		cnpj2    string
		expected bool
	}

	tests := []testCnpjs{
		{"09015016000142", "09015016000223", true},
		{"09015016000223", "09015016000142", true},
		{"09015016000142", "08869363000250", false},
		{"08869363000250", "09015016000142", false},
		{"0901", "09015016000142", false},
		{"09015016000142", "0901", false},
		{"", "09015016000223", false},
		{"09015016000142", "", false},
		{"", "", false},
	}
	for _, test := range tests {
		if IsCnpjRelated(test.cnpj1, test.cnpj2) != test.expected {
			t.Errorf("Fail test for (%s) and (%s), expected : %t got: %t ",
				test.cnpj1, test.cnpj2, !test.expected, test.expected)
		}
	}
}
