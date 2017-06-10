package log

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

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

// Warn ...
func (channel *Channel) Warn(values ...interface{}) {
	channel.write(values...)
}

// Error ...
func (channel *Channel) Error(values ...interface{}) {
	channel.write(values...)

	// Flush on errors
	for _, output := range channel.outputs {
		output.writer.Flush()
	}
}

// write ...
func (channel *Channel) write(values ...interface{}) {
	now := time.Now().Format(time.StampMilli)

	for _, output := range channel.outputs {
		output.mutex.Lock()

		output.writer.WriteString(now)
		output.writer.WriteByte('\t')

		fmt.Fprint(output.writer, values...)
		output.writer.WriteByte('\n')

		output.mutex.Unlock()
	}
}
