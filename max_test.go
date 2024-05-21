package ccmax_test

import (
	"testing"

	ccmax "github.com/Friedfisch/cc-max"
)

func run(t *testing.T, n1, n2, e float64) {
	r := ccmax.Max(n1, n2)
	if r != e {
		t.Errorf("Error (%f, %f): %f did not match %f", n1, n2, r, e)
	}
	r2 := ccmax.MaxCmp(n1, n2)
	if r != r2 {
		t.Errorf("Error (%f, %f): %f did not match %f", n1, n2, r, r2)
	}
	r3 := ccmax.MaxCompact(n1, n2)
	if r != r3 {
		t.Errorf("Error (%f, %f): %f did not match %f", n1, n2, r, r3)
	}
}

func TestMax(t *testing.T) {
	run(t, 3, 3, 3)
	run(t, -3, -3, -3)
	run(t, 0, 0, 0)

	run(t, 3, 0, 3)
	run(t, 4, 0, 4)
	run(t, -3, 0, 0)

	run(t, 0, 3, 3)
	run(t, 0, -3, 0)

	// does not work
	/*run(t, math.Inf(1), math.Inf(1), math.Inf(1))
	run(t, math.Inf(-1), math.Inf(-1), math.Inf(-1))
	run(t, math.Inf(1), 0, math.Inf(1))
	run(t, math.Inf(-1), 0, 0)
	run(t, 0, math.Inf(1), math.Inf(1))
	run(t, 0, math.Inf(-1), 0)
	*/
}
