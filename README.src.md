# {name}

{go:header}

A logging system that allows you to connect one log to multiple writers (e.g. 2 files and 1 TCP connection).

{go:install}

## Example

```go
hello := log.New()                          // Create a new log
hello.AddWriter(log.File("hello.log"))      // Add a writer

hello.Info("Hello World %d %d %d", 1, 2, 3) // Write non-critical data (buffered)
hello.Error("Something went wrong")         // Force an immediate I/O flush for error messages
```

{go:footer}
