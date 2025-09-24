package randomness

import "math/rand"

// ------------------------------------------------------------------------------------------------
func GenerateRandomNumbersInt(max, howMany int) []int {
	nums := make([]int, howMany)

	for pos := range howMany {
		nums[pos] = rand.Intn(max)
	}

	return nums
}

// ------------------------------------------------------------------------------------------------
func GenerateRandomNumbersUInt8(max, howMany int) []uint8 {
	nums := make([]uint8, howMany)

	for pos := range howMany {
		nums[pos] = uint8(rand.Intn(max))
	}

	return nums
}
