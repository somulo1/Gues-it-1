package helperfunctions

import (
	"math"
	"testing"
)

// Helper function to compare floats within a tolerance
func nearlyEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

// Test for CalculateMean
func TestCalculateMean(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	expectedMean := 30.0

	result := CalculateMean(data)

	if !nearlyEqual(result, expectedMean, 1e-9) {
		t.Errorf("CalculateMean() = %v; want %v", result, expectedMean)
	}
}

// Test for StandardDeviation
func TestStandardDeviation(t *testing.T) {
	data := []float64{10, 20, 30, 40, 50}
	// Expected standard deviation value calculated manually or using a trusted tool
	expectedStdDev := 15.81

	result := StandardDeviation(data)

	if !nearlyEqual(result, expectedStdDev, 1e-2) {
		t.Errorf("StandardDeviation() = %v; want %v", result, expectedStdDev)
	}
}

// Edge case: empty data for StandardDeviation
func TestStandardDeviationEmpty(t *testing.T) {
	data := []float64{}
	expectedStdDev := 0.0

	result := StandardDeviation(data)

	if !nearlyEqual(result, expectedStdDev, 1e-9) {
		t.Errorf("StandardDeviation() with empty data = %v; want %v", result, expectedStdDev)
	}
}

// Edge case: single data point for StandardDeviation
func TestStandardDeviationSingle(t *testing.T) {
	data := []float64{10}
	expectedStdDev := 0.0

	result := StandardDeviation(data)

	if !nearlyEqual(result, expectedStdDev, 1e-9) {
		t.Errorf("StandardDeviation() with single data point = %v; want %v", result, expectedStdDev)
	}
}
