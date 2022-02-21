package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_erryStrBuilder_Write(t *testing.T) {
	tests := []struct {
		str, err   string
		wStr, wErr string
	}{{
		err:  snap,
		wErr: snap,
	}, {
		str:  snap,
		wStr: snap,
	}}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("str=%s/err=%s", tt.str, tt.err), func(t *testing.T) {
			b := &erryStrBuilder{err: tt.err}

			_, err := b.Write([]byte(tt.str))

			req := require.New(t)
			if tt.wErr != "" {
				req.EqualError(err, tt.wErr)
			} else {
				req.NoError(err)
			}
			req.Equal(tt.wStr, b.String())
		})
	}
}

const snap = "snap"

// erryStrBuilder returns error on Write calls if its err is non-empty.
type erryStrBuilder struct {
	strings.Builder
	err string
}

func (wr *erryStrBuilder) Write(p []byte) (int, error) {
	if wr.err != "" {
		return 0, errors.New(wr.err)
	}
	return wr.Builder.Write(p)
}
