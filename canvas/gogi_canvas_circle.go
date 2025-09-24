package canvas

import "github.com/ewaldhorn/gogi/colour"

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) DrawCircle(cx, cy int, radius int, col colour.Colour) {
	x := 0
	y := radius
	p := 3 - 2*radius

	drawOctants := func(x, y int) {
		m.ColourPutPixel(cx+x, cy+y, col)
		m.ColourPutPixel(cx-x, cy+y, col)
		m.ColourPutPixel(cx+x, cy-y, col)
		m.ColourPutPixel(cx-x, cy-y, col)
		m.ColourPutPixel(cx+y, cy+x, col)
		m.ColourPutPixel(cx-y, cy+x, col)
		m.ColourPutPixel(cx+y, cy-x, col)
		m.ColourPutPixel(cx-y, cy-x, col)
	}

	for x <= y {
		drawOctants(x, y)
		x++

		if p < 0 {
			p = p + 4*x + 6
		} else {
			y--
			p = p + 4*(x-y) + 10
		}
	}
}

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) DrawFilledCircle(cx, cy int, radius int, col colour.Colour) {
	x := 0
	y := radius
	p := 3 - 2*radius

	// Helper function to draw a horizontal line
	// Assumes startX <= endX
	drawHorizontalLine := func(x1, x2, y int) {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for i := x1; i <= x2; i++ {
			m.ColourPutPixel(i, y, col)
		}
	}

	// Function to handle drawing based on the octants.
	// Instead of just drawing the pixel, we fill the horizontal space.
	drawAndFillOctants := func(x, y int) {
		// Octants 1 & 8 (Top Right & Bottom Right)
		drawHorizontalLine(cx-x, cx+x, cy+y)
		drawHorizontalLine(cx-x, cx+x, cy-y)

		// Octants 2 & 7 (Top Right & Bottom Right)
		// Note: When x and y are close, these might overlap or be the same as above.
		// We ensure we only draw distinct lines.
		if x != y { // Only draw these if they are distinct from the previous lines
			drawHorizontalLine(cx-y, cx+y, cy+x)
			drawHorizontalLine(cx-y, cx+y, cy-x)
		}
	}

	for x <= y {
		// Draw the current scanlines based on the calculated x, y points
		drawAndFillOctants(x, y)

		x++

		if p < 0 {
			p = p + 4*x + 6
		} else {
			y--
			p = p + 4*(x-y) + 10
		}
	}
}
