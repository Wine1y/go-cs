package sort

import "cmp"

func MergeSort[T cmp.Ordered](slice []T) []T {
	if len(slice) == 1 {
		return slice
	}

	splitI := len(slice) / 2

	left := MergeSort[T](slice[:splitI])
	right := MergeSort[T](slice[splitI:])

	result := make([]T, len(slice))
	i, l, r := 0, 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}

	for l < len(left) {
		result[i] = left[l]
		i++
		l++
	}
	for r < len(right) {
		result[i] = right[r]
		i++
		r++
	}
	return result
}
