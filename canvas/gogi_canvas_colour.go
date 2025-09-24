package canvas

import "github.com/ewaldhorn/gogi/colour"

// ------------------------------------------------------------------------------------------------
// GetColour returns the active colour currently set
func (m *GogiCanvas) GetColour() colour.Colour {
	return m.activeColour
}

// ------------------------------------------------------------------------------------------------
// SetColour sets the active colour to be used when drawing
func (m *GogiCanvas) SetColour(p colour.Colour) {
	m.activeColour = p
}

// ------------------------------------------------------------------------------------------------
// Saves the current colour
func (m *GogiCanvas) SaveColour() {
	m.savedColour = m.activeColour
}

// ------------------------------------------------------------------------------------------------
// Switches and saves the colour
func (m *GogiCanvas) SwitchAndSaveColour(colour colour.Colour) {
	m.SaveColour()
	m.SetColour(colour)
}

// ------------------------------------------------------------------------------------------------
// Restores the saved colour
func (m *GogiCanvas) RestoreColour() {
	m.SetColour(m.savedColour)
}
