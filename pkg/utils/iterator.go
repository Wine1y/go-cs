package utils

type Iterator[T any] interface {
	Next() T
	Ended() bool
}

func CollectIterator[T any](iterator Iterator[T]) []T {
	slice := make([]T, 0)
	for !iterator.Ended() {
		slice = append(slice, iterator.Next())
	}
	return slice
}
