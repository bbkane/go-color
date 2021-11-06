// +build !windows

package color

import "errors"

// Enable enables color printing.
// On Windows it calls SetConsoleMode to enable ANSII colors
// (I think this only works in Win10...).
// Returns an error if called more than once.
func Enable() error {
	if enabled {
		return errors.New("Enable() called more than once")
	}
	setColors()
	enabled = true
	return nil
}
