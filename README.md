# Yet Another Printer (Yap)

**Yap** is a lightweight, configurable logging package for Go — with support for timestamps, caller info, severity levels, and colorful/styled output.

---

## ✨ Features

- 🕒 Optional timestamps
- 📂 Caller file + line number
- 🔧 Function name tracing
- 🎨 Severity colors, styles and symbols
- ⚡ Simple API (`yap.Info("...")`)

---

## 📦 Installation

```bash
go get github.com/jflorberg/yap
```

## 🚀 Quick Start

Initialize the logger with configuration and set it as the default:

```go
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
	yap.Warning("This is a warning")
	yap.Error("Something went wrong")
}
```

## 📖 Background

The reason I created this logger was not only to try something new and publish my first Go package, but also because I was tired of using log packages that were difficult to customize and made terminal logs hard to read.

I prefer using colors and severity levels to get a quicker grasp of what is going on, while keeping the syntax simple and familiar. With functions like `Info`, `Warn`, `Error`, and `Debug`, Yap makes it easy to log for different purposes without clutter or complexity.

## 🛠 Roadmap

- Severity colors & styles customization
- JSON log output
- Custom formatters
- File logging

## 📜 License

MIT — feel free to use, modify, and contribute!
