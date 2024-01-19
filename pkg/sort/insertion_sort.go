package sort

import "cmp"

func InsertionSort[T cmp.Ordered](slice []T) {
	for unsortedStart := 1; unsortedStart < len(slice); unsortedStart++ {
		for i := unsortedStart; i > 0 && slice[i-1] > slice[i]; i-- {
			slice[i-1], slice[i] = slice[i], slice[i-1]
		}
	}
}
