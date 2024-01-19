package sort

import "cmp"

func BubbleSort[T cmp.Ordered](slice []T) {
	for {
		swapped := false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i+1], slice[i] = slice[i], slice[i+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
