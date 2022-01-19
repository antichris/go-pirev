package print_test

import (
	"strings"
	"testing"

	. "github.com/antichris/go-pirev/internal/print"
	"github.com/antichris/go-pirev/legacy"

	"github.com/stretchr/testify/require"
)

func TestLegacyPrinter_Print(t *testing.T) {
	i := legacy.Info{
		Model:        legacy.CM1,
		Revision:     5.5,
		Memory:       legacy.M512,
		Manufacturer: legacy.Qisda,
	}
	wantOut := `                  Model: CM1
               Revision: 5.5
            Memory size: 512 MB
           Manufacturer: Qisda
New-style revision code: false
`

	p := NewLegacyPrinter()
	out := &strings.Builder{}

	p.Print(out, i)

	require.Equal(t, wantOut, out.String())
}
