package color

var enabled = false

// Add wraps a given message in a given color.
// Init must be called first
func Add(color string, message string) string {
	return color + message + Default
}

// If Enabled, then you are good to go!
func Enabled() bool {
	return enabled
}
