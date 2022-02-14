package print

import (
	"fmt"
	"testing"

	"github.com/antichris/go-pirev/bits"
	"github.com/stretchr/testify/require"
)

func Test_printer_Print(t *testing.T) {
	t.Parallel()

	labels := []string{First.String()}
	p := printer{labels: labels}
	vs := values{First: true}

	tests := []struct {
		name   string
		order  Fields
		outErr string
		wN     int
		wErr   string
		wOut   string
	}{{
		order:  Fields{First},
		outErr: snap,
		wErr:   snap,
	}, {
		order: Fields{First},
		wN:    20,
		wOut:  "NoOvervoltage: true\n",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := &erryStrBuilder{err: tt.outErr}

			gotN, err := p.Print(out, tt.order, vs)

			if tt.wErr != "" {
				require.EqualError(t, err, tt.wErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.wN, gotN)
			require.Equal(t, tt.wOut, out.String())
		})
	}
}

func Test_customizeLabels(t *testing.T) {
	t.Parallel()

	type ss = []string
	const f = First
	fStr := f.String()

	tests := []struct {
		custom ss
		want   string
	}{{
		custom: nil,
		want:   fStr,
	}, {
		custom: ss{f: ""},
		want:   fStr,
	}, {
		custom: ss{f: snap},
		want:   snap,
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("s=%q", tt.custom), func(t *testing.T) {
			t.Parallel()
			got := customizeLabels(tt.custom)[f]
			require.Equal(t, tt.want, got)
		})
	}
}

const First = bits.First
