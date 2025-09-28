package canvas

import (
	"testing"

	"github.com/ewaldhorn/gogi/colour"
)

func TestDrawTriangle(t *testing.T) {
	// 1. Create a new canvas for testing.
	canvas := NewCanvas(20, 20)

	// 2. Define a color for drawing and set it on the canvas.
	white := colour.NewColourWhite()
	canvas.SetColour(white)

	// 3. Define the triangle's points.
	p1 := Point{X: 5, Y: 5}
	p2 := Point{X: 15, Y: 10}
	p3 := Point{X: 10, Y: 15}

	// 4. Draw the triangle.
	canvas.DrawTriangle(p1, p2, p3)

	// 5. Verify the result.
	// We'll check the pixels at the vertices of the triangle.
	// A more comprehensive test could check more points along the lines.

	// Check the color of the pixel at the first vertex.
	pixel1 := canvas.GetPixel(p1.X, p1.Y)
	if pixel1 != white {
		t.Errorf("Expected pixel at p1 (%d, %d) to be white, but got %v", p1.X, p1.Y, pixel1)
	}

	// Check the color of the pixel at the second vertex.
	pixel2 := canvas.GetPixel(p2.X, p2.Y)
	if pixel2 != white {
		t.Errorf("Expected pixel at p2 (%d, %d) to be white, but got %v", p2.X, p2.Y, pixel2)
	}

	// Check the color of the pixel at the third vertex.
	pixel3 := canvas.GetPixel(p3.X, p3.Y)
	if pixel3 != white {
		t.Errorf("Expected pixel at p3 (%d, %d) to be white, but got %v", p3.X, p3.Y, pixel3)
	}

	// Check a point that should NOT be part of the triangle.
	// The canvas is initialized to all zeros (black and transparent).
	pixelOutside := canvas.GetPixel(0, 0)
	if pixelOutside.R != 0 || pixelOutside.G != 0 || pixelOutside.B != 0 || pixelOutside.A != 0 {
		t.Errorf("Expected pixel at (0, 0) to be black/transparent, but got %v", pixelOutside)
	}
}
