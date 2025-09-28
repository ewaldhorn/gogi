package canvas

import "math"

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) DrawTriangle(p1, p2, p3 Point) {
	m.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
	m.DrawLine(p2.X, p2.Y, p3.X, p3.Y)
	m.DrawLine(p3.X, p3.Y, p1.X, p1.Y)
}

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) DrawTrianglePointedTo(start, destination Point) {
	// 1. Calculate the direction vector
	dx := destination.X - start.X
	dy := destination.Y - start.Y

	// 2. Calculate the angle of movement (in radians)
	angle := math.Atan2(float64(dy), float64(dx))

	triangleLength := 20.0
	triangleWidth := 10.0

	// Relative coordinates assuming tip is at (0,0) and facing positive X
	// Base point 1 (behind the tip, up relative to the axis)
	base1RelX := -triangleLength
	base1RelY := triangleWidth / 2

	// Base point 2 (behind the tip, down relative to the axis)
	base2RelX := -triangleLength
	base2RelY := -triangleWidth / 2

	// 4. Rotate the relative points to align with the movement angle
	// Rotation matrix: x' = x*cos(a) - y*sin(a), y' = x*sin(a) + y*cos(a)

	cosAngle := math.Cos(angle)
	sinAngle := math.Sin(angle)

	// Rotate Base Point 1
	rotatedBase1X := base1RelX*cosAngle - base1RelY*sinAngle
	rotatedBase1Y := base1RelX*sinAngle + base1RelY*cosAngle

	// Rotate Base Point 2
	rotatedBase2X := base2RelX*cosAngle - base2RelY*sinAngle
	rotatedBase2Y := base2RelX*sinAngle + base2RelY*cosAngle

	// 5. Translate the rotated points to the starting position (which is the tip)
	tip := start
	base1 := Point{X: start.X + int(rotatedBase1X), Y: start.Y + int(rotatedBase1Y)}
	base2 := Point{X: start.X + int(rotatedBase2X), Y: start.Y + int(rotatedBase2Y)}

	m.DrawTriangle(tip, base1, base2)
}
