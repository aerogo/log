package log_test

import (
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/aerogo/log"
	"github.com/stretchr/testify/assert"
)

// writerWithError errors the Write call after `successfulWrites` writes.
type writerWithError struct {
	io.Writer

	countWrites      int
	successfulWrites int
}

func (writer *writerWithError) Write(buffer []byte) (int, error) {
	if writer.countWrites == writer.successfulWrites {
		return 0, errors.New("Artificial error")
	}

	writer.countWrites++
	return writer.Write(buffer)
}

// zeroWriter always writes zero bytes.
type zeroWriter struct {
	io.Writer
}

func (writer *zeroWriter) Write(buffer []byte) (int, error) {
	return 0, nil
}

func TestInfoError(t *testing.T) {
	fileWriter := log.File("hello.log")
	defer os.Remove("hello.log")
	defer fileWriter.Close()

	errorWriter := &writerWithError{
		Writer: fileWriter,
	}

	zero := &zeroWriter{
		Writer: fileWriter,
	}

	hello := log.New()
	hello.AddWriter(fileWriter)
	hello.AddWriter(errorWriter)
	hello.AddWriter(zero)

	hello.Info(
		"Info message %d %f %f %s",
		1,
		float32(3.14),
		3.14,
		"some text",
	)

	hello.Error("Oh noes %s", "Something went wrong")
	time.Sleep(500 * time.Millisecond)
}

func TestInvalidFilePath(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	log.File("")
}

func BenchmarkInfo(b *testing.B) {
	defer os.Remove("hello.log")

	hello := log.New()
	hello.AddWriter(log.File("hello.log"))

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			hello.Info("Hello World")
		}
	})
}
