# {name}

{go:header}

An O(1) constant time logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).

{go:install}

## Example

```go
hello := log.New()                          // Create a new log
hello.AddWriter(log.File("hello.log"))      // Add a writer

hello.Info("Hello World %d %d %d", 1, 2, 3) // Write non-critical data (buffered)
hello.Error("Something went wrong")         // Force an immediate I/O flush for error messages
```

## Under the hood

All `Write` calls are queued into a channel that is read in a separate goroutine. Once the channel receives new data it writes the data to all registered outputs.

{go:footer}
