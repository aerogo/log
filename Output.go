package log

import (
	"bufio"
	"sync"
)

// Output ...
type Output struct {
	writer        *bufio.Writer
	mutex         sync.Mutex
	messageBuffer []byte
}
