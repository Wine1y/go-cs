package queue

/*
Circular queues use circular arrays to store items.
CircularSliceQueue grows at overflow by reallocating the underlying slice.
*/
type CircularSliceQueue[T any] struct {
	slice []T
	front int
	end   int
	size  int
}

func NewCircularSliceQueue[T any]() CircularSliceQueue[T] {
	return CircularSliceQueue[T]{
		slice: make([]T, 1),
		front: -1, end: -1,
		size: 1,
	}
}

func NewCircularSliceQueueWithSize[T any](size int) CircularSliceQueue[T] {
	return CircularSliceQueue[T]{
		slice: make([]T, size),
		front: -1, end: -1,
		size: size,
	}
}

func (queue *CircularSliceQueue[T]) EnQueue(data T) {
	if queue.Full() {
		queue.resize()
	}
	if queue.front == -1 {
		queue.front = 0
	}
	queue.end = (queue.end + 1) % queue.size
	queue.slice[queue.end] = data
}

func (queue *CircularSliceQueue[T]) DeQueue() T {
	item := queue.Front()
	if queue.front == queue.end {
		queue.end = -1
		queue.front = -1
	} else {
		queue.front = (queue.front + 1) % queue.size
	}
	return item
}

func (queue *CircularSliceQueue[T]) Front() T {
	if queue.Empty() {
		panic("Queue underflow: Can't dequeue from empty queue")
	}
	return queue.slice[queue.front]
}

func (queue CircularSliceQueue[T]) Length() int {
	if queue.front > queue.end {
		return queue.size - queue.front + queue.end + 1
	} else {
		return queue.end - queue.front + 1
	}
}

func (queue CircularSliceQueue[T]) Empty() bool {
	return queue.front == -1
}

func (queue CircularSliceQueue[T]) Full() bool {
	return (queue.end+1)%queue.size == queue.front
}

func (queue *CircularSliceQueue[T]) resize() {
	newSlice := make([]T, queue.size*2)
	copy(newSlice, queue.slice)
	queue.slice = newSlice

	if queue.front > queue.end {
		for i := 0; i < queue.front; i++ {
			queue.slice[i+queue.size] = queue.slice[i]
		}
		queue.end = queue.end + queue.size
	}
	queue.size *= 2
}
