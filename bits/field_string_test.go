package bits_test

import (
	"testing"

	. "github.com/antichris/go-pirev/bits"

	"github.com/stretchr/testify/require"
)

func TestID_String(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		f Field
		w string
	}{{
		w: "NoOvervoltage",
	}, {
		f: Memory,
		w: "Memory",
	}, {
		f: 13,
		w: "Field(13)",
	}} {
		tt := tt
		t.Run(tt.f.String(), func(t *testing.T) {
			t.Parallel()
			got := tt.f.String()
			require.Equal(t, tt.w, got)
		})
	}
}
