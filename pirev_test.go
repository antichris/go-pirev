package pirev_test

import (
	"fmt"
	"testing"

	. "github.com/antichris/go-pirev"

	"github.com/stretchr/testify/require"
)

func TestIdentify(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		rev  Code
		want Info
	}{{
		rev: rev,
		want: Info{
			NewStyleRev:     true,
			MemSize:         4 << 10, // 4Ki
			Processor:       BCM2711,
			Manufacturer:    Embest2,
			Model:           FourB,
			Revision:        2,
			WarrantyVoid:    true,
			NoOTPReading:    true,
			NoOTPPrograming: true,
			NoOvervoltage:   true,
		},
	}} {
		tt := tt
		t.Run(fmt.Sprintf("rev=%x", tt.rev), func(t *testing.T) {
			t.Parallel()
			got := Identify(tt.rev)
			require.Equal(t, tt.want, got)
		})
	}
}

const rev = 0xe7c43112
