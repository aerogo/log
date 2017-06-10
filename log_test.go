package log

import (
	"runtime"
	"sync"
	"testing"

	"go.uber.org/zap"
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
	web.AddOutput(File("aero.log"))

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			web.Info("Hello World")
		}
	})
}

func BenchmarkZap(b *testing.B) {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"zap.log"},
		ErrorOutputPaths: []string{"stderr"},
	}
	log, _ := config.Build()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("Hello World")
		}
	})
}
