package helperfunctions

import (
	"math"
	"testing"
)

// TestCalculateMean tests the CalculateMean function
func TestCalculateMean(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{
			name:     "Empty slice",
			data:     []float64{},
			expected: 0,
		},
		{
			name:     "Single element",
			data:     []float64{5.0},
			expected: 5.0,
		},
		{
			name:     "Multiple elements",
			data:     []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			expected: 3.0,
		},
		{
			name:     "Negative numbers",
			data:     []float64{-1.0, -2.0, -3.0, -4.0, -5.0},
			expected: -3.0,
		},
		{
			name:     "Mixed numbers",
			data:     []float64{1.0, -1.0, 0.0, 2.0, -2.0},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateMean(tt.data)
			if !floatEquals(result, tt.expected) {
				t.Errorf("CalculateMean() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// TestStandardDeviation tests the StandardDeviation function
func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{
			name:     "Empty slice",
			data:     []float64{},
			expected: 0,
		},
		{
			name:     "Single element",
			data:     []float64{5.0},
			expected: 0,
		},
		{
			name:     "Multiple elements",
			data:     []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			expected: 1.4142135623730951, // sqrt(2)
		},
		{
			name:     "Negative numbers",
			data:     []float64{-1.0, -2.0, -3.0, -4.0, -5.0},
			expected: 1.4142135623730951, // sqrt(2)
		},
		{
			name:     "Mixed numbers",
			data:     []float64{1.0, -1.0, 0.0, 2.0, -2.0},
			expected: 1.4142135623730951, // sqrt(2)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StandardDeviation(tt.data)
			if !floatEquals(result, tt.expected) {
				t.Errorf("StandardDeviation() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// Helper function to compare float64 values with a tolerance
func floatEquals(a, b float64) bool {
	const tolerance = 1e-9
	return math.Abs(a-b) < tolerance
}
