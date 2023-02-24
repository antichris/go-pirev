package bits_test

import (
	"testing"

	. "github.com/antichris/go-pirev/bits"

	"github.com/stretchr/testify/require"
)

func TestID_String(t *testing.T) {
	t.Parallel()
	for v, want := range map[Field]string{
		0:      "NoOvervoltage",
		Memory: "Memory",
		13:     "Field(13)",
	} {
		v, want := v, want
		t.Run(want, func(t *testing.T) {
			t.Helper()
			t.Parallel()
			require.Equal(t, want, v.String())
		})
	}
}
