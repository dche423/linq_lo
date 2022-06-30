package linq_lo_test

import (
	"testing"

	ll "linq-lo"
)

func BenchmarkFilter(b *testing.B) {
	data := randomIntSlice(sliceLen)
	for i := 0; i < b.N; i++ {
		ll.Filter(data)
	}
}

func BenchmarkFilterLo(b *testing.B) {
	data := randomIntSlice(sliceLen)
	for i := 0; i < b.N; i++ {
		ll.FilterLo(data)
	}
}

func BenchmarkFilterLinq(b *testing.B) {
	data := randomIntSlice(sliceLen)
	for i := 0; i < b.N; i++ {
		ll.FilterLinq(data)
	}
}
