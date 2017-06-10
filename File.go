package log

import "os"

// File provides a simple interface to create a log file.
// Ignores creation errors for the sake of simplicity.
func File(path string) *os.File {
	file, _ := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	return file
}
