package print

import (
	"io"

	"github.com/antichris/go-pirev/bits"
	"github.com/antichris/go-pirev/legacy"
)

// NewLegacyPrinter instance.
func NewLegacyPrinter() *LegacyPrinter {
	return &LegacyPrinter{
		printer: printer{labels: defaultLabels[:]},
		order:   DefaultLegacyOrder(),
	}
}

// LegacyPrinter pretty-prints info of an old-style revision code.
type LegacyPrinter struct {
	printer
	order Fields
}

// Print info divined from an old-style revision code.
func (p LegacyPrinter) Print(w io.Writer, i legacy.Info) (n int, e error) {
	vs := values{
		bits.Model:        i.Model,
		bits.Revision:     i.Revision,
		bits.Memory:       i.Memory,
		bits.Manufacturer: i.Manufacturer,
		bits.NewStyleRev:  false,
	}
	return p.printer.Print(w, p.order, vs)
}

// DefaultLegacyOrder of fields in output of legacy revision code info.
func DefaultLegacyOrder() Fields {
	return filterOrderForLegacy(DefaultOrder)
}

func filterOrderForLegacy(order Fields) Fields {
	s := make(Fields, len(legacyFields))
	i := 0
	for _, f := range order {
		if _, ok := legacyFields[f]; ok {
			s[i] = f
			i++
		}
	}
	return s
}

var legacyFields = map[bits.Field]emptyT{
	bits.Model:        emptyV,
	bits.Revision:     emptyV,
	bits.Memory:       emptyV,
	bits.Manufacturer: emptyV,
	bits.NewStyleRev:  emptyV,
}

type emptyT = struct{}

var emptyV = emptyT{}
