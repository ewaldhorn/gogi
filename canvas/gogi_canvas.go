package canvas

import (
	"unsafe"

	"github.com/ewaldhorn/gogi/colour"
)

// ------------------------------------------------------------------------------------------------
type GogiCanvas struct {
	width, height int
	bufferSize    int
	pixelBuffer   []uint8

	activeColour colour.Colour
	savedColour  colour.Colour
}

// ------------------------------------------------------------------------------------------------
func NewCanvas(width, height int) *GogiCanvas {
	temp := GogiCanvas{
		width:  width,
		height: height,
	}

	temp.bufferSize = width * height * colour.BYTES_PER_PIXEL
	temp.pixelBuffer = make([]uint8, temp.bufferSize)

	return &temp
}

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) GetBuffer() []uint8 {
	bufferCopy := make([]uint8, len(m.pixelBuffer))
	copy(bufferCopy, m.pixelBuffer)
	return bufferCopy
}

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) ClearBuffer() {
	for i := range m.pixelBuffer {
		m.pixelBuffer[i] = 0
	}
}

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) Width() int {
	return m.width
}

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) Height() int {
	return m.height
}

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) GetBufferPointer() uintptr {
	// unsafe.Pointer(&pixelBuffer[0]) gets the address of the first element so that
	// we can pass it back to JavaScript
	return uintptr(unsafe.Pointer(&m.pixelBuffer[0]))
}

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) GetBufferLength() uint32 {
	return uint32(len(m.pixelBuffer))
}

// ------------------------------------------------------------------------------------------------
// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
