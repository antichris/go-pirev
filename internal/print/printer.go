package print

import (
	"fmt"
	"io"

	"github.com/antichris/go-pirev"
	"github.com/antichris/go-pirev/bits"
)

// NewPrinter instance.
func NewPrinter() *Printer {
	return &Printer{
		printer: printer{labels: defaultLabels[:]},
		order:   DefaultOrder,
	}
}

// Printer of Raspberry Pi revision info.
type Printer struct {
	printer
	order Fields
}

// Print info extracted from a new-style revision code.
func (p Printer) Print(w io.Writer, i pirev.Info) (n int, e error) {
	vs := values{
		bits.Model:           i.Model,
		bits.Revision:        fmt.Sprintf("1.%d", i.Revision),
		bits.Memory:          formatMemsize(i.MemSize),
		bits.Processor:       i.Processor,
		bits.Manufacturer:    i.Manufacturer,
		bits.NewStyleRev:     i.NewStyleRev,
		bits.WarrantyVoid:    i.WarrantyVoid,
		bits.NoOTPReading:    i.NoOTPReading,
		bits.NoOTPPrograming: i.NoOTPPrograming,
		bits.NoOvervoltage:   i.NoOvervoltage,
	}
	return p.printer.Print(w, p.order, vs)
}

func formatMemsize(v uint) string {
	const k = 1 << 10 // 1024
	if v < k {
		return fmt.Sprintf("%vMB", v)
	}
	return fmt.Sprintf("%vGB", float32(v)/k)
}
