# Yet Another Printer (Yap)

Yet Another Printer (Yap) is a log package for go.

## Import

`go get github.com/jflorberg/yap`

## Snippet

Initialize the logger with configuration and set as default

```go
package main

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
	yap.Warning("This is a warning")
	yap.Error("Something went wrong")
}
```
