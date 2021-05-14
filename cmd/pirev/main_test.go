package main

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/antichris/go-pirev/legacy"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	var exitCode int

	defer func(a []string, x func(int), o, e io.Writer) {
		args, exit, wrOut, wrErr = a, x, o, e
	}(args, exit, wrOut, wrErr)
	exit = func(code int) { exitCode = code }
	wrOut = &strings.Builder{}
	wrErr = &strings.Builder{}

	args = []string{"foo", "bar"}
	var (
		wantExit = ExitParseErr
		wantOut  = "Usage:  foo"
		wantErr  = `parsing "bar": invalid syntax`
	)
	require.NotPanics(t, func() {
		main()
	})
	require.Equal(t, wantExit, exitCode)
	require.Contains(t, wrOut.(fmt.Stringer).String(), wantOut)
	require.Contains(t, wrErr.(fmt.Stringer).String(), wantErr)
}

func Test_trueMain(t *testing.T) {
	t.Parallel()

	type ss = []string

	tests := []struct {
		args   ss
		outErr string
		errErr string
		wExit  int
		wOut   string
		wErr   string
		wPanic string
	}{{
		wPanic: "runtime error: index out of range [0] with length 0",
	}, {
		args:   ss{"foo"},
		outErr: snap,
		wPanic: snap,
	}, {
		args:  ss{"bar", "--help"},
		wExit: ExitOK,
		wOut:  "Usage:  bar",
	}, {
		args:   ss{"", "qux"},
		errErr: snap,
		wPanic: snap,
	}, {
		args:  ss{"", "corge"},
		wExit: ExitParseErr,
		wErr:  `parsing "corge": invalid syntax`,
	}, {
		args:   ss{"", "800000"},
		outErr: snap,
		wPanic: snap,
	}, {
		args:  ss{"", "0"},
		wExit: ExitLookupErr,
		wErr:  legacy.ErrUnknownRevision.Error(),
	}, {
		args:   ss{"", "15"},
		outErr: snap,
		wPanic: snap,
	}, {
		args:  ss{"", "800000"},
		wExit: ExitOK,
		wOut:  "New-style revision code: true",
	}, {
		args:  ss{"", "15"},
		wExit: ExitOK,
		wOut:  "New-style revision code: false",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf(
			"arg=%q/wPanic=%s", tt.args, tt.wPanic,
		), func(t *testing.T) {
			t.Parallel()
			wrOut := &erryStrBuilder{err: tt.outErr}
			wrErr := &erryStrBuilder{err: tt.errErr}

			var exitCode int
			run := func() {
				exitCode = trueMain(wrOut, wrErr, tt.args)
			}

			if tt.wPanic != "" {
				require.PanicsWithError(t, tt.wPanic, run)
				return
			}
			require.NotPanics(t, run)
			require.Equal(t, tt.wExit, exitCode)
			require.Contains(t, wrOut.String(), tt.wOut)
			require.Contains(t, wrErr.String(), tt.wErr)
		})
	}
}
