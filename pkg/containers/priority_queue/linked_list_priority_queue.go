package priorityqueue

import (
	"cmp"

	linkedlist "github.com/Wine1y/go-cs/pkg/containers/linked_list"
)

type LinkedListPriorityQueue[T cmp.Ordered] struct {
	*linkedlist.SinglyLinkedList[T]
}

func NewLinkedListPriorityQueue[T cmp.Ordered]() LinkedListPriorityQueue[T] {
	return LinkedListPriorityQueue[T]{
		SinglyLinkedList: linkedlist.NewSinglyLinkedList[T](),
	}
}

func (queue *LinkedListPriorityQueue[T]) Insert(data T) {
	queue.SinglyLinkedList.Append(data)
}

func (queue LinkedListPriorityQueue[T]) Length() int {
	return queue.SinglyLinkedList.Length()
}

func (queue LinkedListPriorityQueue[T]) GetMin() T {
	return queue.SinglyLinkedList.Get(queue.getMinIndex()).Data()
}

func (queue *LinkedListPriorityQueue[T]) DeleteMin() T {
	minIndex := queue.getMinIndex()
	node := queue.SinglyLinkedList.Get(minIndex)
	queue.SinglyLinkedList.Delete(minIndex)
	return node.Data()
}

func (queue LinkedListPriorityQueue[T]) GetMax() T {
	return queue.SinglyLinkedList.Get(queue.getMaxIndex()).Data()
}

func (queue *LinkedListPriorityQueue[T]) DeleteMax() T {
	maxIndex := queue.getMaxIndex()
	node := queue.SinglyLinkedList.Get(maxIndex)
	queue.SinglyLinkedList.Delete(maxIndex)
	return node.Data()
}

func (queue LinkedListPriorityQueue[T]) getMinIndex() int {
	if queue.Length() == 0 {
		panic("Can't get minimum item, priority queue is empty")
	}
	minIndex, min := 0, queue.SinglyLinkedList.First().Data()
	iterator := queue.SinglyLinkedList.Iterator()
	for i := 0; !iterator.Ended(); i++ {
		if iterator.Next().Data() < min {
			minIndex = i
		}
	}
	return minIndex
}

func (queue LinkedListPriorityQueue[T]) getMaxIndex() int {
	if queue.Length() == 0 {
		panic("Can't get maximum item, priority queue is empty")
	}
	maxIndex, max := 0, queue.SinglyLinkedList.First().Data()
	iterator := queue.SinglyLinkedList.Iterator()
	for i := 0; !iterator.Ended(); i++ {
		if iterator.Next().Data() > max {
			maxIndex = i
		}
	}
	return maxIndex
}

type SortedLinkedListPriorityQueue[T cmp.Ordered] struct {
	*linkedlist.SinglyLinkedList[T]
}

func NewSortedLinkedListPriorityQueue[T cmp.Ordered]() SortedLinkedListPriorityQueue[T] {
	return SortedLinkedListPriorityQueue[T]{
		SinglyLinkedList: linkedlist.NewSinglyLinkedList[T](),
	}
}

func (queue *SortedLinkedListPriorityQueue[T]) Insert(data T) {
	if queue.SinglyLinkedList.Length() == 0 {
		queue.SinglyLinkedList.Insert(0, data)
		return
	}
	i := 0
	for iterator := queue.SinglyLinkedList.Iterator(); !iterator.Ended(); {
		if iterator.Next().Data() > data {
			break
		}
		i++
	}
	queue.SinglyLinkedList.Insert(i, data)
}

func (queue SortedLinkedListPriorityQueue[T]) Length() int {
	return queue.SinglyLinkedList.Length()
}

func (queue SortedLinkedListPriorityQueue[T]) GetMin() T {
	return queue.SinglyLinkedList.Get(queue.getMinIndex()).Data()
}

func (queue *SortedLinkedListPriorityQueue[T]) DeleteMin() T {
	minIndex := queue.getMinIndex()
	node := queue.SinglyLinkedList.Get(minIndex)
	queue.SinglyLinkedList.Delete(minIndex)
	return node.Data()
}

func (queue SortedLinkedListPriorityQueue[T]) GetMax() T {
	return queue.SinglyLinkedList.Get(queue.getMaxIndex()).Data()
}

func (queue *SortedLinkedListPriorityQueue[T]) DeleteMax() T {
	maxIndex := queue.getMaxIndex()
	node := queue.SinglyLinkedList.Get(maxIndex)
	queue.SinglyLinkedList.Delete(maxIndex)
	return node.Data()
}

func (queue SortedLinkedListPriorityQueue[T]) getMinIndex() int {
	if queue.Length() == 0 {
		panic("Can't get minimum item, priority queue is empty")
	}
	return 0
}

func (queue SortedLinkedListPriorityQueue[T]) getMaxIndex() int {
	if queue.Length() == 0 {
		panic("Can't get maximum item, priority queue is empty")
	}
	return queue.SinglyLinkedList.Length() - 1
}
