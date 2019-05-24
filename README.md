# log

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

An O(1) constant time logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).

## Installation

```shell
go get -u github.com/aerogo/log/...
```

## Example

```go
hello := log.New()                          // Create a new log
hello.AddWriter(log.File("hello.log"))      // Add a writer

hello.Info("Hello World %d %d %d", 1, 2, 3) // Write non-critical data (buffered)
hello.Error("Something went wrong")         // Force an immediate I/O flush for error messages
```

## Under the hood

All `Write` calls are queued into a channel that is read in a separate goroutine. Once the channel receives new data it writes the data to all registered outputs.

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Cedric Fung](https://avatars3.githubusercontent.com/u/2269238?s=70&v=4)](https://github.com/cedricfung) | [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars3.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- | --- |
| [Cedric Fung](https://github.com/cedricfung) | [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://github.com/users/akyoto/sponsorship)

[godoc-image]: https://godoc.org/github.com/aerogo/log?status.svg
[godoc-url]: https://godoc.org/github.com/aerogo/log
[report-image]: https://goreportcard.com/badge/github.com/aerogo/log
[report-url]: https://goreportcard.com/report/github.com/aerogo/log
[tests-image]: https://cloud.drone.io/api/badges/aerogo/log/status.svg
[tests-url]: https://cloud.drone.io/aerogo/log
[coverage-image]: https://codecov.io/gh/aerogo/log/graph/badge.svg
[coverage-url]: https://codecov.io/gh/aerogo/log
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
[sponsor-url]: https://github.com/users/akyoto/sponsorship
