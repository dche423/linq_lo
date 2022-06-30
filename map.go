package linq_lo

import (
	"fmt"

	"github.com/ahmetb/go-linq/v3"
	"github.com/samber/lo"
)

func Map(l []int) []string {
	result := make([]string, len(l), len(l))
	for i, v := range l {
		result[i] = fmt.Sprint(v)
	}
	return result
}

func MapLo(l []int) []string {
	fn := func(x, _ int) string {
		return fmt.Sprint(x)
	}
	return lo.Map(l, fn)
}

func MapLinq(l []int) []string {
	fn := func(x int) string {
		return fmt.Sprint(x)
	}
	var result []string
	linq.From(l).SelectT(fn).ToSlice(&result)
	return result
}
