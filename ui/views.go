package ui

import (
	"github.com/jroimartin/gocui"
)

// Labels and values use the default colour of the terminal (values in bold),
// so that they are readable whatever the colour scheme is. Dark blue was
// nearly impossible to read on a black background.
const (
	boldTextColour    = "\033[0;1m"
	normalTextColour  = "\033[0m"
	resetTextColour   = "\033[0m"
	volumeColour      = "\033[31;2m"
	volumeMutedColour = "\033[31;1m"
	progressColour    = "\033[33;2m"
)

// views sets up all of the views:
func (ui *UserInterface) views(g *gocui.Gui) error {
	ui.viewStatus(g)
	ui.viewVolume(g)
	ui.viewProgress(g)
	ui.viewLog(g)
	ui.viewKeys(g)
	return nil
}
