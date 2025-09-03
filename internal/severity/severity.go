package severity

import (
	"github.com/jflorberg/yap/internal/ansi"
)

type Severity struct {
	Name  string
	Color ansi.Color
	Style ansi.Style
}

var Info Severity = Severity{
	Name:  "INFO",
	Color: ansi.BrightBlue,
	Style: ansi.NoStyle,
}

var Warn Severity = Severity{
	Name:  "WARN",
	Color: ansi.BrightYellow,
	Style: ansi.NoStyle,
}

var Error Severity = Severity{
	Name:  "ERROR",
	Color: ansi.Red,
	Style: ansi.NoStyle,
}

var Debug Severity = Severity{
	Name:  "DEBUG",
	Color: ansi.Gray,
	Style: ansi.NoStyle,
}

var Fatal Severity = Severity{
	Name:  "FATAL",
	Color: ansi.BrightRed,
	Style: ansi.Bold,
}
