package log

import "bytes"

// Channel ...
type Channel struct {
	Name string
	log  *Log
}

// NewChannel creates a new channel.
func (log *Log) NewChannel(name string) *Channel {
	return &Channel{
		Name: "[" + name + "]",
		log:  log,
	}
}

// Info ...
func (channel *Channel) Info(values ...interface{}) {
	values = append([]interface{}{channel.Name}, values...)
	channel.log.Info(values...)
}

// Error ...
func (channel *Channel) Error(values ...interface{}) {
	values = append([]interface{}{channel.Name}, values...)
	channel.log.Error(values...)
}

// Write implements the io.Writer interface.
func (channel *Channel) Write(b []byte) (int, error) {
	channel.log.Info(channel.Name, bytes.TrimSpace(b))
	return len(b), nil
}
