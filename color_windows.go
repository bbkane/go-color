package color

import (
	"os"
	"syscall"
)

func Init() error {
	if initialized {
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
	setColors()
	initialized = true
	return nil
}
