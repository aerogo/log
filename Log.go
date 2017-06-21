package log

import (
	"bufio"
	"fmt"
	"io"
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
		output.writer.WriteString(now)

		for _, value := range values {
			output.writer.WriteString(separator)
			fmt.Fprint(output.writer, value)
		}

		output.writer.WriteByte('\n')
		output.mutex.Unlock()
	}
}
