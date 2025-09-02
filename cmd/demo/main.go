package main

import "github.com/jflorberg/yap"

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
	yap.Info("App started")
	yap.Warn("This is a warning")
	yap.Error("Something went wrong")
	thisIsAFunc()

	yap.Fatal("This is fatal")
}

func thisIsAFunc() {
	yap.Debug("This is debug trace")
	yap.Debugf("This is debug trace: %s:%2f:%d", "with formatting", 10.25, 80)
	yap.Print("Test Print")
}
