package log

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"
)

const separator = " | "
const bufferSize = 8192
const flushThreshold = bufferSize / 2

// Log is a log data source used for a specific part of your application,
// e.g. "web", "database", "api" or other categories. It can be connected
// to multiple outputs.
type Log struct {
	outputs []*output
}

// New creates a new Log.
func New() *Log {
	log := &Log{}
	sleepTime := 250 * time.Millisecond

	go func() {
		for {
			time.Sleep(sleepTime)
			log.Flush()
		}
	}()

	return log
}

// AddOutput adds an output to the log.
func (log *Log) AddOutput(writer io.Writer) {
	out := &output{
		writer:        writer,
		messageBuffer: make([]byte, 0, bufferSize),
	}

	log.outputs = append(log.outputs, out)
}

// Info writes non-critical information to the log.
// Unlike Error, it does not guarantee that the message will have been
// written persistenly to disk at the time this function returns.
func (log *Log) Info(values ...interface{}) {
	log.write(values...)
}

// Error writes critical information to the log.
// It will instantly flush the I/O buffers and guarantees that the message
// will have been written persistenly to disk at the time this function returns.
func (log *Log) Error(values ...interface{}) {
	log.write(values...)
	log.Flush()
}

// Flush forces the currently buffered data to be flushed to all outputs.
// A flush usually guarantees that the data has been written permanently to disk.
func (log *Log) Flush() {
	for _, output := range log.outputs {
		output.mutex.Lock()
		output.writer.Write(output.messageBuffer)
		output.messageBuffer = output.messageBuffer[:0]
		output.mutex.Unlock()
	}
}

// Write implements the io.Writer interface.
func (log *Log) Write(b []byte) (int, error) {
	log.write(bytes.TrimSpace(b))
	return len(b), nil
}

// write is the core function implementing the serialization of data types.
func (log *Log) write(values ...interface{}) {
	now := time.Now().Format(time.StampMilli)

	for _, output := range log.outputs {
		output.mutex.Lock()
		b := append(output.messageBuffer, now...)

		for _, value := range values {
			b = append(b, separator...)

			switch value := value.(type) {
			case string:
				b = append(b, value...)
			case int:
				b = strconv.AppendInt(b, int64(value), 10)
			case float64:
				b = strconv.AppendFloat(b, value, 'f', 5, 64)
			case float32:
				b = strconv.AppendFloat(b, float64(value), 'f', 5, 32)
			case byte:
				b = append(b, value)
			case []byte:
				b = append(b, value...)
			default:
				b = append(b, fmt.Sprint(value)...)
			}
		}

		b = append(b, '\n')

		if len(b) > flushThreshold {
			output.writer.Write(b)
			output.messageBuffer = b[:0]
		} else {
			output.messageBuffer = b
		}

		output.mutex.Unlock()
	}
}
