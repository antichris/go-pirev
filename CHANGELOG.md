# Changelog

This project adheres to [Semantic Versioning][semver2].


## [0.5.0] - 2023-02-25

### Changed

- Go version in `go.mod` to the oldest stable (1.19)
- All module dependencies to the latest
- The order of `struct` fields for reduced memory consumption in
	- legacy `Info`
	- tests
- Table driven test cases to be defined inline, as the parameter for `range`
- Tests of `Stringer`s to be terser
- Empty `interface` to be spelled as `any`


## [0.4.0] - 2023-01-29

### Added

- The CM4S model

### Changed

- `bits.Field` bitmask type to `bits.Value` from `uint8`
- `strings_test` failure cases to use the maximum value of each type
- `go:generate` directives to come after the `package` declarations
- Format of this changelog to be better aligned with [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)


## [0.3.0] - 2022-04-22

### Added

- The Zero 2 W model
- Badges and library installation instructions to the README

### Changed

- Code formatting using [`gofumpt`](https://github.com/mvdan/gofumpt)
- The location of `tools.go` from the module root to an internal package

### Fixed

- Printers to use the proper Go technique for specifying the label width
- Tests to instantiate `require` assertions where that makes sense
- Miscellaneous minor code and documentation legibility issues


## [0.2.1] - 2022-02-14

### Fixed

- The package name in `tools.go`
- Misspelling in a comment in `cmd/pirev/main.go`


## [0.2.0] - 2022-01-19

### Added

- `tools.go` with imports for:
	- `golang.org/x/tools/cmd/stringer`
	- `honnef.co/go/tools/cmd/staticcheck`
- Space between the number and the unit when reporting memory sizes
- Comment on `const revBase` in `cmd/pirev/main.go`

### Changed

- Go version in `go.mod` to 1.17

### Removed

- Unused `const revBase` from `pirev.go`

### Fixed

- Name of the `print` package in its documentation comment


## [0.1.1] - 2021-05-14

### Fixed

- Documentation
	- Main CLI package path in README installation instructions
	- Capitalization, wording and other minor fixes in `pirev/bits`


## [0.1.0] - 2021-05-14

Initial release


[semver2]: https://semver.org/spec/v2.0.0.html

[0.5.0]: https://github.com/antichris/go-pirev/releases/v0.5.0
[0.4.0]: https://github.com/antichris/go-pirev/releases/v0.4.0
[0.3.0]: https://github.com/antichris/go-pirev/releases/v0.3.0
[0.2.1]: https://github.com/antichris/go-pirev/releases/v0.2.1
[0.2.0]: https://github.com/antichris/go-pirev/releases/v0.2.0
[0.1.1]: https://github.com/antichris/go-pirev/releases/v0.1.1
[0.1.0]: https://github.com/antichris/go-pirev/releases/v0.1.0
