package log

import (
	"bufio"
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
func (channel *Channel) Info(values ...string) {
	now := time.Now().Format(time.StampMilli)

	for _, output := range channel.outputs {
		output.mutex.Lock()

		output.writer.WriteString(now)
		output.writer.WriteByte('\t')

		for num, value := range values {
			output.writer.WriteString(value)

			if num >= 1 {
				output.writer.WriteByte(' ')
			}
		}
		output.writer.WriteByte('\n')

		output.mutex.Unlock()
	}
}
