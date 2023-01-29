// Package bits does the low (bit-)level parsing and assembly of
// new-style Raspberry Pi revision codes.
package bits

//go:generate go run golang.org/x/tools/cmd/stringer -type=Field

import (
	"errors"
	"fmt"
	"math/bits"
)

// Parse a Raspberry Pi revision code into a slice of values.
func Parse(r Code) Values {
	s := make(Values, NumTotal)
	// HACK Seems crazy, but that's how integer overflow works.
	for f := Last; f <= Last; f-- {
		s[f] = Value(r) & f.Bitmask()
		r >>= f.Bitsize()
	}
	return s
}

// Code is a Raspberry Pi revision code.
type Code uint32

// Values is an alias for a Value slice.
type Values = []Value

// The Value of a Field in a revision code.
type Value uint8

// Field constants.
const (
	NoOvervoltage   Field = iota // Overvoltage disallowed
	NoOTPPrograming              // OTP programming disallowed
	NoOTPReading                 // OTP reading disallowed
	Unused1                      //
	WarrantyVoid                 // Warranty has been voided by overclocking
	Unused2                      //
	NewStyleRev                  // New-style revision code
	Memory                       // Memory size in multiples of 256
	Manufacturer
	Processor
	Model
	Revision

	// NumTotal of known Fields.
	NumTotal // Taking advantage of iota continuing its effect here.
	// Last known field.
	Last = NumTotal - 1 // This automatically tracks the last field.
	// First of known Fields, short for "Field(0)" to iterate from.
	First Field = 0
)

// A Field of revision code.
//
// Keep in mind that, as an unsigned integer, Field cannot have negative
// values, and `f := Field(0) - 1` overflows to `f == Field(MaxUint)`.
// This means that you have to compare `f <= Last` if/when iterating
// down from Last.
type Field uint

// Bitmask is an all-ones mask derived from the bitsize of the Field.
func (f Field) Bitmask() Value { return bitmasks[f] }

// Bitsize is the length of Field value in number of bits.
func (f Field) Bitsize() uint8 { return bitsizes[f] }

// Validate returns error if v is not a valid Value for the Field.
func (f Field) Validate(v Value) (err error) {
	if v&^f.Bitmask() != 0 {
		err = fmt.Errorf(
			"field %s takes %d bit value, %#x is %d bits",
			f, f.Bitsize(), v, bits.Len8(uint8(v)),
		)
	}
	return
}

// Naming ideas: assemble/compile/construct/synthesize

// Assemble a slice of Values into a revision code.
func Assemble(s Values) (r Code, err error) {
	n := Field(len(s))
	if n > NumTotal {
		err = ErrOverflow
		return
	}
	for f := First; f < n; f++ {
		v := s[f]
		if err = f.Validate(v); err != nil {
			return
		}
		r = r<<f.Bitsize() | Code(v)
	}
	for f := n; f <= Last; f++ {
		r <<= f.Bitsize()
	}
	return
}

// ErrOverflow is returned when a function gets passed more Values than
// NumTotal.
var ErrOverflow = errors.New("more Values than NumTotal")

// bitmasks of Fields.
var bitmasks = func(s [NumTotal]uint8) (m [NumTotal]Value) {
	// TODO Verify that the array argument is not copied here.
	for i, l := range s {
		m[i] = 1<<l - 1
	}
	return
}(bitsizes)

// bitsizes of fields.
var bitsizes = [NumTotal]uint8{
	NoOvervoltage:   1,
	NoOTPPrograming: 1,
	NoOTPReading:    1,
	Unused1:         3,
	WarrantyVoid:    1,
	Unused2:         1,
	NewStyleRev:     1,
	Memory:          3,
	Manufacturer:    4,
	Processor:       4,
	Model:           8,
	Revision:        4,
}

// CodeBitsize is the number of bits in a Raspberry Pi revision code.
const CodeBitsize = 32

// "index out of range" means bitsizes not adding up to CodeBitsize.
var _ = [1]struct{}{}[0+
	CodeBitsize-
	bitsizes[NoOvervoltage]-
	bitsizes[NoOTPPrograming]-
	bitsizes[NoOTPReading]-
	bitsizes[Unused1]-
	bitsizes[WarrantyVoid]-
	bitsizes[Unused2]-
	bitsizes[NewStyleRev]-
	bitsizes[Memory]-
	bitsizes[Manufacturer]-
	bitsizes[Processor]-
	bitsizes[Model]-
	bitsizes[Revision]-
	0]
