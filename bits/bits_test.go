package bits_test

import (
	"fmt"
	"testing"

	. "github.com/antichris/go-pirev/bits"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		rev  Code
		want Values
	}{{
		rev: rev,
		want: Values{
			NewStyleRev: 1,
			Memory:      4,
			Processor:   BCM2711,
			Model:       FourB,
			Revision:    2,
		},
	}, {
		rev: revFirstBit,
		want: Values{
			First: 1,
			Last:  0,
		},
	}} {
		tt := tt
		t.Run(fmt.Sprintf("rev=%#x", tt.rev), func(t *testing.T) {
			t.Parallel()
			got := Parse(tt.rev)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestAssemble(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		parts Values
		want  Code
		wErr  string
	}{{
		want: 0,
	}, {
		parts: Values{
			Last + 1: 1,
		},
		wErr: ErrOverflow.Error(),
	}, {
		parts: Values{
			Memory: 0xff,
		},
		wErr: "field Memory takes 3 bit value, 0xff is 8 bits",
	}, {
		parts: Values{
			NoOvervoltage: 1,
		},
		want: 0x80000000,
	}, {
		parts: Values{
			NewStyleRev: 1,
			Memory:      4,
			Processor:   BCM2711,
			Model:       FourB,
			Revision:    2,
		},
		want: rev,
	}} {
		tt := tt
		t.Run(fmt.Sprintf("parts=%v", tt.parts), func(t *testing.T) {
			t.Parallel()

			got, err := Assemble(tt.parts)

			req := require.New(t)
			if tt.wErr != "" {
				req.EqualError(err, tt.wErr)
			} else {
				req.NoError(err)
			}
			req.Equal(tt.want, got)
		})
	}
}

const (
	rev         = 0xc03112
	revFirstBit = 0x80000000
	BCM2711     = 3
	FourB       = 0x11
)
