package canvas

// ------------------------------------------------------------------------------------------------
// DrawLine draws a line between (x0, y0) and (x1, y1) using Bresenham's algorithm.
// It uses only integer arithmetic.
func (m *GogiCanvas) DrawLine(x0, y0, x1, y1 int) {
	// Determine if the line is steep (more vertical than horizontal)
	// This helps in swapping x and y coordinates to always iterate along the major axis.
	steep := abs(y1-y0) > abs(x1-x0)

	// If the line is steep, swap x and y coordinates for calculation.
	// This ensures we always iterate along the x-axis (or what becomes the x-axis).
	if steep {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}

	// Ensure that x0 is less than or equal to x1.
	// If not, swap the start and end points to draw from left to right.
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	// Calculate the change in x and y.
	dx := x1 - x0
	dy := abs(y1 - y0)

	// Initialize the error term.
	// This term tracks how far off the line we are from the ideal path.
	error := dx / 2

	// Determine the direction of y-increment.
	// If y1 is greater than y0, we increment y; otherwise, we decrement y.
	yStep := 1
	if y0 > y1 {
		yStep = -1
	}

	// y-coordinate for the current pixel, initialized to y0.
	y := y0

	// Iterate along the major axis (which is now x due to potential swapping).
	for x := x0; x <= x1; x++ {
		// If the line was steep, swap x and y back before setting the pixel.
		// This translates the calculated (x, y) back to the original coordinate system.
		if steep {
			m.PutPixel(y, x)
		} else {
			m.PutPixel(x, y)
		}

		// Update the error term.
		// We add the change in y (dy) to the error.
		error -= dy

		// If the error becomes negative, it means we've crossed the midpoint
		// between two potential pixels on the minor axis, so we need to
		// increment/decrement the minor axis coordinate (y) and reset the error.
		if error < 0 {
			y += yStep
			error += dx // Add dx back to reset the error for the next step.
		}
	}
}
