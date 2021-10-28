package color

// These will be set to the proper color codes if Init() returns successfully
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

var initialized = false

// setColors actually sets the colors - called from Init()
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

// Add wraps a given message in a given color.
// Init must be called first
func Add(color, message string) string {
	return color + message + Reset
}
