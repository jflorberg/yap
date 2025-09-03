package main

import (
	"github.com/jflorberg/yap"
	"github.com/jflorberg/yap/color"
	"github.com/jflorberg/yap/style"
	"github.com/jflorberg/yap/symbols"
)

func init() {
	cfg := &yap.Config{
		UseTimestamp:      true,
		UseCallerFile:     true,
		UseSeverityColors: true,
		UseSeverityStyles: true,
		UseCalledFunc:     true,
	}
	log := yap.New(cfg)
	yap.SetDefault(log)
}

func main() {
	yap.Info("App started", symbols.Rocket)
	yap.Warn("This is a warning")
	yap.Error("Something went wrong")
	thisIsAFunc()

	yap.Print(color.Blue("Blue output"))
	yap.Print(style.Bold(color.Blue("Bold blue output")))

	yap.Fatal("This is fatal")
}

func thisIsAFunc() {
	yap.Debug("This is debug trace")
	yap.Debugf("This is %s with %s", "a debug trace", "with formatting")
	yap.Print("This is a regular print")
}
