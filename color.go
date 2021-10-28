package color

var (
	Reset  = ""
	Bold   = ""
	Red    = ""
	Green  = ""
	Yellow = ""
	Blue   = ""
	Purple = ""
	Cyan   = ""
	Gray   = ""
	White  = ""
)

func setColors() {
	Reset = "\033[0m"
	Bold = "\033[1m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Purple = "\033[35m"
	Cyan = "\033[36m"
	Gray = "\033[37m"
	White = "\033[97m"
}

// Ize is an alias for the Colorize function
func Ize(color, message string) string {
	return Colorize(color, message)
}

// Colorize wraps a given message in a given color.
func Colorize(color, message string) string {
	return color + message + Reset
}
