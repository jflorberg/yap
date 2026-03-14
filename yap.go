package yap

import (
	"fmt"
	"io"
	"os"

	"github.com/jflorberg/yap/internal/severity"
)

// Yap is the main logger instance
type Yap struct {
	cfg Config
	out io.Writer
}

// Config controls logger behavior
type Config struct {
	Output            io.Writer // default: os.Stdout
	UseTimestamp      bool      // include timestamp
	UseCallerFile     bool      // include file:line
	UseCalledFunc     bool      // include function name
	UseSeverityColors bool      // include colors for severity
	UseSeverityStyles bool      // include styles for severity
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
	Output:            os.Stdout,
	UseTimestamp:      true,
	UseCallerFile:     true,
	UseCalledFunc:     true,
	UseSeverityColors: true,
	UseSeverityStyles: true,
})

// SetDefault sets the package-level default logger
func SetDefault(l *Yap) { std = l }

// Default returns the default logger, creating a fallback if not set
func Default() *Yap { return std }

// Log functions

func (l *Yap) Info(v ...any) {
	l.output(&severity.Info, v...)
}

func (l *Yap) Infof(format string, v ...any) {
	l.output(&severity.Info, fmt.Sprintf(format, v...))
}

func (l *Yap) Warn(v ...any) {
	l.output(&severity.Warn, v...)
}

func (l *Yap) Warnf(format string, v ...any) {
	l.output(&severity.Warn, fmt.Sprintf(format, v...))
}

func (l *Yap) Error(v ...any) {
	l.output(&severity.Error, v...)
}

func (l *Yap) Errorf(format string, v ...any) {
	l.output(&severity.Error, fmt.Sprintf(format, v...))
}

func (l *Yap) Debug(v ...any) {
	l.output(&severity.Debug, v...)
}

func (l *Yap) Debugf(format string, v ...any) {
	l.output(&severity.Debug, fmt.Sprintf(format, v...))
}

func (l *Yap) Print(v ...any) {
	l.output(nil, v...)
}

func (l *Yap) Printf(format string, v ...any) {
	l.output(nil, fmt.Sprintf(format, v...))
}

func (l *Yap) Fatal(v ...any) {
	l.output(&severity.Fatal, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Yap) Fatalf(format string, v ...any) {
	l.output(&severity.Fatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Package-level log functions that use the default logger

func Info(v ...any) {
	std.Info(v...)
}

func Infof(format string, v ...any) {
	std.Infof(format, v...)
}

func Warn(v ...any) {
	std.Warn(v...)
}

func Warnf(format string, v ...any) {
	std.Warnf(format, v...)
}

func Error(v ...any) {
	std.Error(v...)
}

func Errorf(format string, v ...any) {
	std.Errorf(format, v...)
}

func Debug(v ...any) {
	std.Debug(v...)
}

func Debugf(format string, v ...any) {
	std.Debugf(format, v...)
}

func Print(v ...any) {
	std.Print(v...)
}

func Printf(format string, v ...any) {
	std.Printf(format, v...)
}

// Fatal is equivalent to [Print] followed by a call to [os.Exit](1).
func Fatal(v ...any) {
	std.Fatal(v...)
}

// Fatalf is equivalent to [Printf] followed by a call to [os.Exit](1).
func Fatalf(format string, v ...any) {
	std.Fatalf(format, v...)
}
