package rand

import (
	"testing"
)

// unseeded is the first value produced by math/rand.Int63() if its
// internal random number generated has not been seeded.
const unseeded = 5577006791947779410

func TestInts(t *testing.T) {
	seen := make(map[int64]bool)
	ok := func(x int64) {
		if len(seen) < 30 { // output some random numbers for user inspection
			t.Logf("%d", x)
		}
		if x == unseeded {
			t.Errorf("using unseeded, unsafe random number generator")
		}
		if x < 0 {
			t.Errorf("generated a negative number")
		}
		if seen[x] {
			t.Errorf("highly improbable collision. you probably broke something")
		}
		seen[x] = true
	}

	for i := 0; i < 1000; i++ {
		x := Int63()
		ok(x)
	}
}
