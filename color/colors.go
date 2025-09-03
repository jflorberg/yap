package color

import (
	"github.com/jflorberg/yap/internal/ansi"
)

func Black(text string) string         { return ansi.Black.Apply(text) }
func Red(text string) string           { return ansi.Red.Apply(text) }
func Green(text string) string         { return ansi.Green.Apply(text) }
func Yellow(text string) string        { return ansi.Yellow.Apply(text) }
func Blue(text string) string          { return ansi.Blue.Apply(text) }
func Magenta(text string) string       { return ansi.Magenta.Apply(text) }
func Cyan(text string) string          { return ansi.Cyan.Apply(text) }
func White(text string) string         { return ansi.White.Apply(text) }
func Gray(text string) string          { return ansi.Gray.Apply(text) }
func BrightBlack(text string) string   { return ansi.BrightBlack.Apply(text) }
func BrightRed(text string) string     { return ansi.BrightRed.Apply(text) }
func BrightGreen(text string) string   { return ansi.BrightGreen.Apply(text) }
func BrightYellow(text string) string  { return ansi.BrightYellow.Apply(text) }
func BrightBlue(text string) string    { return ansi.BrightBlue.Apply(text) }
func BrightMagenta(text string) string { return ansi.BrightMagenta.Apply(text) }
func BrightCyan(text string) string    { return ansi.BrightCyan.Apply(text) }
func BrightWhite(text string) string   { return ansi.BrightWhite.Apply(text) }
