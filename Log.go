package log

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"time"
)

const separator = " | "

// Log is used for a specific part of your application, e.g. "web", "database", "api", etc.
type Log struct {
	outputs []*Output
}

// New ...
func New() *Log {
	log := &Log{}

	go func() {
		ticker := time.NewTicker(250 * time.Millisecond)
		defer ticker.Stop()

		for {
			<-ticker.C
			log.Flush()
		}
	}()

	return log
}

// AddOutput ...
func (log *Log) AddOutput(writer io.Writer) {
	output := &Output{
		writer: bufio.NewWriter(writer),
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
		output.writer.Flush()
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
				b = strconv.AppendFloat(b, value.(float64), 'f', 3, 64)
			case float32:
				b = strconv.AppendFloat(b, float64(value.(float32)), 'f', 3, 32)
			case byte:
				b = append(b, value.(byte))
			case []byte:
				b = append(b, value.([]byte)...)
			default:
				b = append(b, fmt.Sprint(value)...)
			}
		}

		b = append(b, '\n')
		output.writer.Write(b)
		output.messageBuffer = b[:0]
		output.mutex.Unlock()
	}
}
