package log

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const separator = " | "
const bufferSize = 8192
const flushThreshold = bufferSize / 2

// Log is used for a specific part of your application, e.g. "web", "database", "api", etc.
type Log struct {
	outputs []*Output
}

// New ...
func New() *Log {
	log := &Log{}

	go func() {
		for {
			time.Sleep(250 * time.Millisecond)
			log.Flush()
		}
	}()

	return log
}

// AddOutput ...
func (log *Log) AddOutput(writer io.Writer) {
	output := &Output{
		writer:        writer,
		messageBuffer: make([]byte, 0, bufferSize),
	}

	log.outputs = append(log.outputs, output)
}

// Info ...
func (log *Log) Info(values ...interface{}) {
	log.write(values...)
}

// Error ...
func (log *Log) Error(values ...interface{}) {
	log.write(values...)
	log.Flush()
}

// Flush ...
func (log *Log) Flush() {
	for _, output := range log.outputs {
		output.mutex.Lock()
		output.writer.Write(output.messageBuffer)
		output.messageBuffer = output.messageBuffer[:0]
		output.mutex.Unlock()
	}
}

// write ...
func (log *Log) write(values ...interface{}) {
	now := time.Now().Format(time.StampMilli)

	for _, output := range log.outputs {
		output.mutex.Lock()
		b := append(output.messageBuffer, now...)

		for _, value := range values {
			b = append(b, separator...)

			switch value.(type) {
			case string:
				b = append(b, value.(string)...)
			case int:
				b = strconv.AppendInt(b, int64(value.(int)), 10)
			case float64:
				b = strconv.AppendFloat(b, value.(float64), 'f', 5, 64)
			case float32:
				b = strconv.AppendFloat(b, float64(value.(float32)), 'f', 5, 32)
			case byte:
				b = append(b, value.(byte))
			case []byte:
				b = append(b, value.([]byte)...)
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
