package priorityqueue

import "cmp"

/*
Priority queue is a data structure used to find the maximum or minimum element among a
collection of elements.
Ascending priority queues are used to find the minimum element (DeleteMin).
Descending priority queues are used to find the maximum element (DeleteMax).
Priority queues are used for data compression, graph algorithms and sorting.
*/
type PriorityQueue[T cmp.Ordered] interface {
	//Add new item to the priority queue
	Insert(data T)
	//Get priority queue length
	Length() int
}

type AscendingPriorityQueue[T cmp.Ordered] interface {
	PriorityQueue[T]
	//Return minimum item without removing it from the queue
	GetMin() T
	//Remove and return minimum item
	DeleteMin() T
}

type DescendingPriorityQueue[T cmp.Ordered] interface {
	PriorityQueue[T]
	//Return maximum item without removing it from the queue
	GetMax() T
	//Remove and return maximum item
	DeleteMax() T
}

type DoubleEndedPriorityQueue[T cmp.Ordered] interface {
	AscendingPriorityQueue[T]
	DescendingPriorityQueue[T]
}
