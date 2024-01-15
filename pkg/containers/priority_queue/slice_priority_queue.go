package priorityqueue

import "cmp"

type SlicePriorityQueue[T cmp.Ordered] []T

func NewSlicePriorityQueue[T cmp.Ordered]() SlicePriorityQueue[T] {
	return make(SlicePriorityQueue[T], 0)
}

func (queue *SlicePriorityQueue[T]) Insert(data T) {
	*queue = append(*queue, data)
}

func (queue SlicePriorityQueue[T]) Length() int {
	return len(queue)
}

func (queue SlicePriorityQueue[T]) GetMin() T {
	return queue[queue.getMinIndex()]
}

func (queue *SlicePriorityQueue[T]) DeleteMin() T {
	minIndex := queue.getMinIndex()
	item := (*queue)[minIndex]
	queue.popIndex(minIndex)
	return item
}

func (queue SlicePriorityQueue[T]) GetMax() T {
	return queue[queue.getMaxIndex()]
}

func (queue *SlicePriorityQueue[T]) DeleteMax() T {
	maxIndex := queue.getMaxIndex()
	item := (*queue)[maxIndex]
	queue.popIndex(maxIndex)
	return item
}

func (queue SlicePriorityQueue[T]) getMinIndex() int {
	if queue.Length() == 0 {
		panic("Can't get minimum item, priority queue is empty")
	}
	minI, min := 0, queue[0]
	for i, item := range queue {
		if item < min {
			min = item
			minI = i
		}
	}
	return minI
}

func (queue SlicePriorityQueue[T]) getMaxIndex() int {
	if queue.Length() == 0 {
		panic("Can't get maximum item, priority queue is empty")
	}
	maxI, max := 0, queue[0]
	for i, item := range queue {
		if item > max {
			max = item
			maxI = i
		}
	}
	return maxI
}

func (queue *SlicePriorityQueue[T]) popIndex(i int) {
	switch i {
	case queue.Length() - 1:
		*queue = (*queue)[:queue.Length()-1]
	case 0:
		if queue.Length() > 1 {
			*queue = (*queue)[1:]
		} else {
			*queue = make(SlicePriorityQueue[T], 0)
		}
	default:
		new := make(SlicePriorityQueue[T], queue.Length()-1)
		copy(new, (*queue)[:i])
		copy(new[i:], (*queue)[i+1:])
		*queue = new
	}
}

type SortedSlicePriorityQueue[T cmp.Ordered] []T

func NewSortedSlicePriorityQueue[T cmp.Ordered]() SortedSlicePriorityQueue[T] {
	return make(SortedSlicePriorityQueue[T], 0)
}

func (queue *SortedSlicePriorityQueue[T]) Insert(data T) {
	i := 0
	//Binary search could be used to find insert index in O(logN) instead of O(N)
	for i < queue.Length() && data > (*queue)[i] {
		i++
	}
	switch i {
	case queue.Length() - 1:
		*queue = append(*queue, data)
	case 0:
		new := make(SortedSlicePriorityQueue[T], queue.Length()+1)
		new[0] = data
		copy(new[1:], *queue)
		*queue = new
	default:
		new := make(SortedSlicePriorityQueue[T], queue.Length()+1)
		copy(new, (*queue)[:i])
		new[i] = data
		copy(new[i+1:], (*queue)[i:])
		*queue = new
	}
}

func (queue SortedSlicePriorityQueue[T]) Length() int {
	return len(queue)
}

func (queue SortedSlicePriorityQueue[T]) GetMin() T {
	return queue[queue.getMinIndex()]
}

func (queue *SortedSlicePriorityQueue[T]) DeleteMin() T {
	minIndex := queue.getMinIndex()
	item := (*queue)[minIndex]
	queue.popIndex(minIndex)
	return item
}

func (queue SortedSlicePriorityQueue[T]) GetMax() T {
	return queue[queue.getMaxIndex()]
}

func (queue *SortedSlicePriorityQueue[T]) DeleteMax() T {
	maxIndex := queue.getMaxIndex()
	item := (*queue)[maxIndex]
	queue.popIndex(maxIndex)
	return item
}

func (queue SortedSlicePriorityQueue[T]) getMinIndex() int {
	if queue.Length() == 0 {
		panic("Can't get minimum item, priority queue is empty")
	}
	return 0
}

func (queue SortedSlicePriorityQueue[T]) getMaxIndex() int {
	if queue.Length() == 0 {
		panic("Can't get minimum item, priority queue is empty")
	}
	return queue.Length() - 1
}

func (queue *SortedSlicePriorityQueue[T]) popIndex(i int) {
	switch i {
	case queue.Length() - 1:
		*queue = (*queue)[:queue.Length()-1]
	case 0:
		if queue.Length() > 1 {
			*queue = (*queue)[1:]
		} else {
			*queue = make(SortedSlicePriorityQueue[T], 0)
		}
	default:
		new := make(SortedSlicePriorityQueue[T], queue.Length()-1)
		copy(new, (*queue)[:i])
		copy(new[i:], (*queue)[i+1:])
		*queue = new
	}
}
