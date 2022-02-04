package color

import (
	"os"

	"golang.org/x/sys/windows"
)

// Enable enables color printing on Windows and is a no-op
// on non-Windows platforms (which enable color printing by default)
func Enable() error {
	// Code adapted from: https://github.com/jedib0t/go-pretty

	// See https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences?redirectedfrom=MSDN#text-formatting
	// and
	// https://docs.microsoft.com/en-us/windows/console/setconsolemode

	handle := windows.Handle(os.Stdout.Fd())

	var stdoutConsoleMode uint32
	var flags uint32 = windows.ENABLE_PROCESSED_OUTPUT |
		windows.ENABLE_WRAP_AT_EOL_OUTPUT |
		windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING

	err := windows.SetConsoleMode(
		handle,
		stdoutConsoleMode|flags,
	)
	if err != nil {
		return err
	}
	setColors()
	enabled = true

	return nil
}
