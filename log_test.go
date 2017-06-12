package log

import (
	"testing"

	"go.uber.org/zap"
)

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
