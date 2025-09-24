package colour

import (
	"math/rand/v2"
)

const MAX_COLOUR_VALUE uint8 = 255
const BYTES_PER_PIXEL = 4

// ------------------------------------------------------------------------------------------------
type Colour struct {
	R, G, B, A uint8
}

// ------------------------------------------------------------------------------------------------
// an empty colour is used for transparent effects
func (c *Colour) IsEmpty() bool {
	return c.A == 0 && c.B == 0 && c.G == 0 && c.R == 0
}

// ------------------------------------------------------------------------------------------------
func (c *Colour) FadeALittle() {
	c.R = uint8(float32(c.R) * 0.98)
	c.G = uint8(float32(c.G) * 0.93)
	c.B = uint8(float32(c.B) * 0.93)
}

// ------------------------------------------------------------------------------------------------
func (c *Colour) DecreaseAlpha() {
	if c.A > 0 {
		c.A -= 1
	}
}

// ------------------------------------------------------------------------------------------------
func NewColour(r, g, b, a uint8) Colour {
	return Colour{R: r, G: g, B: b, A: a}
}

// ------------------------------------------------------------------------------------------------
func NewColourWhite() Colour {
	return Colour{R: MAX_COLOUR_VALUE, G: MAX_COLOUR_VALUE, B: MAX_COLOUR_VALUE, A: MAX_COLOUR_VALUE}
}

// ------------------------------------------------------------------------------------------------
func NewColourBlack() Colour {
	return Colour{A: MAX_COLOUR_VALUE}
}

// ------------------------------------------------------------------------------------------------
// empty colour signals to the renderer not to draw anything and is used to create
// transparent "gaps" in images. This is a legacy feature to support some very
// old file formats.
func NewColourEmpty() Colour {
	return Colour{}
}

// ------------------------------------------------------------------------------------------------
func NewRandomColour() Colour {
	return Colour{
		R: uint8(rand.Float32() * float32(MAX_COLOUR_VALUE)),
		G: uint8(rand.Float32() * float32(MAX_COLOUR_VALUE)),
		B: uint8(rand.Float32() * float32(MAX_COLOUR_VALUE)),
		A: MAX_COLOUR_VALUE,
	}
}

// ------------------------------------------------------------------------------------------------
// Built using information from https://en.wikipedia.org/wiki/Grayscale
// and https://stackoverflow.com/questions/42516203/converting-rgba-image-to-grayscale-golang
func (c *Colour) ConvertToGrayscale() {
	shadeOfGray := uint8((0.299*(float64(c.R)) + 0.587*(float64(c.G)) + 0.144*(float64(c.B))) / 256)

	c.R = shadeOfGray
	c.G = shadeOfGray
	c.B = shadeOfGray
}

// ------------------------------------------------------------------------------------------------
// HSLToRGB converts HSL values to RGB.
// h, s, l are expected to be in the range [0, 1].
// Returns r, g, b in the range [0, 255].
func HSLToRGB(h, s, l float64) (r, g, b uint8) {
	var red, green, blue float64

	if s == 0 {
		red = l
		green = l
		blue = l
	} else {
		hue2rgb := func(p, q, t float64) float64 {
			if t < 0 {
				t += 1
			}
			if t > 1 {
				t -= 1
			}
			if t < 1.0/6.0 {
				return p + (q-p)*6*t
			}
			if t < 1.0/2.0 {
				return q
			}
			if t < 2.0/3.0 {
				return p + (q-p)*(2.0/3.0-t)*6
			}
			return p
		}

		q := 0.0
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		red = hue2rgb(p, q, h+1.0/3.0)
		green = hue2rgb(p, q, h)
		blue = hue2rgb(p, q, h-1.0/3.0)
	}

	return uint8(red * 255), uint8(green * 255), uint8(blue * 255)
}

// ------------------------------------------------------------------------------------------------
// MapValue maps a value from one range to another.
func MapValue(value, inMin, inMax, outMin, outMax float64) float64 {
	return (value-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
