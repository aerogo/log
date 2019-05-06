# {name}

{go:header}

A logging system that allows you to connect one log to multiple outputs (e.g. file system).

## Installation

```go
go get github.com/aerogo/log
```

## Example

```go
hello := log.New()                     // Create a new log
hello.AddOutput(log.File("hello.log")) // Add an output

hello.Info("Hello World", 1, 2, 3)     // Write non-critical data (buffered)
hello.Error("Something went wrong")    // Force an immediate I/O flush for error messages
```

{go:footer}
