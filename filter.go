package linq_lo

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/samber/lo"
)

func Filter(l []int) []int {
	fn := func(x int) bool {
		return x%13 == 0
	}

	var result []int
	for _, v := range l {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterLo(l []int) []int {
	fn := func(x, _ int) bool {
		return x%13 == 0
	}
	return lo.Filter(l, fn)
}

func FilterLinq(l []int) []int {
	fn := func(x int) bool {
		return x%13 == 0
	}
	var result []int
	linq.From(l).WhereT(fn).ToSlice(&result)
	return result
}
