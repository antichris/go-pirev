package legacy_test

import (
	"testing"

	. "github.com/antichris/go-pirev/legacy"

	"github.com/stretchr/testify/require"
)

func TestModel_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		v Model
		w string
	}{{
		w: "A",
	}, {
		v: CM1,
		w: "CM1",
	}, {
		v: 0xff,
		w: "Model(255)",
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

func TestMemory_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		v Memory
		w string
	}{{
		w: "256 MB",
	}, {
		v: M512,
		w: "512 MB",
	}, {
		v: 2,
		w: "Memory(2)",
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

func TestManufacturer_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		v Manufacturer
		w string
	}{{
		w: "Egoman",
	}, {
		v: SonyUK,
		w: "Sony UK",
	}, {
		v: 4,
		w: "Manufacturer(4)",
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
