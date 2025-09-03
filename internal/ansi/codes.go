package ansi

import "fmt"

type Color string

const (
	// Basic ANSI color codes
	NoColor Color = ""
	Reset   Color = "\033[0m"
	Black   Color = "\033[30m%s\033[0m"
	Red     Color = "\033[31m%s\033[0m"
	Green   Color = "\033[32m%s\033[0m"
	Yellow  Color = "\033[33m%s\033[0m"
	Blue    Color = "\033[34m%s\033[0m"
	Magenta Color = "\033[35m%s\033[0m"
	Cyan    Color = "\033[36m%s\033[0m"
	White   Color = "\033[37m%s\033[0m"
	Gray    Color = "\033[90m%s\033[0m"

	// Bright versions
	BrightBlack   Color = "\033[90m%s\033[0m"
	BrightRed     Color = "\033[91m%s\033[0m"
	BrightGreen   Color = "\033[92m%s\033[0m"
	BrightYellow  Color = "\033[93m%s\033[0m"
	BrightBlue    Color = "\033[94m%s\033[0m"
	BrightMagenta Color = "\033[95m%s\033[0m"
	BrightCyan    Color = "\033[96m%s\033[0m"
	BrightWhite   Color = "\033[97m%s\033[0m"
)

// Applies color to the provided string
func (c Color) Apply(text string) string {
	if c == NoColor {
		return text
	}
	return fmt.Sprintf(string(c), text)
}

type Style string

const (
	// Styles
	NoStyle   Style = ""
	Bold      Style = "\033[1m%s\033[0m"
	Underline Style = "\033[4m%s\033[0m"
	Inverse   Style = "\033[7m%s\033[0m"
)

// Applies color to the provided string
func (s Style) Apply(text string) string {
	if s == NoStyle {
		return text
	}
	return fmt.Sprintf(string(s), text)
}
