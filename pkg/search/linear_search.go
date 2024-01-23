package search

import "cmp"

func LinearSearch[T comparable](slice []T, item T) (index int, found bool) {
	for index = 0; index < len(slice); index++ {
		if slice[index] == item {
			return index, true
		}
	}
	return -1, false
}

func OrderedLinearSearch[T cmp.Ordered](slice []T, item T) (index int, found bool) {
	for index = 0; index < len(slice); index++ {
		if slice[index] == item {
			return index, true
		}
		if slice[index] > item {
			break
		}
	}
	return -1, false
}
