# Changelog

This project adheres to [Semantic Versioning][semver2].


## 0.2.1

### Fixed

- The package name in `tools.go`
- Misspelling in a comment in `cmd/pirev/main.go`


## 0.2.0

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


## 0.1.1

### Fixed

- Documentation
	- Main CLI package path in README installation instructions
	- Capitalization, wording and other minor fixes in `pirev/bits`


## 0.1.0

Initial release


[semver2]: https://semver.org/spec/v2.0.0.html
