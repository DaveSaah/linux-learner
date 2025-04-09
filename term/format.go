package term

// Terminal formatting constants
const (
	reset     = "\033[0m"
	bold      = "\033[1m"
	italic    = "\033[3m"
	underline = "\033[4m"
)

// Color struct defines the ANSI color codes for text and background
type Color struct {
	TextColor string
	BgColor   string
}

// Predefined color values
var (
	Black   = Color{TextColor: "\033[30m"}
	Red     = Color{TextColor: "\033[31m"}
	Green   = Color{TextColor: "\033[32m"}
	Yellow  = Color{TextColor: "\033[33m"}
	Blue    = Color{TextColor: "\033[34m"}
	Magenta = Color{TextColor: "\033[35m"}
	Cyan    = Color{TextColor: "\033[36m"}
	White   = Color{TextColor: "\033[37m"}

	// Background colors
	BgBlack   = Color{BgColor: "\033[40m"}
	BgRed     = Color{BgColor: "\033[41m"}
	BgGreen   = Color{BgColor: "\033[42m"}
	BgYellow  = Color{BgColor: "\033[43m"}
	BgBlue    = Color{BgColor: "\033[44m"}
	BgMagenta = Color{BgColor: "\033[45m"}
	BgCyan    = Color{BgColor: "\033[46m"}
	BgWhite   = Color{BgColor: "\033[47m"}
)

// Bold returns the ANSI escape code for bold text
func Bold(text string) string {
	return bold + text + reset
}

// Italic returns the ANSI escape code for italic text
func Italic(text string) string {
	return italic + text + reset
}

// Underline returns the ANSI escape code for underlined text
func Underline(text string) string {
	return underline + text + reset
}

// ColorText applies the given text color to the text
func ColorText(text string, color string) string {
	return getColorCode(color) + text + reset
}

// BgColorText applies the given background color to the text
func BgColorText(text string, color string) string {
	return getBgColorCode(color) + text + reset
}

// ResetTextStyle resets all formatting styles applied to text
func ResetTextStyle() string {
	return reset
}

// Helper function to get color code by string
func getColorCode(colorName string) string {
	switch colorName {
	case "black":
		return Black.TextColor
	case "red":
		return Red.TextColor
	case "green":
		return Green.TextColor
	case "yellow":
		return Yellow.TextColor
	case "blue":
		return Blue.TextColor
	case "magenta":
		return Magenta.TextColor
	case "cyan":
		return Cyan.TextColor
	case "white":
		return White.TextColor
	default:
		return "" // Invalid color
	}
}

// Helper function to get background color code by string
func getBgColorCode(colorName string) string {
	switch colorName {
	case "black":
		return BgBlack.BgColor
	case "red":
		return BgRed.BgColor
	case "green":
		return BgGreen.BgColor
	case "yellow":
		return BgYellow.BgColor
	case "blue":
		return BgBlue.BgColor
	case "magenta":
		return BgMagenta.BgColor
	case "cyan":
		return BgCyan.BgColor
	case "white":
		return BgWhite.BgColor
	default:
		return "" // Invalid background color
	}
}
