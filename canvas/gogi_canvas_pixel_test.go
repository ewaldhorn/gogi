package canvas

import (
	"testing"

	"github.com/ewaldhorn/gogi/colour"
)

// ------------------------------------------------------------------------------------------------
func TestGetAndColourPutPixel(t *testing.T) {
	canvas := NewCanvas(10, 10)
	red := colour.NewColour(255, 0, 0, 255)

	// Test putting a pixel with full alpha
	canvas.ColourPutPixel(2, 3, red)
	pixel := canvas.GetPixel(2, 3)

	if pixel != red {
		t.Errorf("Expected pixel at (2, 3) to be red, but got %v", pixel)
	}

	// Test getting a pixel outside the bounds
	pixelOutside := canvas.GetPixel(11, 11)
	if pixelOutside.R != 0 || pixelOutside.G != 0 || pixelOutside.B != 0 || pixelOutside.A != 0 {
		t.Errorf("Expected pixel outside bounds to be empty, but got %v", pixelOutside)
	}
}

// ------------------------------------------------------------------------------------------------
func TestPutPixel(t *testing.T) {
	canvas := NewCanvas(10, 10)
	blue := colour.NewColour(0, 0, 255, 255)

	canvas.SetColour(blue)
	canvas.PutPixel(4, 5)
	pixel := canvas.GetPixel(4, 5)

	if pixel != blue {
		t.Errorf("Expected pixel at (4, 5) to be blue, but got %v", pixel)
	}
}

// ------------------------------------------------------------------------------------------------
func TestColourPutPixelAlphaBlending(t *testing.T) {
	canvas := NewCanvas(10, 10)

	// Fill background with white
	white := colour.NewColour(255, 255, 255, 255)
	canvas.DrawRectangle(0, 0, 10, 10, white)

	// Foreground color with 50% alpha
	red50 := colour.NewColour(255, 0, 0, 127)

	canvas.ColourPutPixel(5, 5, red50)
	pixel := canvas.GetPixel(5, 5)

	// Expected blended color
	// R = (255 * 127 + 255 * (255-127)) / 255 = 255
	// G = (0 * 127 + 255 * (255-127)) / 255 = 128
	// B = (0 * 127 + 255 * (255-127)) / 255 = 128
	// A = 127
	expected := colour.NewColour(255, 128, 128, 127)


	if pixel.R != expected.R || pixel.G != expected.G || pixel.B != expected.B {
		t.Errorf("Expected blended pixel at (5, 5) to be %v, but got %v", expected, pixel)
	}
}

// ------------------------------------------------------------------------------------------------
func TestBlendComponent(t *testing.T) {
	// 50% blend of white (255) and black (0) should be grey (127)
	grey := blendComponent(255, 0, 127)
	if grey != 127 {
		t.Errorf("Expected 127, got %d", grey)
	}

	// 25% blend of blue (255) and red (0)
	// (255 * 64 + 0 * (255-64)) / 255 = 64
	val := blendComponent(255, 0, 64)
	if val != 64 {
		t.Errorf("Expected 64, got %d", val)
	}
}
