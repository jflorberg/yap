# Yet Another Printer (Yap)

**Yap** is a lightweight, configurable logging package for Go — with support for timestamps, caller info, severity levels, and colorful/styled output.

---

## ✨ Features

- 🕒 Optional timestamps
- 📂 Caller file + line number
- 🔧 Function name tracing
- 🎨 Severity colors, styles and symbols
- ⚡ Simple API (`yap.Info("...")`)
- 🧩 Instance-based loggers for flexible configurations

---

## 📦 Installation

```bash
go get github.com/jflorberg/yap
```

## 🚀 Quick Start

### Global logger (default behavior)

Initialize a logger and set it as the package default:

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
	yap.Warn("This is a warning")
	yap.Error("Something went wrong")
}
```

### 🧩 Instance-based logger

You can also create independent logger instances with different configurations.

This is useful when different parts of your application require different logging styles (for example HTTP request logs vs application logs).

```go
package main

import (
	"github.com/jflorberg/yap"
)

func main() {

	httpLogger := yap.New(&yap.Config{
		UseTimestamp:      true,
		UseCallerFile:     false,
		UseCalledFunc:     false,
		UseSeverityColors: true,
		UseSeverityStyles: false,
	})

	httpLogger.Infof("GET /api/characters 200 12ms")
}
```

This logger does not affect the global default logger used by `yap.Info()` and friends.

## ⚙️ Configuration

```go
type Config struct {
	Output            io.Writer // default: os.Stdout
	UseTimestamp      bool      // include timestamp
	UseCallerFile     bool      // include file:line
	UseCalledFunc     bool      // include function name
	UseSeverityColors bool      // include colors for severity
	UseSeverityStyles bool      // include styles for severity
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
