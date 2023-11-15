package ulidpb

import (
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestRoundTrip(t *testing.T) {
	x := ulid.Make()

	y := FromULID(x).AsULID()

	if x != y {
		t.Errorf("round trip failed: %s != %s", x, y)
	}
}
