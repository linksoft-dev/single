package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInputParser peform test over inputparser
func TestInputParser(t *testing.T) {
	testCases := []struct {
		expected, model Model
	}{
		{model: NewModel("test", "EMAIL@TEST.COM"), expected: Model{Name: "TEST", Email: "email@test.com"}},
		{model: NewModel("test  with  extra spaces", "email@TEST.COM"), expected: Model{Name: "TEST WITH EXTRA SPACES", Email: "email@test.com"}},
	}
	for _, testCase := range testCases {
		testname := fmt.Sprintf("model: '%v'", testCase.model)
		t.Run(testname, func(t *testing.T) {
			inputParser(&testCase.model)
			assert.Equal(t, testCase.expected, testCase.model)
		})
	}
}
