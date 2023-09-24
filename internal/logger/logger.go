package logger

import (
	"fmt"
	"runtime"
	"time"
)

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) MemoryUsage() {
	var m runtime.MemStats
	for {
		time.Sleep(3 * time.Second)
		runtime.ReadMemStats(&m)
		data := fmt.Sprintf("Current alloc: %.2f MB, total alloc: %.2f MB, GC cycles: %d ",
			float64(m.Alloc)/1024/1024,
			float64(m.TotalAlloc)/1024/1024,
			m.NumGC,
		)
		fmt.Printf("\r%s", data)
	}
}
