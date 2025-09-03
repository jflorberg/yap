package style

import (
	"github.com/jflorberg/yap/internal/ansi"
)

func Bold(text string) string      { return ansi.Bold.Apply(text) }
func Underline(text string) string { return ansi.Underline.Apply(text) }
func Inverse(text string) string   { return ansi.Inverse.Apply(text) }
