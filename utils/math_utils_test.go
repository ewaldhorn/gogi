package utils

import (
	"math"
	"testing"
)

// ------------------------------------------------------------------------------------------------
func TestClampUInt8(t *testing.T) {
	tests := []struct {
		initial  uint8
		diff     int
		expected uint8
		name     string
	}{
		{initial: 100, diff: 50, expected: 150, name: "Add within range"},
		{initial: 10, diff: -5, expected: 5, name: "Subtract within range"},
		{initial: 5, diff: -10, expected: 0, name: "Clamp to 0 (negative result)"},
		{initial: 0, diff: -100, expected: 0, name: "Clamp to 0 (already 0, negative diff)"},
		{initial: 200, diff: 100, expected: 255, name: "Clamp to 255 (positive result)"},
		{initial: 250, diff: 10, expected: 255, name: "Clamp to 255 (near max, positive diff)"},
		{initial: 255, diff: 5, expected: 255, name: "Clamp to 255 (already max, positive diff)"},
		{initial: 128, diff: 0, expected: 128, name: "Zero diff"},
		{initial: 0, diff: 0, expected: 0, name: "Zero initial, zero diff"},
		{initial: 255, diff: 0, expected: 255, name: "Max initial, zero diff"},
		{initial: 255, diff: -10, expected: 245, name: "Subtract from max"},
		{initial: 0, diff: 50, expected: 50, name: "Add to zero"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClampUInt8(tt.initial, tt.diff)
			if got != tt.expected {
				t.Errorf("ClampUInt8(%d, %d) = %d; want %d", tt.initial, tt.diff, got, tt.expected)
			}
		})
	}
}

// ------------------------------------------------------------------------------------------------
func TestClampIntTo(t *testing.T) {
	tests := []struct {
		initial  int
		min      int
		max      int
		expected int
		name     string
	}{
		{initial: 50, min: 0, max: 100, expected: 50, name: "Within range"},
		{initial: -10, min: 0, max: 100, expected: 0, name: "Below min"},
		{initial: 110, min: 0, max: 100, expected: 100, name: "Above max"},
		{initial: 0, min: 0, max: 100, expected: 0, name: "At min"},
		{initial: 100, min: 0, max: 100, expected: 100, name: "At max"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClampIntTo(tt.initial, tt.min, tt.max)
			if got != tt.expected {
				t.Errorf("ClampIntTo(%d, %d, %d) = %d; want %d", tt.initial, tt.min, tt.max, got, tt.expected)
			}
		})
	}
}

// ------------------------------------------------------------------------------------------------
func TestMapValue(t *testing.T) {
	tests := []struct {
		value    float64
		inMin    float64
		inMax    float64
		outMin   float64
		outMax   float64
		expected float64
		name     string
	}{
		{value: 50, inMin: 0, inMax: 100, outMin: 0, outMax: 1, expected: 0.5, name: "Simple map"},
		{value: 0, inMin: -1, inMax: 1, outMin: 0, outMax: 100, expected: 50, name: "Negative input range"},
		{value: 5, inMin: 0, inMax: 10, outMin: 10, outMax: 20, expected: 15, name: "Shifted output range"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapValue(tt.value, tt.inMin, tt.inMax, tt.outMin, tt.outMax)
			if got != tt.expected {
				t.Errorf("MapValue() = %f; want %f", got, tt.expected)
			}
		})
	}
}

// ------------------------------------------------------------------------------------------------
func TestRandomSignFloat(t *testing.T) {
	n := 123.45
	result := RandomSignFloat(n)
	if math.Abs(result) != n {
		t.Errorf("Expected absolute value of result to be %f, but got %f", n, math.Abs(result))
	}
}

// ------------------------------------------------------------------------------------------------
func TestRandomSignInt(t *testing.T) {
	n := 123
	result := RandomSignInt(n)
	if result != n && result != -n {
		t.Errorf("Expected result to be %d or %d, but got %d", n, -n, result)
	}
}
