package pirev_test

import (
	"testing"

	. "github.com/antichris/go-pirev"

	"github.com/stretchr/testify/require"
)

func TestManufacturer_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		v Manufacturer
		w string
	}{{
		w: "Sony UK",
	}, {
		v: Embest2,
		w: "Embest",
	}, {
		v: 6,
		w: "Manufacturer(6)",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.w, func(t *testing.T) {
			t.Parallel()
			got := tt.v.String()
			require.Equal(t, tt.w, got)
		})
	}
}

func TestProcessor_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		v Processor
		w string
	}{{
		w: "BCM2835",
	}, {
		v: BCM2711,
		w: "BCM2711",
	}, {
		v: 4,
		w: "Processor(4)",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.w, func(t *testing.T) {
			t.Parallel()
			got := tt.v.String()
			require.Equal(t, tt.w, got)
		})
	}
}

func TestModel_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		v Model
		w string
	}{{
		w: "A",
	}, {
		v: FourB,
		w: "4B",
	}, {
		v: 21,
		w: "Model(21)",
	}}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.w, func(t *testing.T) {
			t.Parallel()
			got := tt.v.String()
			require.Equal(t, tt.w, got)
		})
	}
}
