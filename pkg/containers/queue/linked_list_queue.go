package queue

type listNode[T any] struct {
	data T
	next *listNode[T]
}

/*
Linked list queue can't be full because linked list have no max capacity.
Linked list queue uses extra space to store pointers
*/
type LinkedListQueue[T any] struct {
	front  *listNode[T]
	end    *listNode[T]
	length int
}

func NewLinkedListQueue[T any]() LinkedListQueue[T] {
	return LinkedListQueue[T]{
		front:  nil,
		end:    nil,
		length: 0,
	}
}

func (queue *LinkedListQueue[T]) EnQueue(data T) {
	node := listNode[T]{data: data, next: nil}
	if queue.Empty() {
		queue.front = &node
		queue.end = &node
	} else {
		queue.end.next = &node
		queue.end = &node
	}
	queue.length++
}

func (queue *LinkedListQueue[T]) DeQueue() T {
	item := queue.Front()
	queue.front = queue.front.next
	queue.length--
	return item
}

func (queue *LinkedListQueue[T]) Front() T {
	if queue.Empty() {
		panic("Queue underflow: Can't dequeue from empty queue")
	}
	return queue.front.data
}

func (queue LinkedListQueue[T]) Length() int {
	return queue.length
}

func (queue LinkedListQueue[T]) Empty() bool {
	return queue.Length() == 0
}

func (queue LinkedListQueue[T]) Full() bool {
	return false
}
