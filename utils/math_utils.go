package utils

import "math/rand/v2"

// ------------------------------------------------------------------------------------------------
// RandomSignFloat returns n or -n with a roughly 50% chance.
func RandomSignFloat(n float64) float64 {
	if rand.IntN(2) == 0 {
		return -n
	}
	return n
}

// ------------------------------------------------------------------------------------------------
// RandomSignInt returns n or -n with a roughly 50% chance.
func RandomSignInt(n int) int {
	if rand.IntN(2) == 0 {
		return -n
	}
	return n
}

// ------------------------------------------------------------------------------------------------
func ClampUInt8(initial uint8, diff int) uint8 {
	newValue := int(initial) + diff

	if newValue < 0 {
		return 0
	}

	if newValue > 255 {
		return 255
	}

	return uint8(newValue)
}

// ------------------------------------------------------------------------------------------------
func ClampIntTo(initial int, min int, max int) int {
	if initial < min {
		return min
	}

	if initial > max {
		return max
	}

	return initial
}

// ------------------------------------------------------------------------------------------------
// MapValue maps a value from one range to another.
func MapValue(value, inMin, inMax, outMin, outMax float64) float64 {
	return (value-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}

/*
// TestClampUInt8 contains test cases for the ClampUInt8 function.
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

*/
