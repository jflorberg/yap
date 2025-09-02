package severity

import (
	"github.com/jflorberg/yap/pkg/color"
	"github.com/jflorberg/yap/pkg/style"
)

type Severity struct {
	Name  string
	Color color.ANSI_COLOR
	Style style.ANSI_STYLE
}

var Info Severity = Severity{
	Name:  "INFO",
	Color: color.BrightBlue,
	Style: style.None,
}

var Warn Severity = Severity{
	Name:  "WARN",
	Color: color.BrightYellow,
	Style: style.None,
}

var Error Severity = Severity{
	Name:  "ERROR",
	Color: color.Red,
	Style: style.None,
}

var Debug Severity = Severity{
	Name:  "DEBUG",
	Color: color.Gray,
	Style: style.None,
}

var Fatal Severity = Severity{
	Name:  "FATAL",
	Color: color.BrightRed,
	Style: style.Bold,
}
