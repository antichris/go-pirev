package print_test

import (
	"strings"
	"testing"

	"github.com/antichris/go-pirev"
	. "github.com/antichris/go-pirev/internal/print"

	"github.com/stretchr/testify/require"
)

func TestPrinter_Print(t *testing.T) {
	for _, tt := range []struct {
		i    pirev.Info
		wOut string
	}{{
		i: pirev.Info{
			Model:        pirev.Internal,
			Revision:     13,
			MemSize:      1 << 10,
			Processor:    pirev.BCM2711,
			Manufacturer: pirev.Stadium,
		},
		wOut: `                          Model: Internal use only
                       Revision: 1.13
                    Memory size: 1 GB
                      Processor: BCM2711
                   Manufacturer: Stadium
        New-style revision code: false
Warranty voided by overclocking: false
         OTP reading disallowed: false
     OTP programming disallowed: false
         Overvoltage disallowed: false
`,
	}, {
		i: pirev.Info{
			MemSize: 1<<10 - 1,
		},
		wOut: `                          Model: A
                       Revision: 1.0
                    Memory size: 1023 MB
                      Processor: BCM2835
                   Manufacturer: Sony UK
        New-style revision code: false
Warranty voided by overclocking: false
         OTP reading disallowed: false
     OTP programming disallowed: false
         Overvoltage disallowed: false
`,
	}} {
		t.Run("", func(t *testing.T) {
			p := NewPrinter()
			out := &strings.Builder{}

			p.Print(out, tt.i)

			require.Equal(t, tt.wOut, out.String())
		})
	}
}
