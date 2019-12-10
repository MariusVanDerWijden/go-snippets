package search

// Search calls the provided function with increasing input
// from [0,n) and returns the input i if the provided function
// returns true, can be used to find the first occurence
// in an array that satisfies a given predicate.
func Search(n int, f func(i int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}
	return -1
}
