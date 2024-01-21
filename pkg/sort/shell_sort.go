package sort

import (
	"cmp"
)

var gaps []int = []int{701, 301, 132, 57, 23, 10, 4, 1}

func ShellSort[T cmp.Ordered](slice []T) {
	for _, gap := range gaps {
		for i := gap; i < len(slice); i++ {
			for j := i - gap; j >= 0; j -= gap {
				if slice[j] > slice[j+gap] {
					slice[j+gap], slice[j] = slice[j], slice[j+gap]
				}
			}
		}
	}
}
