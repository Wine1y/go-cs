package heap

import "cmp"

type BinaryHeap[T cmp.Ordered] struct {
	data      []T
	isOrdered HeapOrder[T]
}

func NewBinaryHeap[T cmp.Ordered](order HeapOrder[T]) BinaryHeap[T] {
	return BinaryHeap[T]{
		data:      make([]T, 0),
		isOrdered: order,
	}
}

func NewBinaryHeapWithSize[T cmp.Ordered](size int, order HeapOrder[T]) BinaryHeap[T] {
	return BinaryHeap[T]{
		data:      make([]T, 0, size),
		isOrdered: order,
	}
}

func Heapify[T cmp.Ordered](slice []T, order HeapOrder[T]) BinaryHeap[T] {
	heap := NewBinaryHeapWithSize[T](len(slice), order)
	for _, value := range slice {
		heap.Push(value)
	}
	return heap
}

func (heap *BinaryHeap[T]) Push(data T) {
	heap.data = append(heap.data, data)
	heap.siftUp(heap.Length() - 1)
}

func (heap *BinaryHeap[T]) Pop() T {
	if heap.Length() == 0 {
		panic("Can't pop from an empty heap")
	}
	node := heap.Peek()
	heap.swap(0, heap.Length()-1)
	heap.data = heap.data[:heap.Length()-1]
	heap.siftDown(0)
	return node
}

func (heap BinaryHeap[T]) Peek() T {
	return heap.data[0]
}

func (heap *BinaryHeap[T]) Replace(data T) {
	if heap.Length() == 0 {
		panic("Can't replace, heap has no root")
	}
	heap.data[0] = data
	heap.siftDown(0)
}

func (heap BinaryHeap[T]) Length() int {
	return len(heap.data)
}

func (heap *BinaryHeap[T]) siftUp(i int) {
	for parent := (i - 1) / 2; !heap.isOrdered(heap.at(parent), heap.at(i)); {
		heap.swap(parent, i)
		i = parent
		parent = (i - 1) / 2
	}
}

func (heap *BinaryHeap[T]) siftDown(i int) {
	for {
		left, right := i*2+1, i*2+2
		leftOrdered := left >= heap.Length() || heap.isOrdered(heap.at(i), heap.at(left))
		rightOrdered := right >= heap.Length() || heap.isOrdered(heap.at(i), heap.at(right))
		if leftOrdered && rightOrdered {
			break
		}
		var swapWith int
		if right >= heap.Length() || heap.isOrdered(heap.at(left), heap.at(right)) {
			swapWith = left
		} else {
			swapWith = right
		}
		heap.swap(i, swapWith)
		i = swapWith
	}
}

func (heap *BinaryHeap[T]) swap(a, b int) {
	heap.data[a], heap.data[b] = heap.data[b], heap.data[a]
}

func (heap BinaryHeap[T]) at(i int) T {
	return heap.data[i]
}
