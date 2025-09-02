package style

import "fmt"

type ANSI_STYLE string

const (
	// Styles
	None      ANSI_STYLE = ""
	Bold      ANSI_STYLE = "\033[1m%s\033[0m"
	Underline ANSI_STYLE = "\033[4m%s\033[0m"
	Inverse   ANSI_STYLE = "\033[7m%s\033[0m"
)

func StyleText(s ANSI_STYLE, text string) string {
	if s == None {
		return text
	}
	return fmt.Sprintf(string(s), text)
}
