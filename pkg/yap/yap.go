package yap

import (
	"io"
	"os"
)

// Yap is the main logger instance
type Yap struct {
	cfg Config
	out io.Writer
}

// Config controls logger behavior
type Config struct {
	Timestamp bool      // include timestamp
	Caller    bool      // include file:line
	Output    io.Writer // default: os.Stdout
}

// New creates a new Yap logger with config
func New(cfg *Config) *Yap {
	if cfg == nil {
		cfg = &Config{}
	}
	if cfg.Output == nil {
		cfg.Output = os.Stdout
	}
	return &Yap{
		cfg: *cfg,
		out: cfg.Output,
	}
}

var std = New(&Config{
	Timestamp: true,
	Caller:    true,
})

// SetDefault sets the package-level default logger
func SetDefault(l *Yap) { std = l }

// Default returns the default logger, creating a fallback if not set
func Default() *Yap { return std }

func Info(v ...any) {
	std.output("INFO", v...)
}

func Warning(v ...any) {
	std.output("WARN", v...)
}

func Error(v ...any) {
	std.output("ERROR", v...)
}
