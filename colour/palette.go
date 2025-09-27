package colour

import "github.com/ewaldhorn/gogi/utils"

// ------------------------------------------------------------------------------------------------
func GetFirePalette() []Colour {
	fire := make([]Colour, 256)

	for i := range 256 {
		l := float64(i) / 255.0

		hDeg := utils.MapValue(float64(i), 0, 255, 0, 85)
		h := hDeg / 360.0

		s := 1.0

		r, g, b := HSLToRGB(h, s, l)

		fire[i] = NewColour(r, g, b, 255)
	}

	return fire
}

// ------------------------------------------------------------------------------------------------
func GetAllRedPalette() []Colour {
	allRed := make([]Colour, 256)

	for i := range 256 {
		l := float64(i) / 255.0

		hDeg := utils.MapValue(float64(i), 0, 255, 0, 119)
		h := hDeg / 360.0

		s := 1.0

		r, _, _ := HSLToRGB(h, s, l)

		allRed[i] = NewColour(r, 0, 0, 255)
	}

	return allRed
}
