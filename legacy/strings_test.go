package legacy_test

import (
	"fmt"
	"testing"

	. "github.com/antichris/go-pirev/legacy"

	"github.com/stretchr/testify/require"
)

func TestModel_String(t *testing.T) {
	t.Parallel()
	for v, want := range map[Model]string{
		0:    "A",
		CM1:  "CM1",
		0xff: "Model(255)",
	} {
		testStringer(t, v, want)
	}
}

func TestMemory_String(t *testing.T) {
	t.Parallel()
	for v, want := range map[Memory]string{
		0:    "256 MB",
		M512: "512 MB",
		2:    "Memory(2)",
	} {
		testStringer(t, v, want)
	}
}

func TestManufacturer_String(t *testing.T) {
	t.Parallel()
	for v, want := range map[Manufacturer]string{
		0:      "Egoman",
		SonyUK: "Sony UK",
		4:      "Manufacturer(4)",
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
