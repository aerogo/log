package log

import (
	"io"
	"sync"
)

// Output ...
type Output struct {
	writer        io.Writer
	mutex         sync.Mutex
	messageBuffer []byte
}
