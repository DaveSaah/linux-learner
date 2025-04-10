package term

// ANSI style escape sequences for text formatting
const (
	reset     = "\033[0m"
	bold      = "\033[1m"
	italic    = "\033[3m"
	underline = "\033[4m"
)

// Style returns ANSI escape codes for text styles
func Bold(text string) string      { return bold + text + reset }
func Italic(text string) string    { return italic + text + reset }
func Underline(text string) string { return underline + text + reset }
func ResetTextStyle() string       { return reset }

// BoldUnderline returns the text formatted with both bold and underline styles
func BoldUnderline(text string) string {
	return bold + underline + text + reset
}

// ansi holds the escape codes for text and background color
type ansi struct {
	fg string // foreground
	bg string // background
}

// Colors provides convenient access to text and background colors
var Colors = struct {
	Black   ansi
	Red     ansi
	Green   ansi
	Yellow  ansi
	Blue    ansi
	Magenta ansi
	Cyan    ansi
	White   ansi
}{
	Black:   ansi{fg: "\033[30m", bg: "\033[40m"},
	Red:     ansi{fg: "\033[31m", bg: "\033[41m"},
	Green:   ansi{fg: "\033[32m", bg: "\033[42m"},
	Yellow:  ansi{fg: "\033[33m", bg: "\033[43m"},
	Blue:    ansi{fg: "\033[34m", bg: "\033[44m"},
	Magenta: ansi{fg: "\033[35m", bg: "\033[45m"},
	Cyan:    ansi{fg: "\033[36m", bg: "\033[46m"},
	White:   ansi{fg: "\033[37m", bg: "\033[47m"},
}

// ColorText applies the foreground color from an ansi color struct.
func ColorText(text string, color ansi) string {
	return color.fg + text + reset
}

// BgColorText applies the background color from an ansi color struct.
func BgColorText(text string, color ansi) string {
	return color.bg + text + reset
}
