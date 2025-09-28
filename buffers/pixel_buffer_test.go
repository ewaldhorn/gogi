package buffers

import (
	"testing"

	"github.com/ewaldhorn/gogi/colour"
)

// ------------------------------------------------------------------------------------------------
func TestNewPixelBuffer(t *testing.T) {
	// Test successful creation
	t.Run("Successful creation", func(t *testing.T) {
		width, height := 10, 10
		buffer := make([]uint8, width*height*RGBABytesPerPixel)
		pb := NewPixelBuffer(width, height, buffer)

		if pb.width != width || pb.height != height {
			t.Errorf("Expected dimensions %dx%d, got %dx%d", width, height, pb.width, pb.height)
		}
	})
}

// ------------------------------------------------------------------------------------------------
func TestGetAndColourPutPixel(t *testing.T) {
	width, height := 10, 10
	buffer := make([]uint8, width*height*RGBABytesPerPixel)
	pb := NewPixelBuffer(width, height, buffer)
	red := colour.NewColour(255, 0, 0, 255)

	// Test putting a pixel with full alpha
	pb.ColourPutPixel(2, 3, red)
	pixel := pb.GetPixel(2, 3)

	if pixel != red {
		t.Errorf("Expected pixel at (2, 3) to be red, but got %v", pixel)
	}

	// Test getting a pixel outside the bounds
	pixelOutside := pb.GetPixel(11, 11)
	if pixelOutside.R != 0 || pixelOutside.G != 0 || pixelOutside.B != 0 || pixelOutside.A != 0 {
		t.Errorf("Expected pixel outside bounds to be empty, but got %v", pixelOutside)
	}
}

// ------------------------------------------------------------------------------------------------
func TestBlendColors(t *testing.T) {
	// Background: White, opaque
	bg := colour.NewColour(255, 255, 255, 255)
	// Foreground: Red, 50% transparent
	fg := colour.NewColour(255, 0, 0, 128)

	blended := blendColors(fg, bg)

	// Expected values from the formula
	// outA = 128 + (255 * (255-128))/255 = 128 + 127 = 255
	// outR = (255*128 + (255*255*(255-128))/255) / 255 = (32640 + 32385) / 255 = 255
	// outG = (0*128 + (255*255*(255-128))/255) / 255 = 127
	// outB = (0*128 + (255*255*(255-128))/255) / 255 = 127

	if blended.R != 255 || blended.G != 127 || blended.B != 127 || blended.A != 255 {
		t.Errorf("Expected (255, 127, 127, 255), got %v", blended)
	}
}
