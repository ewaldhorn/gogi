package canvas

import "github.com/ewaldhorn/gogi/colour"

// ------------------------------------------------------------------------------------------------
func (m *GogiCanvas) DrawRectangle(x, y, width, height int, drawColour colour.Colour) {
	for py := y; py < y+height; py++ {
		for px := x; px < x+width; px++ {
			if px >= 0 && px < m.width && py >= 0 && py < m.height {
				index := (py*m.width + px) * colour.BYTES_PER_PIXEL

				m.pixelBuffer[index+0] = drawColour.R
				m.pixelBuffer[index+1] = drawColour.G
				m.pixelBuffer[index+2] = drawColour.B
				m.pixelBuffer[index+3] = 255
			}
		}
	}
}
