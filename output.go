package log

import (
	"io"
	"sync"
)

// output represents a buffered device that can be used to write log messages to.
type output struct {
	writer        io.Writer
	mutex         sync.Mutex
	messageBuffer []byte
}
