package yap

import (
	"fmt"
	"runtime"
	"time"

	"github.com/jflorberg/yap/pkg/color"
	"github.com/jflorberg/yap/pkg/severity"
	"github.com/jflorberg/yap/pkg/style"
)

// internal formatter
func (y *Yap) output(severity *severity.Severity, v ...any) {
	var prefix string

	// timestamp
	if y.cfg.UseTimestamp {
		prefix += time.Now().Format("2006-01-02 15:04:05") + " "
	}

	// caller
	if y.cfg.UseCallerFile {
		pc, file, line, ok := runtime.Caller(2) // 2 = skip this + caller
		if ok {
			prefix += fmt.Sprintf("%s:%d ", shortFile(file), line)
		}

		if y.cfg.UseCalledFunc {
			fn := runtime.FuncForPC(pc)
			if fn != nil {
				prefix += fmt.Sprintf("%s() ", shortFunc(fn.Name()))
			}
		}
	}

	if severity != nil {
		if y.cfg.UseSeverityStyles && y.cfg.UseSeverityColors {
			prefix += fmt.Sprintf("[%s] ", style.StyleText(severity.Style, color.ColorText(severity.Color, severity.Name)))
		} else if y.cfg.UseSeverityColors {
			prefix += fmt.Sprintf("[%s] ", color.ColorText(severity.Color, severity.Name))
		} else {
			prefix += fmt.Sprintf("[%s] ", severity.Name)
		}
	}

	// build final msg
	msg := fmt.Sprint(v...)
	output := fmt.Sprintf("%s: %s \n", prefix, msg)

	_, _ = y.out.Write([]byte(output))
}

// utility: strip path, keep only file name
func shortFile(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}

func shortFunc(name string) string {
	// strip package path, keep only function part
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' || name[i] == '.' {
			return name[i+1:]
		}
	}
	return name
}
