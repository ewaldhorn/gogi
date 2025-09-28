//go:build !tinygo

package buffers

import (
	"testing"
)

// ------------------------------------------------------------------------------------------------
func TestNewPixelBufferPanic(t *testing.T) {
	// Test for panic with wrong buffer size
	t.Run("Panic on buffer size mismatch", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("The code did not panic")
			}
			expectedPanic := "buffer size mismatch"
			if r != expectedPanic {
				t.Errorf("Expected panic message '%s', but got '%v'", expectedPanic, r)
			}
		}()

		width, height := 10, 10
		buffer := make([]uint8, 1) // Wrong size
		NewPixelBuffer(width, height, buffer)
	})
}
