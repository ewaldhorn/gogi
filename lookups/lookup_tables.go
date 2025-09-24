// Package lookups provides lookup tables for trigonometric functions.
// This is useful for animations where absolute accuracy is less important than raw speed.
package lookups

// There are basic Get functions and then more advanced Interpolating Get functions for when
// accuracy matters a bit more.

import "math"

// ------------------------------------------------------------------------------------------------
const (
	LOOKUP_TABLE_SIZE = 4096
	MAX_ANGLE         = 2 * math.Pi
)

// ------------------------------------------------------------------------------------------------
type LookupTables struct {
	sineTable, cosineTable []float64
}

// ------------------------------------------------------------------------------------------------
func NewLookupTables() *LookupTables {
	return &LookupTables{
		sineTable:   createSinLookupTable(LOOKUP_TABLE_SIZE),
		cosineTable: createCosLookupTable(LOOKUP_TABLE_SIZE),
	}

	// for i := range LOOKUP_TABLE_SIZE {
	// 	angle := float64(i) / LOOKUP_TABLE_SIZE * 2 * math.Pi // Map index to 0-2*Pi
	// 	sineTable[i] = math.Sin(angle)

	// 	angle = MAX_ANGLE * float64(i) / float64(LOOKUP_TABLE_SIZE-1)
	// 	cosTable[i] = math.Cos(angle)
	// }
}

// ------------------------------------------------------------------------------------------------
// Sin retrieves the sine value from the lookup table for a given angle.
// This is a basic lookup without interpolation.
func (l *LookupTables) Sin(angle float64) float64 {
	return getSinFromLookup(angle, l.sineTable)
}

// ------------------------------------------------------------------------------------------------
// SinI retrieves the sine value from the lookup table for a given angle.
// This function uses interpolation to provide a more accurate result.
func (l *LookupTables) SinI(angle float64) float64 {
	return getSinFromLookupInterpolated(angle, l.sineTable)
}

// ------------------------------------------------------------------------------------------------
func (l *LookupTables) Cos(angle float64) float64 {
	return getCosFromLookup(angle, l.cosineTable)
}

// ------------------------------------------------------------------------------------------------
// createSinLookupTable creates a sine lookup table for angles from 0 to 2*PI radians.
// numEntries determines the resolution of the table.
func createSinLookupTable(numEntries int) []float64 {
	// Use a slice instead of an array for more flexibility,
	// especially if the size is determined at runtime.
	sinLookup := make([]float64, numEntries)

	// Calculate the step size for the angle
	// Each entry represents an angle increment.
	angleStep := MAX_ANGLE / float64(numEntries)

	// Populate the lookup table
	for i := range numEntries {
		// Calculate the actual angle for the current index
		angle := float64(i) * angleStep
		sinLookup[i] = math.Sin(angle)
	}

	return sinLookup
}

// ------------------------------------------------------------------------------------------------
// getSinFromLookup retrieves the sine value from the lookup table for a given angle.
// This is a basic lookup without interpolation.
func getSinFromLookup(angle float64, lookupTable []float64) float64 {
	numEntries := len(lookupTable)
	// Normalize the angle to be within [0, 2*PI)
	normalizedAngle := math.Mod(angle, MAX_ANGLE)
	if normalizedAngle < 0 {
		normalizedAngle += MAX_ANGLE
	}

	// Map the angle to an index in the lookup table
	// This scales the angle to the table's index range.
	index := int(normalizedAngle / MAX_ANGLE * float64(numEntries))

	// Ensure the index is within bounds (should be, due to normalization and scaling)
	if index >= numEntries {
		index = numEntries - 1 // Should ideally not happen if angle is precisely 2*PI etc.
	}

	return lookupTable[index]
}

// ------------------------------------------------------------------------------------------------
// createCosLookupTable generates a lookup table for cosine values.
// It takes the number of entries as input and returns a slice of float64.
// The table covers a full 2*Pi cycle, with each entry representing
// an increment in angle.
func createCosLookupTable(numEntries int) []float64 {
	// Initialize a slice to store the cosine values.
	// Using a slice provides flexibility for dynamic sizing.
	cosLookup := make([]float64, numEntries)

	// Calculate the step size for the angle.
	// This determines the angular increment between each entry in the table.
	angleStep := MAX_ANGLE / float64(numEntries)

	// Populate the lookup table with cosine values.
	for i := range numEntries {
		// Calculate the actual angle for the current index.
		// The angle starts at 0 and goes up to (2*Pi - angleStep).
		angle := float64(i) * angleStep
		// Compute the cosine of the calculated angle and store it.
		cosLookup[i] = math.Cos(angle)
	}

	return cosLookup
}

// ------------------------------------------------------------------------------------------------
// getSinFromLookupInterpolated retrieves the sine value from the lookup table for a given angle,
// using linear interpolation for better accuracy.
func getSinFromLookupInterpolated(angle float64, lookupTable []float64) float64 {
	numEntries := len(lookupTable)
	angleRange := MAX_ANGLE // The total range covered by the table

	// Normalize the angle to be within [0, 2*PI)
	normalizedAngle := math.Mod(angle, angleRange)
	if normalizedAngle < 0 {
		normalizedAngle += angleRange
	}

	// Calculate the "floating point" index
	// This gives us the precise position within the table, including fractions.
	floatIndex := normalizedAngle / angleRange * float64(numEntries)

	// Get the indices of the two surrounding table entries
	idx1 := int(math.Floor(floatIndex))
	idx2 := int(math.Ceil(floatIndex))

	// Handle edge cases for the last entry
	if idx2 >= numEntries {
		idx2 = 0 // Wrap around for angles near 2*PI
	}

	// Get the values from the lookup table
	val1 := lookupTable[idx1]
	val2 := lookupTable[idx2]

	// Calculate the fractional part (how far between idx1 and idx2 the angle is)
	fraction := floatIndex - float64(idx1)

	// Perform linear interpolation
	return val1 + fraction*(val2-val1)
}

// ------------------------------------------------------------------------------------------------
// getCosFromLookup retrieves the sine value from the lookup table for a given angle.
// This is a basic lookup without interpolation.
func getCosFromLookup(angle float64, lookupTable []float64) float64 {
	numEntries := len(lookupTable)
	// Normalize the angle to be within [0, 2*PI)
	normalizedAngle := math.Mod(angle, MAX_ANGLE)
	if normalizedAngle < 0 {
		normalizedAngle += MAX_ANGLE
	}

	// Map the angle to an index in the lookup table
	// This scales the angle to the table's index range.
	index := int(normalizedAngle / MAX_ANGLE * float64(numEntries))

	// Ensure the index is within bounds (should be, due to normalization and scaling)
	if index >= numEntries {
		index = numEntries - 1 // Should ideally not happen if angle is precisely 2*PI etc.
	}

	return lookupTable[index]
}
