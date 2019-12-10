package search

func BenchmarkSearch(b *testing.B) {
	for i := 1; i < b.N; i++ {
		arr := make([]byte, i)
		arr[i-1] = 1
		idx := Search(i, func(i int) bool {
			return arr[i] != 0
		})
		assert.Equal(b, i-1, idx, "")
	}
}

func BenchmarkForLoop(b *testing.B) {
	for i := 1; i < b.N; i++ {
		arr := make([]byte, i)
		arr[i-1] = 1
		var idx int
		for k, ele := range arr {
			if ele != 0 {
				idx = k
			}
		}
		assert.Equal(b, i-1, idx, "")
	}
}
