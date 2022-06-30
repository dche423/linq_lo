package linq_lo_test

import (
	"testing"

	ll "linq-lo"
)

func BenchmarkMap(b *testing.B) {
	data := randomIntSlice(sliceLen)
	for i := 0; i < b.N; i++ {
		ll.Map(data)
	}
}

func BenchmarkMapLo(b *testing.B) {
	data := randomIntSlice(sliceLen)
	for i := 0; i < b.N; i++ {
		ll.MapLo(data)
	}
}

func BenchmarkMapLinq(b *testing.B) {
	data := randomIntSlice(sliceLen)
	for i := 0; i < b.N; i++ {
		ll.MapLinq(data)
	}
}
