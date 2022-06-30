package linq_lo_test

import (
	"math/rand"
	"testing"

	ll "linq-lo"
)

const sliceLen = 1e5

func randomIntSlice(n int) []int {
	l := make([]int, n, n)
	for i := 0; i < n; i++ {
		l[i] = rand.Intn(n)
	}
	return l
}
func BenchmarkUniq(b *testing.B) {
	data := randomIntSlice(sliceLen)
	for i := 0; i < b.N; i++ {
		ll.Uniq(data)
	}
}

func BenchmarkUniqLinq(b *testing.B) {
	data := randomIntSlice(sliceLen)
	for i := 0; i < b.N; i++ {
		ll.UniqLinq(data)
	}
}

func BenchmarkUniqLo(b *testing.B) {
	data := randomIntSlice(sliceLen)
	for i := 0; i < b.N; i++ {
		ll.UniqLo(data)
	}
}
