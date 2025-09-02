package color

import (
	"fmt"
)

type ANSI_COLOR string

const (
	// Basic ANSI color codes
	None    ANSI_COLOR = ""
	Reset   ANSI_COLOR = "\033[0m"
	Black   ANSI_COLOR = "\033[30m%s\033[0m"
	Red     ANSI_COLOR = "\033[31m%s\033[0m"
	Green   ANSI_COLOR = "\033[32m%s\033[0m"
	Yellow  ANSI_COLOR = "\033[33m%s\033[0m"
	Blue    ANSI_COLOR = "\033[34m%s\033[0m"
	Magenta ANSI_COLOR = "\033[35m%s\033[0m"
	Cyan    ANSI_COLOR = "\033[36m%s\033[0m"
	White   ANSI_COLOR = "\033[37m%s\033[0m"
	Gray               = "\033[90m%s\033[0m"

	// Bright versions
	BrightBlack   ANSI_COLOR = "\033[90m%s\033[0m"
	BrightRed     ANSI_COLOR = "\033[91m%s\033[0m"
	BrightGreen   ANSI_COLOR = "\033[92m%s\033[0m"
	BrightYellow  ANSI_COLOR = "\033[93m%s\033[0m"
	BrightBlue    ANSI_COLOR = "\033[94m%s\033[0m"
	BrightMagenta ANSI_COLOR = "\033[95m%s\033[0m"
	BrightCyan    ANSI_COLOR = "\033[96m%s\033[0m"
	BrightWhite   ANSI_COLOR = "\033[97m%s\033[0m"
)

func ColorText(c ANSI_COLOR, text string) string {
	if c == None {
		return text
	}
	return fmt.Sprintf(string(c), text)
}
