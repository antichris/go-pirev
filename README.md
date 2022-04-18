# go-pirev

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

The source code of this project is released under [Mozilla Public License Version 2.0][mpl]. See [LICENSE](LICENSE).

[mpl]: https://www.mozilla.org/en-US/MPL/2.0/
	"Mozilla Public License, version 2.0"
