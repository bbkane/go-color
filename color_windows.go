package color

import (
	"errors"
	"os"
	"syscall"
)

// Enable enables color printing.
// On Windows it calls SetConsoleMode to enable ANSII colors
// (I think this only works in Win10...).
// Returns an error if called more than once.
func Enable() error {
	// See https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences?redirectedfrom=MSDN#text-formatting
	// and
	// https://docs.microsoft.com/en-us/windows/console/setconsolemode
	if enabled {
		return errors.New("Init() called more than once")
	}

	// Try to make ANSI work
	handle := syscall.Handle(os.Stdout.Fd())
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
	// If it fails, fallback to no colors
	if _, _, err := setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004); err != nil && err.Error() != "The operation completed successfully." {
		return err
	}

	// Try to make ANSI work
	handle := syscall.Handle(os.Stderr.Fd())
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
	// If it fails, fallback to no colors
	if _, _, err := setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004); err != nil && err.Error() != "The operation completed successfully." {
		return err
	}

	setColors()
	enabled = true
	return nil
}
