package log

import (
	"bytes"
	"io"
	"sync"
)

// Output ...
type Output struct {
	buffer bytes.Buffer
	writer io.Writer
	mutex  sync.Mutex
}

// sync ...
func (output *Output) sync() {
	data := output.buffer.Bytes()
	output.writer.Write(data)
	output.buffer.Reset()
}
