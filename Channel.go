package log

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

const separator = " | "

// Channel is used for a specific part of your application, e.g. "web", "database", "api", etc.
type Channel struct {
	name    string
	outputs []*Output
}

// NewChannel ...
func NewChannel(name string) *Channel {
	channel := &Channel{
		name: name,
	}

	go func() {
		ticker := time.NewTicker(250 * time.Millisecond)
		defer ticker.Stop()

		for {
			<-ticker.C
			channel.Flush()
		}
	}()

	return channel
}

// AddOutput ...
func (channel *Channel) AddOutput(writer io.Writer) {
	output := &Output{
		writer: bufio.NewWriter(writer),
	}

	channel.outputs = append(channel.outputs, output)
}

// Info ...
func (channel *Channel) Info(values ...interface{}) {
	channel.write(values...)
}

// Error ...
func (channel *Channel) Error(values ...interface{}) {
	channel.write(values...)
	channel.Flush()
}

// Flush ...
func (channel *Channel) Flush() {
	for _, output := range channel.outputs {
		output.mutex.Lock()
		output.writer.Flush()
		output.mutex.Unlock()
	}
}

// write ...
func (channel *Channel) write(values ...interface{}) {
	now := time.Now().Format(time.StampMilli)

	for _, output := range channel.outputs {
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
