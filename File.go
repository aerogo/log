package log

import "os"

// File provides a simple interface to create a log file.
// The given file path must be writable, otherwise it will panic.
func File(path string) *os.File {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	return file
}
