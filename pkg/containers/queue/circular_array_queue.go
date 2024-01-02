package queue

type CircularArrayQueue[T any] struct {
	array []T
	front int
	end   int
	size  int
}

func NewCircularArrayQueue[T any](size int) CircularArrayQueue[T] {
	return CircularArrayQueue[T]{
		array: make([]T, size),
		front: -1, end: -1,
		size: size,
	}
}

func (queue *CircularArrayQueue[T]) EnQueue(data T) {
	if queue.Full() {
		panic("Queue overflow: Can't enqueue into full queue")
	}
	if queue.front == -1 {
		queue.front = 0
	}
	queue.end = (queue.end + 1) % queue.size
	queue.array[queue.end] = data
}

func (queue *CircularArrayQueue[T]) DeQueue() T {
	item := queue.Front()
	if queue.front == queue.end {
		queue.end = -1
		queue.front = -1
	} else {
		queue.front = (queue.front + 1) % queue.size
	}
	return item
}

func (queue *CircularArrayQueue[T]) Front() T {
	if queue.Empty() {
		panic("Queue underflow: Can't dequeue from empty queue")
	}
	return queue.array[queue.front]
}

func (queue CircularArrayQueue[T]) Length() int {
	if queue.front > queue.end {
		return queue.size - queue.front + queue.end + 1
	} else {
		return queue.end - queue.front + 1
	}
}

func (queue CircularArrayQueue[T]) Empty() bool {
	return queue.front == -1
}

func (queue CircularArrayQueue[T]) Full() bool {
	return (queue.end+1)%queue.size == queue.front
}
