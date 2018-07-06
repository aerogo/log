# log

[![Godoc reference][godoc-image]][godoc-url]
[![Go report card][goreportcard-image]][goreportcard-url]
[![Tests][travis-image]][travis-url]
[![Code coverage][codecov-image]][codecov-url]
[![License][license-image]][license-url]

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

## Author

| [![Eduard Urbach on Twitter](https://gravatar.com/avatar/16ed4d41a5f244d1b10de1b791657989?s=70)](https://twitter.com/eduardurbach "Follow @eduardurbach on Twitter") |
|---|
| [Eduard Urbach](https://eduardurbach.com) |

[godoc-image]: https://godoc.org/github.com/aerogo/log?status.svg
[godoc-url]: https://godoc.org/github.com/aerogo/log
[goreportcard-image]: https://goreportcard.com/badge/github.com/aerogo/log
[goreportcard-url]: https://goreportcard.com/report/github.com/aerogo/log
[travis-image]: https://travis-ci.org/aerogo/log.svg?branch=master
[travis-url]: https://travis-ci.org/aerogo/log
[codecov-image]: https://codecov.io/gh/aerogo/log/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/aerogo/log
[license-image]: https://img.shields.io/badge/license-MIT-blue.svg
[license-url]: https://github.com/aerogo/log/blob/master/LICENSE
