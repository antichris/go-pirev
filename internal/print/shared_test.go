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

	for _, tt := range []struct {
		name   string
		outErr string
		wErr   string
		wOut   string
		order  Fields
		wN     int
	}{{
		order:  Fields{First},
		outErr: snap,
		wErr:   snap,
	}, {
		order: Fields{First},
		wN:    20,
		wOut:  "NoOvervoltage: true\n",
	}} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := &erryStrBuilder{err: tt.outErr}

			gotN, err := p.Print(out, tt.order, vs)

			req := require.New(t)
			if tt.wErr != "" {
				req.EqualError(err, tt.wErr)
			} else {
				req.NoError(err)
			}
			req.Equal(tt.wN, gotN)
			req.Equal(tt.wOut, out.String())
		})
	}
}

func Test_customizeLabels(t *testing.T) {
	t.Parallel()

	type ss = []string
	const f = First
	fStr := f.String()

	for _, tt := range []struct {
		want   string
		custom ss
	}{{
		custom: nil,
		want:   fStr,
	}, {
		custom: ss{f: ""},
		want:   fStr,
	}, {
		custom: ss{f: snap},
		want:   snap,
	}} {
		tt := tt
		t.Run(fmt.Sprintf("s=%q", tt.custom), func(t *testing.T) {
			t.Parallel()
			got := customizeLabels(tt.custom)[f]
			require.Equal(t, tt.want, got)
		})
	}
}

const First = bits.First
