package calculations

import (
	"math"
	"testing"
)

func TestFindAverage(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := 3.0

	result := FindAverage(input)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestFindMedian(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := 3.0

	result := FindMedian(input)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestFindVariance(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := 2.0

	result := FindVariance(input)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestFindStandardDeviation(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := 1.41421356237
	tolerance := 1e-9

	result := math.Sqrt(FindVariance(input))
	if math.Abs(result-expected) > tolerance {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
