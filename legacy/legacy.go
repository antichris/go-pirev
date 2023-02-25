// Package legacy describes Raspberry Pi old-style revision codes.
package legacy

//go:generate go run golang.org/x/tools/cmd/stringer -output=strings.go -type=Manufacturer,Memory -linecomment

import (
	"errors"

	"github.com/antichris/go-pirev"
)

// Describe an old-style revision code.
func Describe(r RevCode) (i Info, err error) {
	if r > Last {
		err = ErrRevisionOutOfRange
		return
	}
	i = known[r]
	if i.Revision == 0.0 {
		// Since RevInfo.Revision defaults to 0.0, we can tell when it
		// is not in the known table.
		err = ErrUnknownRevision
	}
	return
}

// RevCode is a Raspberry Pi revision code. Also, an alias for uint16.
type RevCode = uint16

// Info about hardware with an old-style revision code.
type Info struct {
	Revision     Revision
	Memory       Memory
	Model        Model
	Manufacturer Manufacturer
}

// Last is the last old-style revision code according to the official
// Raspberry Pi documentation.
const Last RevCode = 0x0015

var (
	// ErrRevisionOutOfRange happens when asked to Describe a revision
	// number greater than the Last known.
	ErrRevisionOutOfRange = errors.New("revision out of range")
	// ErrUnknownRevision happens when asked to Describe a revision number
	// that is not specified in the Raspberry Pi documentation.
	ErrUnknownRevision = errors.New("unknown revision")
)

type (
	// Model code.
	Model = pirev.Model
	// Revision number.
	Revision float32
	// Memory size.
	Memory uint8
	// Manufacturer code.
	Manufacturer uint8
)

// Legacy model constants.
const (
	A     = Model(pirev.A)
	B     = Model(pirev.B)
	APlus = Model(pirev.APlus) // A+
	BPlus = Model(pirev.BPlus) // B+
	CM1   = Model(pirev.CM1)
)

// Memory size constants.
const (
	M256 Memory = iota // 256 MB
	M512               // 512 MB
)

// Legacy manufacturer constants.
const (
	Egoman Manufacturer = iota
	SonyUK              // Sony UK
	Qisda
	Embest
)

var known = [Last + 1]Info{
	// 0x0000: {},
	// 0x0001: {},
	0x0002: {1.0, M256, B, Egoman},
	0x0003: {1.0, M256, B, Egoman},
	0x0004: {2.0, M256, B, SonyUK},
	0x0005: {2.0, M256, B, Qisda},
	0x0006: {2.0, M256, B, Egoman},
	0x0007: {2.0, M256, A, Egoman},
	0x0008: {2.0, M256, A, SonyUK},
	0x0009: {2.0, M256, A, Qisda},
	// 0x000a: {},
	// 0x000b: {},
	// 0x000c: {},
	0x000d: {2.0, M512, B, Egoman},
	0x000e: {2.0, M512, B, SonyUK},
	0x000f: {2.0, M512, B, Egoman},
	0x0010: {1.2, M512, BPlus, SonyUK},
	0x0011: {1.0, M512, CM1, SonyUK},
	0x0012: {1.1, M256, APlus, SonyUK},
	0x0013: {1.2, M512, BPlus, Embest},
	0x0014: {1.0, M512, CM1, Embest},
	0x0015: {1.1, M256, APlus, Embest}, // or M512: docs are ambiguous.
}
