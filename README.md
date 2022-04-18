# go-pirev

[![godoc-badge]][godoc]
[![release-badge]][latest-release]
[![license-badge]][license]
[![goreport-badge]][goreport]

## Raspberry Pi revision code parser in Go

A Go library for parsing and extracting hardware information from a Raspberry Pi revision code as structured data or raw bit values.

Also includes a basic CLI tool that uses this library to display the details of a Raspberry Pi revision code.

## Installation

### Library

In the directory of your module run:

```sh
go get github.com/antichris/go-pirev@latest
```

### CLI tool

```sh
go install github.com/antichris/go-pirev/cmd/pirev@latest
```

## License

The source code of this project is released under [Mozilla Public License Version 2.0][mpl]. See [LICENSE].

[mpl]: https://www.mozilla.org/en-US/MPL/2.0/
	"Mozilla Public License, version 2.0"

[license]: LICENSE

[godoc]: https://pkg.go.dev/github.com/antichris/go-pirev
[latest-release]: https://github.com/antichris/go-pirev/releases/latest
[goreport]: https://goreportcard.com/report/github.com/antichris/go-pirev

[godoc-badge]: https://godoc.org/github.com/antichris/go-pirev?status.svg
[release-badge]: https://img.shields.io/github/release/antichris/go-pirev
[license-badge]: https://img.shields.io/github/license/antichris/go-pirev
[goreport-badge]: https://goreportcard.com/badge/github.com/antichris/go-pirev?status.svg
