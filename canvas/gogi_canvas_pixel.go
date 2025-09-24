package canvas

import "github.com/ewaldhorn/gogi/colour"

// ------------------------------------------------------------------------------------------------
// ColourPutPixel draws a single pixel at coordinates x,y using the specified
// colour. Does nothing if coordinates fall outside the canvas dimensions.
// Now supports alpha blending.
func (m *GogiCanvas) ColourPutPixel(x, y int, p colour.Colour) {
	const bytesPerPixel = 4
	offset := (x * bytesPerPixel) + (y * bytesPerPixel * m.width)

	// don't bother if we are outside our area
	if offset < 0 || offset+bytesPerPixel >= m.bufferSize {
		return
	}

	if p.A == 0 {
		// nothing to do, it's transparent
		return
	}

	if p.A == 255 {
		// full alpha, use this colour only
		m.pixelBuffer[offset] = p.R
		m.pixelBuffer[offset+1] = p.G
		m.pixelBuffer[offset+2] = p.B
		m.pixelBuffer[offset+3] = p.A
		return
	}

	// not full alpha, time to blend
	blendedColour := blendColors(p, m.GetPixel(x, y))
	m.pixelBuffer[offset] = blendedColour.R
	m.pixelBuffer[offset+1] = blendedColour.G
	m.pixelBuffer[offset+2] = blendedColour.B
	m.pixelBuffer[offset+3] = blendedColour.A
}

// ------------------------------------------------------------------------------------------------
func blendColors(fg, bg colour.Colour) colour.Colour {
	return colour.Colour{
		R: blendComponent(fg.R, bg.R, fg.A),
		G: blendComponent(fg.G, bg.G, fg.A),
		B: blendComponent(fg.B, bg.B, fg.A),
		A: fg.A,
	}
}

// ------------------------------------------------------------------------------------------------
// blendComponent blends a single color component.
func blendComponent(fg, bg, alpha uint8) uint8 {
	return uint8((uint16(fg)*uint16(alpha) + uint16(bg)*uint16(255-alpha)) / 255)
}

// ------------------------------------------------------------------------------------------------
// GetPixel returns the colour of the pixel at a given location.
func (m *GogiCanvas) GetPixel(x, y int) colour.Colour {
	offset := (x * 4) + (y * 4 * m.width)

	// don't bother if we are outside our area
	if offset < 0 || offset >= m.bufferSize {
		return colour.Colour{}
	}

	colour := colour.Colour{}

	colour.R = m.pixelBuffer[offset]
	colour.G = m.pixelBuffer[offset+1]
	colour.B = m.pixelBuffer[offset+2]
	colour.A = m.pixelBuffer[offset+3]

	return colour
}

// ------------------------------------------------------------------------------------------------
// PutPixel draws a single pixel at coordinates x,y using the active colour.
// Active colour can be set using SetColour(). Does nothing if coordinates
// fall outside the canvas dimensions.
func (m *GogiCanvas) PutPixel(x, y int) {
	m.ColourPutPixel(x, y, m.activeColour)
}
