package search

import (
	"cmp"
	"math"
)

func JumpSearch[T cmp.Ordered](slice []T, item T) (index int, found bool) {
	step := int(math.Sqrt(float64(len(slice))))

	start := 0
	for {
		next := start + step
		if next >= len(slice) || slice[next] > item {
			break
		}
		start += step
	}

	for i := start; i < start+step && i < len(slice); i++ {
		if slice[i] == item {
			return i, true
		}
	}
	return -1, false
}
