package log

import (
	"fmt"
	"io"
	"sync/atomic"
)

const (
	// How many log messages can be buffered until the call blocks
	bufferCapacity = 1024
)

// Log is a log data source used for a specific part of your application,
// e.g. "web", "database", "api" or other categories. It can be connected
// to multiple writers.
type Log struct {
	writers  atomic.Value
	messages chan []byte
}

// New creates a new Log.
func New() *Log {
	log := &Log{
		messages: make(chan []byte, bufferCapacity),
	}

	log.writers.Store([]io.Writer{})

	go func() {
		for msg := range log.messages {
			log.write(msg)
		}
	}()

	return log
}

// AddWriter adds an output to the log.
func (log *Log) AddWriter(writer io.Writer) {
	newWriters := append(log.writers.Load().([]io.Writer), writer)
	log.writers.Store(newWriters)
}

// Info writes non-critical information to the log.
// Unlike Error, it does not guarantee that the message will have been
// written persistenly to disk at the time this function returns.
func (log *Log) Info(format string, values ...interface{}) {
	fmt.Fprintf(log, format+"\n", values...)
}

// Error writes critical information to the log.
// It will instantly flush the I/O buffers and guarantees that the message
// will have been written persistenly to disk at the time this function returns.
func (log *Log) Error(format string, values ...interface{}) {
	fmt.Fprintf(log, format+"\n", values...)
	// TODO: Flush.
}

// Write implements the io.Writer interface.
// As long as buffer capacity is available,
// this call will not block and have O(1) behaviour,
// regardless of how many writers are used.
func (log *Log) Write(b []byte) (n int, err error) {
	tmp := make([]byte, len(b))
	copy(tmp, b)
	log.messages <- tmp
	return len(b), nil
}

// write writes the given slice of bytes to all registered writers immediately.
func (log *Log) write(b []byte) {
	for _, writer := range log.writers.Load().([]io.Writer) {
		_, _ = writer.Write(b)
	}
}
