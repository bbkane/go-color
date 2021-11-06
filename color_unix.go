//go:build !windows
// +build !windows

package color

import "errors"

// Enable enables color printing.
// It doesn't currently seem to work on Windows
// and I don't really have access to a Windows machine
// to fix it.
func Enable() error {
	if enabled {
		return errors.New("Enable() called more than once")
	}
	setColors()
	enabled = true
	return nil
}
