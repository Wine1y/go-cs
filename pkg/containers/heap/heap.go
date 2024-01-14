package heap

import "cmp"

/*
Heap is a tree-like data structure. Heap must satisfy the heap property.
In a max-heap for any given node, the node should be less than or equal to any of it's child nodes.
In a min-heap for any given node, the node should be greater than or equal to any of it's child nodes.
Heaps are used for sorting, priority queue implementations or graph algorithms.
*/
type Heap[T cmp.Ordered] interface {
	//Remove and return maximum item of a max-heap or minimum item of a min-heap
	Pop() T
	//Get maximum item of a max-heap or minimum item of a min-heap without removing it
	Peek() T
	//Add new key to the heap
	Push(data T)
	//Pop root and push a new key
	Replace(data T)
	//Get heap length
	Length() int
}

type HeapOrder[T cmp.Ordered] func(parent, child T) bool

func MaxHeapOrder[T cmp.Ordered](parent, child T) bool {
	return parent >= child
}

func MinHeapOrder[T cmp.Ordered](parent, child T) bool {
	return parent <= child
}
