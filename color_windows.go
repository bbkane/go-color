package color

import (
	"os"
	"syscall"
)

// colorizeFileWindows calls SetConsoleMode to enable ANSII colors to work.
// This should probably only be called for os.Stdout.Fd() and os.Stderr.Fd()
func colorizeFileWindows(fp *os.File) error {
	// Try to make ANSI work
	handle := syscall.Handle(fp)
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
	// If it fails, fallback to no colors
	if _, _, err := setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004); err != nil && err.Error() != "The operation completed successfully." {
		return err
	}
}

// Enable enables color printing.
// On Windows it calls SetConsoleMode to enable ANSII colors
// (I think this only works in Win10...).
// Returns an error if called more than once.
func Enable() error {
	// See https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences?redirectedfrom=MSDN#text-formatting
	// and
	// https://docs.microsoft.com/en-us/windows/console/setconsolemode
	if initialized {
		return errors.New("Init() called more than once")
	}

	err := colorizeFileWindows(os.Stdout.Fd())
	if err != nil {
		return err
	}
	err := colorizeFileWindows(os.Stderr.Fd())
	if err != nil {
		return err
	}

	setColors()
	initialized = true
	return nil
}
