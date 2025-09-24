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
	sqrtTable  []float64
	radToIndex float64
	palette    [256]colour.Colour
)

// ----------------------------------------------------------------------------
type Scenario struct {
	t                         float64
	renderWidth, renderHeight int
	renderBuffer              buffers.PixelBuffer
}

// ----------------------------------------------------------------------------
func fastSin(val float64) float64 {
	idx := int(val*radToIndex) & (SINE_TABLE_SIZE - 1)
	return sineTable[idx]
}

// ----------------------------------------------------------------------------
func fastSqrt(val float64) float64 {
	if val < 0 || int(val) >= len(sqrtTable) {
		return math.Sqrt(val)
	}
	return sqrtTable[int(val)]
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
	// Initialize sine lookup table
	radToIndex = SINE_TABLE_SIZE / (2 * math.Pi)
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

	// Initialize square root lookup table
	maxDistSq := scenario.renderWidth/2*scenario.renderWidth/2 + scenario.renderHeight/2*scenario.renderHeight/2
	sqrtTable = make([]float64, maxDistSq+1)
	for i := range sqrtTable {
		sqrtTable[i] = math.Sqrt(float64(i))
	}

	// Initialize color palette
	for i := range 256 {
		colorInput := float64(i)
		r := clampByte(fastSin(colorInput*0.02+0.0)*127.0 + 128.0)
		g := clampByte(fastSin(colorInput*0.02+2.0)*64.0 + 190.0)
		b := clampByte(fastSin(colorInput*0.02+4.0)*127.0 + 128.0)
		palette[i] = colour.NewColour(r, g, b, 255)
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

	renderWidth := scenario.renderWidth
	renderHeight := scenario.renderHeight
	renderWidthHalf := renderWidth / 2
	renderHeightHalf := renderHeight / 2

	// first calculate the smaller buffer
	for y := range renderHeight {
		ny := float64(y) * 0.045
		for x := range renderWidth {
			nx := float64(x) * 0.045

			val1 := fastSin(nx + scenario.t)
			val2 := fastSin(ny + scenario.t*0.5)
			val3 := fastSin((nx+ny)*0.7 + scenario.t*0.8)

			dx := float64(x - renderWidthHalf)
			dy := float64(y - renderHeightHalf)
			dist_sq := dx*dx + dy*dy
			dist := fastSqrt(dist_sq) * 0.01
			val4 := fastSin(dist + scenario.t*0.3)

			plasmaValue := (val1 + val2 + val3 + val4) / 4.0

			// Map plasma value to palette index
			// plasmaValue is in [-1, 1]. Map it to [0, 255].
			idx := int((plasmaValue + 1.0) * 127.5)
			if idx > 255 {
				idx = 255
			}
			if idx < 0 {
				idx = 0
			}

			scenario.renderBuffer.ColourPutPixel(x, y, palette[idx])
		}
	}

	// now actually render it by upscaling
	for y := range CANVAS_HEIGHT {
		for x := range CANVAS_WIDTH {
			// Map x, y from CANVAS_WIDTH/HEIGHT to RENDER_WIDTH/HEIGHT
			srcX := x * renderWidth / CANVAS_WIDTH
			srcY := y * renderHeight / CANVAS_HEIGHT
			c := scenario.renderBuffer.GetPixel(srcX, srcY)
			gameCanvas.ColourPutPixel(x, y, c)
		}
	}
}
