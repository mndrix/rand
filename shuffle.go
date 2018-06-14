// +build go1.10

package rand

// Shuffle pseudo-randomizes the order of elements using the default Source.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func Shuffle(n int, swap func(i, j int)) {
	globalSource.Shuffle(n, swap)
}
