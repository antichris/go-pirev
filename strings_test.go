package pirev_test

import (
	"fmt"
	"testing"

	. "github.com/antichris/go-pirev"

	"github.com/stretchr/testify/require"
)

func TestManufacturer_String(t *testing.T) {
	t.Parallel()
	for v, want := range map[Manufacturer]string{
		0:                "Sony UK",
		Embest2:          "Embest",
		^Manufacturer(0): "Manufacturer(255)",
	} {
		testStringer(t, v, want)
	}
}

func TestProcessor_String(t *testing.T) {
	t.Parallel()
	for v, want := range map[Processor]string{
		0:             "BCM2835",
		BCM2711:       "BCM2711",
		^Processor(0): "Processor(255)",
	} {
		testStringer(t, v, want)
	}
}

func TestModel_String(t *testing.T) {
	t.Parallel()
	for v, want := range map[Model]string{
		0:         "A",
		FourB:     "4B",
		^Model(0): "Model(255)",
	} {
		testStringer(t, v, want)
	}
}

func testStringer(t *testing.T, v fmt.Stringer, want string) bool {
	t.Helper()
	return t.Run(want, func(t *testing.T) {
		t.Helper()
		t.Parallel()
		require.Equal(t, want, v.String())
	})
}
