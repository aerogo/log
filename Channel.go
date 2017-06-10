package log

import (
	"io"
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
		writer: writer,
	}

	channel.outputs = append(channel.outputs, output)
}

// Info ...
func (channel *Channel) Info(values ...string) {
	for _, output := range channel.outputs {
		output.mutex.Lock()

		for _, value := range values {
			output.buffer.WriteString(value)
		}
		output.buffer.WriteByte('\n')

		output.sync()
		output.mutex.Unlock()
	}
}
