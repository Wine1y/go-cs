package sort

import (
	"cmp"

	"github.com/Wine1y/go-cs/pkg/containers/heap"
)

func HeapSort[T cmp.Ordered](slice []T) {
	heap := heap.HeapifyBinaryHeap[T](slice, heap.MinHeapOrder)
	for i := 0; i < len(slice); i++ {
		slice[i] = heap.Pop()
	}
}
