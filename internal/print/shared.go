package print

import (
	"fmt"
	"io"

	"github.com/antichris/go-pirev/bits"
)

type printer struct {
	labels []string
}

func (p printer) Print(wr io.Writer, order Fields, s values) (n int, e error) {
	format := fmt.Sprintf("%%%ds: %%v\n", labelWidth(p.labels, order))

	for _, f := range order {
		_n, e := fmt.Fprintf(wr, format, p.labels[f], s[f])
		n += _n
		if e != nil {
			return n, e
		}
	}
	return
}

func labelWidth(labels []string, fields Fields) (n int) {
	for _, f := range fields {
		l := len(labels[f])
		if l > n {
			n = l
		}
	}
	return n
}

// Fields alias bits.Field slice.
type Fields = []bits.Field

type values = []interface{}

// DefaultOrder of fields in output.
var DefaultOrder = Fields{
	bits.Model,
	bits.Revision,
	bits.Memory,
	bits.Processor,
	bits.Manufacturer,
	bits.NewStyleRev,
	bits.WarrantyVoid,
	bits.NoOTPReading,
	bits.NoOTPPrograming,
	bits.NoOvervoltage,
}

// defaultLabels are based on Field names, except where custom values assigned.
var defaultLabels = customizeLabels([]string{
	bits.Memory:          "Memory size",
	bits.NewStyleRev:     "New-style revision code",
	bits.WarrantyVoid:    "Warranty voided by overclocking",
	bits.NoOTPReading:    "OTP reading disallowed",
	bits.NoOTPPrograming: "OTP programming disallowed",
	bits.NoOvervoltage:   "Overvoltage disallowed",
})

func customizeLabels(custom []string) (s [bits.NumTotal]string) {
	l := len(custom)
	for i := range s {
		if i < l && custom[i] != "" {
			s[i] = custom[i]
			continue
		}
		s[i] = bits.Field(i).String()
	}
	return s
}
