package search

import "cmp"

type SearchingFunction[T comparable] func(slice []T, item T) (index int, found bool)
type OrderedSearchingFunction[T cmp.Ordered] func(slice []T, item T) (index int, found bool)
