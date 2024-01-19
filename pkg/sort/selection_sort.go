package sort

import "cmp"

func SelectionSort[T cmp.Ordered](slice []T) {
	for unsortedStart := 0; unsortedStart < len(slice)-1; unsortedStart++ {
		minI := unsortedStart
		for i := minI + 1; i < len(slice); i++ {
			if slice[i] < slice[minI] {
				minI = i
			}
		}
		slice[unsortedStart], slice[minI] = slice[minI], slice[unsortedStart]
	}
}
