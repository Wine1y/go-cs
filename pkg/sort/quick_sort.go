package sort

import (
	"cmp"
)

func QuickSort[T cmp.Ordered](slice []T) {
	storeI, pivotI := 0, len(slice)-1
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] <= slice[pivotI] {
			slice[i], slice[storeI] = slice[storeI], slice[i]
			storeI++
		}
	}

	slice[storeI], slice[pivotI] = slice[pivotI], slice[storeI]

	if storeI > 1 {
		QuickSort[T](slice[:storeI])
	}
	if len(slice)-(storeI+1) > 1 {
		QuickSort[T](slice[storeI+1:])
	}
}
