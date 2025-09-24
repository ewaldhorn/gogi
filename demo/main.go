// main.go
package main

import (
	"math"

	"github.com/ewaldhorn/gogi/buffers"
	"github.com/ewaldhorn/gogi/canvas"
	"github.com/ewaldhorn/gogi/colour"
)

// ----------------------------------------------------------------------------
const (
	// canvas properties
	CANVAS_WIDTH    = 800
	CANVAS_HEIGHT   = 600
	HALF_WIDTH      = CANVAS_WIDTH / 2
	HALF_HEIGHT     = CANVAS_HEIGHT / 2
	SINE_TABLE_SIZE = 4096 * 4 // Adjust for desired precision vs. memory
)

// ----------------------------------------------------------------------------
var (
	gameCanvas *canvas.GogiCanvas
	scenario   Scenario
	BLACK      colour.Colour
	sineTable  [SINE_TABLE_SIZE]float64
)

// ----------------------------------------------------------------------------
type Scenario struct {
	t                         float64
	renderWidth, renderHeight int
	renderBuffer              buffers.PixelBuffer
}

// ----------------------------------------------------------------------------
func fastSin(val float64) float64 {
	// Map val to an index in the table, handling periodicity
	// This mapping depends on the expected range of 'val'
	// For example, if val typically ranges from -X to X, you need to normalize it
	// and wrap around the table.
	// A simple example assuming val is scaled to 0 to 2*Pi somewhere:
	idx := int(val*(SINE_TABLE_SIZE/(2*math.Pi))) % SINE_TABLE_SIZE
	if idx < 0 {
		idx += SINE_TABLE_SIZE // Handle negative results from modulo
	}
	return sineTable[idx]
}

// ----------------------------------------------------------------------------
func main() {
	// The main function is empty as initGame and update are exported for WASM.
}

// ----------------------------------------------------------------------------
func clampByte(val float64) uint8 {
	if val < 0 {
		return 0
	}
	if val > 255 {
		return 255
	}
	return uint8(val)
}

// ----------------------------------------------------------------------------
//
//export initGame
func initGame() {
	for i := range SINE_TABLE_SIZE {
		angle := float64(i) / SINE_TABLE_SIZE * 2 * math.Pi // Map index to 0-2*Pi
		sineTable[i] = math.Sin(angle)
	}

	gameCanvas = canvas.NewCanvas(CANVAS_WIDTH, CANVAS_HEIGHT)
	gameCanvas.ClearBuffer()

	BLACK = colour.NewColourBlack()

	var bufSize = (CANVAS_HEIGHT / 2) * (CANVAS_WIDTH / 2) * 4

	scenario = Scenario{
		t:            0.0,
		renderWidth:  CANVAS_WIDTH / 2,
		renderHeight: CANVAS_HEIGHT / 2,
		renderBuffer: *buffers.NewPixelBuffer(CANVAS_WIDTH/2, CANVAS_HEIGHT/2, make([]uint8, bufSize)),
	}
}

// ----------------------------------------------------------------------------
// getBufferPointer returns the memory address of the start of the pixelBuffer.
// This allows JavaScript to locate the buffer within the WebAssembly memory.
//
//export getBufferPointer
func getBufferPointer() uintptr {
	return gameCanvas.GetBufferPointer()
}

// ----------------------------------------------------------------------------
// getBufferLength returns the total size of the pixelBuffer in bytes.
// JavaScript needs this to know how much memory to read.
//
//export getBufferLength
func getBufferLength() uint32 {
	return gameCanvas.GetBufferLength()
}

// ----------------------------------------------------------------------------
//
//export update
func update() {
	scenario.t += 0.05
	if scenario.t > 1000000 {
		scenario.t = 0.05
	}
	// first calculate the smaller buffer
	for y := range scenario.renderHeight {
		ny := float64(y) * 0.045
		for x := range scenario.renderWidth {
			// Normalize x and y coordinates to a smaller range for sine wave calculations.
			// This scaling factor (e.g., 0.015) affects the "zoom" or "density" of the plasma.
			nx := float64(x) * 0.045

			// --- Plasma Calculation ---
			// This is the core of the plasma effect. We combine multiple sine waves
			// with different frequencies, phases, and offsets.
			// Experiment with these values!

			// Wave 1: Based on X and time
			val1 := fastSin(nx + scenario.t)

			// Wave 2: Based on Y and time
			val2 := fastSin(ny + scenario.t*0.5) // Slower time progression for this wave

			// Wave 3: Based on X+Y and time
			val3 := fastSin((nx+ny)*0.7 + scenario.t*0.8) // Different frequency and speed

			// Wave 4: Based on distance from center and time
			// Calculate distance from center (approximate)
			dx := float64(x - HALF_WIDTH)
			dy := float64(y - HALF_HEIGHT)
			dist := math.Sqrt(dx*dx+dy*dy) * 0.01  // Scale distance
			val4 := fastSin(dist + scenario.t*0.3) // Wave based on radial distance

			// Combine the waves. The weights (e.g., 1.0, 0.8, 0.6, 0.4) can be adjusted.
			// The overall result will be in a range like -4.0 to 4.0 (if all weights are 1.0).
			plasmaValue := (val1 + val2 + val3 + val4) / 4.0 // Normalize to -1.0 to 1.0 range

			// --- Map Plasma Value to Color ---
			// We'll map the plasmaValue (which is between -1.0 and 1.0) to an RGB color.
			// A common technique is to use sine waves again for each color channel,
			// offsetting their phases to get a smooth color transition.

			// Scale plasmaValue to a 0-255 range for color mapping
			// This shifts it from [-1, 1] to [0, 2] and then scales it.
			colorInput := (plasmaValue + 1.0) * 128.0
			// Generate RGB components using sine waves with different phase offsets
			r := clampByte(fastSin(colorInput*0.02+0.0)*127.0 + 128.0) // Red channel
			g := clampByte(fastSin(colorInput*0.02+2.0)*64.0 + 190.0)  // Green channel (offset by 2 radians)
			b := clampByte(fastSin(colorInput*0.02+4.0)*127.0 + 128.0) // Blue channel (offset by 4 radians)

			// Alpha channel (fully opaque)
			a := uint8(255)
			scenario.renderBuffer.ColourPutPixel(x, y, colour.NewColour(r, g, b, a))
		}
	}

	// now actually render it
	for y := range CANVAS_HEIGHT {
		for x := range CANVAS_WIDTH {
			// Map x, y from CANVAS_WIDTH/HEIGHT to RENDER_WIDTH/HEIGHT
			srcX := int(float64(x) / CANVAS_WIDTH * float64(scenario.renderWidth))
			srcY := int(float64(y) / CANVAS_HEIGHT * float64(scenario.renderHeight))
			c := scenario.renderBuffer.GetPixel(srcX, srcY)
			gameCanvas.ColourPutPixel(x, y, c)
		}
	}
}
