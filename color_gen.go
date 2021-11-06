package color

// These will be set to the proper color codes if Enable() returns successfully.
// generated from gen_colors/gen_color.py
var (
	// Default - Returns all attributes to the default state prior to modification
	Default = ""

	// Bold - Applies brightness/intensity flag to foreground color
	Bold = ""

	// Nobold - Removes brightness/intensity flag from foreground color
	Nobold = ""

	// Underline - Adds underline
	Underline = ""

	// Nounderline - Removes underline
	Nounderline = ""

	// Negative - Swaps foreground and background colors
	Negative = ""

	// Positive - Returns foreground/background to normal
	Positive = ""

	// ForegroundBlack - Applies non-bold/bright black to foreground
	ForegroundBlack = ""

	// ForegroundRed - Applies non-bold/bright red to foreground
	ForegroundRed = ""

	// ForegroundGreen - Applies non-bold/bright green to foreground
	ForegroundGreen = ""

	// ForegroundYellow - Applies non-bold/bright yellow to foreground
	ForegroundYellow = ""

	// ForegroundBlue - Applies non-bold/bright blue to foreground
	ForegroundBlue = ""

	// ForegroundMagenta - Applies non-bold/bright magenta to foreground
	ForegroundMagenta = ""

	// ForegroundCyan - Applies non-bold/bright cyan to foreground
	ForegroundCyan = ""

	// ForegroundWhite - Applies non-bold/bright white to foreground
	ForegroundWhite = ""

	// ForegroundExtended - Applies extended color value to the foreground (see details below)
	ForegroundExtended = ""

	// ForegroundDefault - Applies only the foreground portion of the defaults (see 0)
	ForegroundDefault = ""

	// BackgroundBlack - Applies non-bold/bright black to background
	BackgroundBlack = ""

	// BackgroundRed - Applies non-bold/bright red to background
	BackgroundRed = ""

	// BackgroundGreen - Applies non-bold/bright green to background
	BackgroundGreen = ""

	// BackgroundYellow - Applies non-bold/bright yellow to background
	BackgroundYellow = ""

	// BackgroundBlue - Applies non-bold/bright blue to background
	BackgroundBlue = ""

	// BackgroundMagenta - Applies non-bold/bright magenta to background
	BackgroundMagenta = ""

	// BackgroundCyan - Applies non-bold/bright cyan to background
	BackgroundCyan = ""

	// BackgroundWhite - Applies non-bold/bright white to background
	BackgroundWhite = ""

	// BackgroundExtended - Applies extended color value to the background (see details below)
	BackgroundExtended = ""

	// BackgroundDefault - Applies only the background portion of the defaults (see 0)
	BackgroundDefault = ""

	// BrightForegroundBlack - Applies bold/bright black to foreground
	BrightForegroundBlack = ""

	// BrightForegroundRed - Applies bold/bright red to foreground
	BrightForegroundRed = ""

	// BrightForegroundGreen - Applies bold/bright green to foreground
	BrightForegroundGreen = ""

	// BrightForegroundYellow - Applies bold/bright yellow to foreground
	BrightForegroundYellow = ""

	// BrightForegroundBlue - Applies bold/bright blue to foreground
	BrightForegroundBlue = ""

	// BrightForegroundMagenta - Applies bold/bright magenta to foreground
	BrightForegroundMagenta = ""

	// BrightForegroundCyan - Applies bold/bright cyan to foreground
	BrightForegroundCyan = ""

	// BrightForegroundWhite - Applies bold/bright white to foreground
	BrightForegroundWhite = ""

	// BrightBackgroundBlack - Applies bold/bright black to background
	BrightBackgroundBlack = ""

	// BrightBackgroundRed - Applies bold/bright red to background
	BrightBackgroundRed = ""

	// BrightBackgroundGreen - Applies bold/bright green to background
	BrightBackgroundGreen = ""

	// BrightBackgroundYellow - Applies bold/bright yellow to background
	BrightBackgroundYellow = ""

	// BrightBackgroundBlue - Applies bold/bright blue to background
	BrightBackgroundBlue = ""

	// BrightBackgroundMagenta - Applies bold/bright magenta to background
	BrightBackgroundMagenta = ""

	// BrightBackgroundCyan - Applies bold/bright cyan to background
	BrightBackgroundCyan = ""

	// BrightBackgroundWhite - Applies bold/bright white to background
	BrightBackgroundWhite = ""
)

func setColors() {
	Default = "\033[0m"
	Bold = "\033[1m"
	Nobold = "\033[22m"
	Underline = "\033[4m"
	Nounderline = "\033[24m"
	Negative = "\033[7m"
	Positive = "\033[27m"
	ForegroundBlack = "\033[30m"
	ForegroundRed = "\033[31m"
	ForegroundGreen = "\033[32m"
	ForegroundYellow = "\033[33m"
	ForegroundBlue = "\033[34m"
	ForegroundMagenta = "\033[35m"
	ForegroundCyan = "\033[36m"
	ForegroundWhite = "\033[37m"
	ForegroundExtended = "\033[38m"
	ForegroundDefault = "\033[39m"
	BackgroundBlack = "\033[40m"
	BackgroundRed = "\033[41m"
	BackgroundGreen = "\033[42m"
	BackgroundYellow = "\033[43m"
	BackgroundBlue = "\033[44m"
	BackgroundMagenta = "\033[45m"
	BackgroundCyan = "\033[46m"
	BackgroundWhite = "\033[47m"
	BackgroundExtended = "\033[48m"
	BackgroundDefault = "\033[49m"
	BrightForegroundBlack = "\033[90m"
	BrightForegroundRed = "\033[91m"
	BrightForegroundGreen = "\033[92m"
	BrightForegroundYellow = "\033[93m"
	BrightForegroundBlue = "\033[94m"
	BrightForegroundMagenta = "\033[95m"
	BrightForegroundCyan = "\033[96m"
	BrightForegroundWhite = "\033[97m"
	BrightBackgroundBlack = "\033[100m"
	BrightBackgroundRed = "\033[101m"
	BrightBackgroundGreen = "\033[102m"
	BrightBackgroundYellow = "\033[103m"
	BrightBackgroundBlue = "\033[104m"
	BrightBackgroundMagenta = "\033[105m"
	BrightBackgroundCyan = "\033[106m"
	BrightBackgroundWhite = "\033[107m"
}
