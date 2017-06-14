package log

import (
	"os"
	"testing"

	std "log"

	"go.uber.org/zap"
)

func BenchmarkAero(b *testing.B) {
	os.Remove("aero.log")

	web := NewChannel("web")
	web.AddOutput(File("aero.log"))

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			web.Info("Hello World", 1, 2, 3, 4)
		}
	})
}

func BenchmarkZap(b *testing.B) {
	os.Remove("zap.log")

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"zap.log"},
		ErrorOutputPaths: []string{"stderr"},
	}
	log, _ := config.Build()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info(
				"Hello World",
				zap.Int("a", 1),
				zap.Int("b", 2),
				zap.Int("c", 3),
				zap.Int("d", 4),
			)
		}
	})
}

func BenchmarkStd(b *testing.B) {
	os.Remove("std.log")

	f, _ := os.Create("std.log")
	stdLog := std.New(f, "", 0)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			stdLog.Println("Hello World", 1, 2, 3, 4)
		}
	})
}
