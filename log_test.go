package log

import (
	"os"
	"runtime"
	"sync"
	"testing"
)

func bench(b *testing.B, fun func()) {
	b.ReportAllocs()

	wg := sync.WaitGroup{}
	tasks := runtime.NumCPU()
	wg.Add(tasks)

	b.ResetTimer()

	for i := 0; i < tasks; i++ {
		go func() {
			for n := 0; n < b.N; n++ {
				fun()
			}

			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkAeroLog(b *testing.B) {
	web := NewChannel("web")
	file, _ := os.Create("aero.log")
	web.AddOutput(file)

	bench(b, func() {
		web.Info("Hello World")
	})
}
