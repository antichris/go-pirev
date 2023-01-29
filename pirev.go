// Package pirev parses Raspberry Pi new-style revision codes.
package pirev

//go:generate go run golang.org/x/tools/cmd/stringer -output=strings.go -type=Manufacturer,Processor,Model -linecomment

import "github.com/antichris/go-pirev/bits"

// Identify hardware from a revision code.
func Identify(r Code) Info {
	p := bits.Parse(bits.Code(r))
	return Info{
		Revision:        uint8(p[bits.Revision]),
		Model:           Model(p[bits.Model]),
		Processor:       Processor(p[bits.Processor]),
		Manufacturer:    Manufacturer(p[bits.Manufacturer]),
		MemSize:         256 << p[bits.Memory],
		NewStyleRev:     p[bits.NewStyleRev] != 0,
		WarrantyVoid:    p[bits.WarrantyVoid] != 0,
		NoOTPReading:    p[bits.NoOTPReading] != 0,
		NoOTPPrograming: p[bits.NoOTPPrograming] != 0,
		NoOvervoltage:   p[bits.NoOvervoltage] != 0,
	}
}

// Code is a Raspberry Pi revision code.
type Code uint32

// Info describes Raspberry Pi hardware as identified by revision code.
type Info struct {
	Revision        uint8 // 1.#
	Model           Model
	Processor       Processor
	Manufacturer    Manufacturer
	MemSize         uint // In megabytes
	NewStyleRev     bool // New-style revision code
	WarrantyVoid    bool // Warranty has been voided by overclocking
	NoOTPReading    bool // OTP reading disallowed
	NoOTPPrograming bool // OTP programming disallowed
	NoOvervoltage   bool // Overvoltage disallowed
}

type (
	// Model is an 8-bit value.
	Model uint8
	// Processor is a 4-bit value.
	Processor uint8
	// Manufacturer is a 4-bit value.
	Manufacturer uint8
)

// Model constants.
const (
	A Model = iota
	B
	APlus // A+
	BPlus // B+
	TwoB  // 2B
	Alpha // Alpha (early prototype)
	CM1
	_Unspecified0
	ThreeB // 3B
	Zero
	CM3
	_Unspecified1
	ZeroW       // Zero W
	ThreeBPlus  // 3B+
	ThreeAPlus  // 3A+
	Internal    // Internal use only
	CM3Plus     // CM3+
	FourB       // 4B
	Zero2W      // Zero 2 W
	FourHundred // 400
	CM4
	CM4S
)

// Processor type constants.
const (
	BCM2835 Processor = iota
	BCM2836
	BCM2837
	BCM2711
)

// Manufacturer constants.
const (
	SonyUK Manufacturer = iota // Sony UK
	Egoman
	Embest
	SonyJapan // Sony Japan
	// Embest2 is the same as Embest, if the docs can be believed.
	Embest2 // Embest
	Stadium
)
