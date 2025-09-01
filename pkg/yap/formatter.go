package yap

import (
	"fmt"
	"runtime"
	"time"
)

// internal formatter
func (y *Yap) output(level string, v ...any) {
	var prefix string

	// timestamp
	if y.cfg.Timestamp {
		prefix += time.Now().Format("2006-01-02 15:04:05") + " "
	}

	// caller
	if y.cfg.Caller {
		_, file, line, ok := runtime.Caller(2) // 2 = skip this + caller
		if ok {
			prefix += fmt.Sprintf("%s:%d ", shortFile(file), line)
		}
	}

	// level
	prefix += fmt.Sprintf("[%s] ", level)

	// build final msg
	msg := fmt.Sprint(v...)
	output := prefix + msg + "\n"

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
