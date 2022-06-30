package linq_lo

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/samber/lo"
)

// Uniq home brew uniq
func Uniq(l []int) []int {
	m := make(map[int]struct{}, len(l))
	result := make([]int, 0, len(l))

	for _, v := range l {
		if _, ok := m[v]; ok {
			continue
		}

		m[v] = struct{}{}
		result = append(result, v)
	}
	return result
}

// UniqLo uniq with lo
func UniqLo(l []int) []int {
	return lo.Uniq(l)
}

// UniqLinq uniq with linq
func UniqLinq(l []int) []int {
	var result []int
	linq.From(l).Distinct().ToSlice(&result)
	return result
}
