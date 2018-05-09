package rand // import "github.com/mndrix/rand"

import (
	crypto "crypto/rand"
	"encoding/binary"
	math "math/rand"
	"sync"
)

// global value that performs all random number operations
var rand = math.New(&source{})

// math/rand Source using entropy from crypto/rand
type source struct {
	data [8]byte // 64 bits
	mtx  sync.Mutex
}

// needed by math/rand.Source, but we don't require seeding
func (src *source) Seed(_ int64) {}

// needed by math/rand.Source
func (src *source) Int63() int64 {
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
	x := binary.BigEndian.Uint64(data)
	return int64(x >> 1) // need 63 bit number
}

// Intn is like Intn in math/rand
func Intn(n int) int {
	return rand.Intn(n)
}

// Int63 is like Int63 in math/rand
func Int63() int64 {
	return rand.Int63()
}

// Int63n is like Int63n in math/rand
func Int63n(n int64) int64 {
	return rand.Int63n(n)
}
