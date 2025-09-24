package buffers

import "github.com/ewaldhorn/gogi/colour"

// ------------------------------------------------------------------------------------------------
type PixelBuffer struct {
	width, height int
	bufferSize    int
	pixels        []uint8
	bytesPerPixel int
}

// ------------------------------------------------------------------------------------------------
func NewPixelBuffer(width, height int, buffer []uint8) *PixelBuffer {
	newBuffer := make([]uint8, len(buffer))
	copy(newBuffer, buffer)

	return &PixelBuffer{
		width: width, height: height, pixels: newBuffer, bytesPerPixel: 4, bufferSize: len(newBuffer),
	}
}

// ------------------------------------------------------------------------------------------------
func (p *PixelBuffer) GetPixel(x, y int) colour.Colour {
	if x < 0 || x >= p.width || y < 0 || y >= p.height {
		return colour.Colour{}
	}

	offset := (y*p.width + x) * p.bytesPerPixel

	col := colour.Colour{}

	col.R = p.pixels[offset]
	col.G = p.pixels[offset+1]
	col.B = p.pixels[offset+2]
	col.A = p.pixels[offset+3]

	return col
}

// ------------------------------------------------------------------------------------------------
func (m *PixelBuffer) ColourPutPixel(x, y int, p colour.Colour) {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return
	}

	// if the foreground is transparent, there's nothing to do
	if p.A == 0 {
		return
	}

	offset := (y*m.width + x) * m.bytesPerPixel

	// if the foreground is fully opaque, we can just overwrite the background
	if p.A == 255 {
		m.pixels[offset] = p.R
		m.pixels[offset+1] = p.G
		m.pixels[offset+2] = p.B
		m.pixels[offset+3] = p.A
		return
	}

	bg := colour.Colour{
		R: m.pixels[offset],
		G: m.pixels[offset+1],
		B: m.pixels[offset+2],
		A: m.pixels[offset+3],
	}

	blendedColour := blendColors(p, bg)
	m.pixels[offset] = blendedColour.R
	m.pixels[offset+1] = blendedColour.G
	m.pixels[offset+2] = blendedColour.B
	m.pixels[offset+3] = blendedColour.A
}

// ------------------------------------------------------------------------------------------------
func blendColors(fg, bg colour.Colour) colour.Colour {
	fgA := uint32(fg.A)
	bgA := uint32(bg.A)

	outA := fgA + (bgA*(255-fgA))/255
	if outA == 0 {
		return colour.Colour{}
	}

	outR := (uint32(fg.R)*fgA + (uint32(bg.R)*bgA*(255-fgA))/255) / outA
	outG := (uint32(fg.G)*fgA + (uint32(fg.G)*bgA*(255-fgA))/255) / outA
	outB := (uint32(fg.B)*fgA + (uint32(fg.B)*bgA*(255-fgA))/255) / outA

	return colour.Colour{
		R: uint8(outR),
		G: uint8(outG),
		B: uint8(outB),
		A: uint8(outA),
	}
}
