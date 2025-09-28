package canvas

import (
	"testing"

	"github.com/ewaldhorn/gogi/colour"
)

func TestDrawRectangle(t *testing.T) {
	// 1. Create a new canvas for testing.
	canvas := NewCanvas(10, 10)

	// 2. Define a color for drawing.
	red := colour.NewColour(255, 0, 0, 255)

	// 3. Define the rectangle's properties.
	x, y := 2, 3
	width, height := 4, 5

	// 4. Draw the rectangle.
	canvas.DrawRectangle(x, y, width, height, red)

	// 5. Verify the result.
	// Check that all pixels within the rectangle are red.
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			pixel := canvas.GetPixel(i, j)
			if pixel != red {
				t.Errorf("Expected pixel at (%d, %d) to be red, but got %v", i, j, pixel)
			}
		}
	}

	// Check some pixels outside the rectangle to ensure they are still the default black.
	// The canvas is initialized to all zeros (black and transparent).
	pixelOutside1 := canvas.GetPixel(0, 0)
	if pixelOutside1.R != 0 || pixelOutside1.G != 0 || pixelOutside1.B != 0 || pixelOutside1.A != 0 {
		t.Errorf("Expected pixel at (0, 0) to be black/transparent, but got %v", pixelOutside1)
	}

	pixelOutside2 := canvas.GetPixel(9, 9)
	if pixelOutside2.R != 0 || pixelOutside2.G != 0 || pixelOutside2.B != 0 || pixelOutside2.A != 0 {
		t.Errorf("Expected pixel at (9, 9) to be black/transparent, but got %v", pixelOutside2)
	}
}
