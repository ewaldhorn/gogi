package randomness

import (
	"testing"
)

// ------------------------------------------------------------------------------------------------
func TestGenerateRandomNumbersInt(t *testing.T) {
	max := 100
	howMany := 10
	nums := GenerateRandomNumbersInt(max, howMany)

	if len(nums) != howMany {
		t.Errorf("Expected slice of length %d, but got %d", howMany, len(nums))
	}

	for _, num := range nums {
		if num < 0 || num >= max {
			t.Errorf("Expected number to be in range [0, %d), but got %d", max, num)
		}
	}
}

// ------------------------------------------------------------------------------------------------
func TestGenerateRandomNumbersUInt8(t *testing.T) {
	max := 255
	howMany := 20
	nums := GenerateRandomNumbersUInt8(max, howMany)

	if len(nums) != howMany {
		t.Errorf("Expected slice of length %d, but got %d", howMany, len(nums))
	}

	for _, num := range nums {
		// The cast to int is important for the check, as a uint8 can't be >= 256
		if int(num) >= max {
			t.Errorf("Expected number to be in range [0, %d), but got %d", max, num)
		}
	}
}
