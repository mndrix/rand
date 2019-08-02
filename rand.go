package rand // import "github.com/mndrix/rand"

import (
	crypto "crypto/rand"
	"encoding/binary"
	math "math/rand"
	"sync"
)

// global value that performs all random number operations
var globalSource = math.New(&source{})

// math/rand Source using entropy from crypto/rand
type source struct {
	data [8]byte // 64 bits
	mtx  sync.Mutex
}

// Float64 returns a pseudo-random number in [0.0,1.0)
func Float64() float64 { return globalSource.Float64() }

// needed by math/rand.Source, but we don't require seeding
func (src *source) Seed(_ int64) {}

// needed by math/rand.Source
func (src *source) Int63() int64 {
	return int64(src.Uint64() >> 1) // need 63 bit number
}

// Uint64 returns a randomly selected 64-bit unsigned integer.
func (src *source) Uint64() uint64 {
	src.mtx.Lock()
	defer src.mtx.Unlock()

	data := src.data[:]
	n, err := crypto.Read(data)
	if err != nil {
		panic("crypto.Read failed: " + err.Error())
	}
	if n != 8 {
		panic("read too few random bytes")
	}
	return binary.BigEndian.Uint64(data)
}

// Intn returns, as an int, a non-negative pseudo-random number in
// [0,n). It panics if n <= 0.
func Intn(n int) int {
	return globalSource.Intn(n)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an
// int64.
func Int63() int64 {
	return globalSource.Int63()
}

// Int63n returns, as an int64, a non-negative pseudo-random number in
// [0,n). It panics if n <= 0.
func Int63n(n int64) int64 {
	return globalSource.Int63n(n)
}

// Uint64 returns a randomly selected 64-bit unsigned integer.
func Uint64() uint64 {
	return globalSource.Uint64()
}
