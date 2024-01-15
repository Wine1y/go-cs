package priorityqueue

import (
	"cmp"

	"github.com/Wine1y/go-cs/pkg/containers/heap"
)

type heapPriorityQueue[T cmp.Ordered] struct {
	heap.BinaryHeap[T]
}

func newHeapPriorityQueue[T cmp.Ordered](heapOrder heap.HeapOrder[T]) heapPriorityQueue[T] {
	return heapPriorityQueue[T]{
		BinaryHeap: heap.NewBinaryHeap[T](heapOrder),
	}
}

func (queue *heapPriorityQueue[T]) Insert(data T) {
	queue.BinaryHeap.Push(data)
}

func (queue heapPriorityQueue[T]) Length() int {
	return queue.BinaryHeap.Length()
}

type HeapAscendingPriorityQueue[T cmp.Ordered] struct {
	heapPriorityQueue[T]
}

func NewHeapAscendingPriorityQueue[T cmp.Ordered]() HeapAscendingPriorityQueue[T] {
	return HeapAscendingPriorityQueue[T]{
		heapPriorityQueue: newHeapPriorityQueue[T](heap.MinHeapOrder),
	}
}

func (queue HeapAscendingPriorityQueue[T]) GetMin() T {
	return queue.BinaryHeap.Peek()
}

func (queue *HeapAscendingPriorityQueue[T]) DeleteMin() T {
	return queue.BinaryHeap.Pop()
}

type HeapDescendingPriorityQueue[T cmp.Ordered] struct {
	heapPriorityQueue[T]
}

func NewHeapDescendingPriorityQueue[T cmp.Ordered]() HeapDescendingPriorityQueue[T] {
	return HeapDescendingPriorityQueue[T]{
		heapPriorityQueue: newHeapPriorityQueue[T](heap.MaxHeapOrder),
	}
}

func (queue HeapDescendingPriorityQueue[T]) GetMax() T {
	return queue.BinaryHeap.Peek()
}

func (queue *HeapDescendingPriorityQueue[T]) DeleteMax() T {
	return queue.BinaryHeap.Pop()
}
