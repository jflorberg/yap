package main

import "github.com/jflorberg/yap/pkg/yap"

func init() {
	cfg := &yap.Config{
		Timestamp: false,
		Caller:    true,
	}
	log := yap.New(cfg)
	yap.SetDefault(log)
}

func main() {
	yap.Info("App started")
	yap.Warning("This is a warning")
	yap.Error("Something went wrong")
}
