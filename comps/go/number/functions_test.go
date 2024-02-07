package number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxFloat64(t *testing.T) {
	testCases := []struct {
		value    []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 10, 5, 3, 1000, 1001}, 1001},
		{[]float64{5885.588997, 0.5566, 5885.588998}, 5885.588998},
		{[]float64{0, 10, 10, 36}, 36},
		{[]float64{100, 10, 10, 36}, 100},
		{[]float64{100.5, 10, 10, 36}, 100.5},
		{[]float64{-1, -1, 1.0001, 1.00001, 0}, 1.0001},
	}
	for _, v := range testCases {
		got := Max(v.value...)
		assert.Equal(t, v.expected, got, "", v.value, v.expected, got)
	}
}

func TestMaxInt(t *testing.T) {
	testCases := []struct {
		value    []int
		expected int
	}{
		{[]int{1, 2, 3, 10, 5, 3, 1000, 1001}, 1001},
		{[]int{5885, 0, 5885}, 5885},
		{[]int{10, 11, 9}, 11},
	}
	for _, v := range testCases {
		got := Max(v.value...)
		assert.Equal(t, v.expected, got, "", v.value, v.expected, got)
	}
}

func TestMinFloat64(t *testing.T) {
	testCases := []struct {
		value    []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 10, 5, 3, 1000, 1001}, 1},
		{[]float64{5885.588997, 0.5566, 5885.588998}, 0.5566},
		{[]float64{0, 10, 10, 36}, 0},
		{[]float64{100.5, 10, 10, 36}, 10},
		{[]float64{-1, 0, 0, 1}, -1},
		{[]float64{-1, -1, -1.0001, 0}, -1.0001},
		{[]float64{-1, -1, -1.0001, -1.00001, 0}, -1.0001},
	}
	for _, v := range testCases {
		got := Min(v.value...)
		assert.Equal(t, v.expected, got, "", v.value, v.expected, got)
	}
}

func TestMinInt(t *testing.T) {
	testCases := []struct {
		value    []int
		expected int
	}{
		{[]int{1, 2, 3, 10, 5, 3, 1000, 1001}, 1},
		{[]int{5885, 0, 5885}, 0},
		{[]int{10, 11, 9}, 9},
		{[]int{-5, -6, 5, 0, 0, 8}, -6},
	}
	for _, v := range testCases {
		got := Min(v.value...)
		assert.Equal(t, v.expected, got, "", v.value, v.expected, got)
	}
}

func TestRoundFloat(t *testing.T) {
	testCases := []struct {
		Name      string
		Value     float64
		Precision int
		Expected  float64
	}{
		{Name: "Pi with 2 decimal places", Value: 3.14159, Precision: 2, Expected: 3.14},
		{Name: "Float with 1 decimal place", Value: 6.666, Precision: 1, Expected: 6.7},
		{Name: "Large float with no decimal places", Value: 123456.789, Precision: 0, Expected: 123456},
		{Name: "Float with 3 decimal places", Value: 0.987654, Precision: 3, Expected: 0.988},
		{Name: "Float with negative precision", Value: 8.999, Precision: -1, Expected: 9.0}, // Precision default to 1
		{Name: "Integer with no decimal places", Value: 42, Precision: 0, Expected: 42.0},
		{Name: "Float with 1 decimal place", Value: 7.1, Precision: 1, Expected: 7.1},
		{Name: "Float with 2 decimal places", Value: 99.99, Precision: 2, Expected: 99.99},
		{Name: "Zero with 2 decimal places", Value: 0, Precision: 2, Expected: 0},
		{Name: "Negative float with 1 decimal place", Value: -5.678, Precision: 1, Expected: -5.7},
		{Name: "Negative float with 2 decimal places", Value: -123.45, Precision: 2, Expected: -123.45},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := RoundFloat(tc.Value, tc.Precision)
			assert.Equal(t, tc.Expected, result, "Expected %v, but got %v", tc.Expected, result)
		})
	}
}
