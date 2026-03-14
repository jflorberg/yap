package main

import (
	"github.com/jflorberg/yap"
)

func init() {
	cfg := &yap.Config{
		UseTimestamp:      true,
		UseCallerFile:     true,
		UseCalledFunc:     true,
		UseSeverityColors: true,
		UseSeverityStyles: true,
	}

	log := yap.New(cfg)
	yap.SetDefault(log)
}

func main() {
	yap.Info("App started")
	yap.Warn("This is a warning")
	yap.Error("Something went wrong")

	httpLogger := yap.New(&yap.Config{
		UseTimestamp:      true,
		UseCallerFile:     false,
		UseCalledFunc:     false,
		UseSeverityColors: true,
		UseSeverityStyles: false,
	})

	httpLogger.Infof("GET /api/characters 200 12ms")
}
