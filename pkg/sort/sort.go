package sort

import "cmp"

type SortingFunction[T cmp.Ordered] func(slice []T) []T
type InPlaceSortingFunction[T cmp.Ordered] func(slice []T)
