package log

import (
	"os"
	"testing"

	std "log"

	"go.uber.org/zap"
)

func BenchmarkAero(b *testing.B) {
	os.Remove("aero.log")

	log := New()
	log.AddOutput(File("aero.log"))

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("Hello World")
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
			log.Info("Hello World")
		}
	})
}

func BenchmarkZapSugar(b *testing.B) {
	os.Remove("zapsugar.log")

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"zapsugar.log"},
		ErrorOutputPaths: []string{"stderr"},
	}
	log, _ := config.Build()
	sugar := log.Sugar()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sugar.Info("Hello World")
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
			stdLog.Println("Hello World")
		}
	})
}

func BenchmarkAero5(b *testing.B) {
	os.Remove("aero5.log")

	log := New()
	log.AddOutput(File("aero5.log"))

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("Hello World", 1, 2, 3.14, 4)
		}
	})
}

func BenchmarkZap5(b *testing.B) {
	os.Remove("zap5.log")

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"zap5.log"},
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
				zap.Float64("c", 3.14),
				zap.Int("d", 4),
			)
		}
	})
}

func BenchmarkZapSugar5(b *testing.B) {
	os.Remove("zapsugar5.log")

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"zapsugar5.log"},
		ErrorOutputPaths: []string{"stderr"},
	}
	log, _ := config.Build()
	sugar := log.Sugar()

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sugar.Info("Hello World", 1, 2, 3.14, 4)
		}
	})
}

func BenchmarkStd5(b *testing.B) {
	os.Remove("std5.log")

	f, _ := os.Create("std5.log")
	stdLog := std.New(f, "", 0)

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			stdLog.Println("Hello World", 1, 2, 3.14, 4)
		}
	})
}
