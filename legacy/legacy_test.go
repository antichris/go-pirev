package legacy_test

import (
	"fmt"
	"testing"

	. "github.com/antichris/go-pirev/legacy"

	"github.com/stretchr/testify/require"
)

func TestDescribe(t *testing.T) {
	t.Parallel()
	tests := []struct {
		rev  uint16
		want Info
		wErr error
	}{{
		rev:  0x0666,
		wErr: ErrRevisionOutOfRange,
	}, {
		rev:  0x0001,
		wErr: ErrUnknownRevision,
	}, {
		rev: 0x0015,
		want: Info{
			Model:        APlus,
			Revision:     1.1,
			Memory:       M256,
			Manufacturer: Embest,
		},
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("rev=%04x", tt.rev), func(t *testing.T) {
			t.Parallel()
			got, err := Describe(tt.rev)
			if tt.wErr != nil {
				require.ErrorIs(t, err, tt.wErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
