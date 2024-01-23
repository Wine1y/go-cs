package search

import "cmp"

func BinarySearch[T cmp.Ordered](slice []T, item T) (index int, found bool) {
	for low, high := 0, len(slice)-1; low <= high; {
		mid := (low + high) / 2
		if slice[mid] == item {
			return mid, true
		}
		if slice[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, false
}
