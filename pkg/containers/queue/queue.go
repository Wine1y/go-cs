package queue

/*
Queues are used to store data, items are added to the queue at one and and popped at another.
Queue is a First In First Out (FIFO) data structure.
Queues can't be randomly accessed.
You can only "EnQueue" something to the queue and "DeQueue" something from it.
You can't insert something in the middle of the Queue.
Queues are used when the order of arrival is important. (Scheduling, IO)
*/
type Queue[T any] interface {
	//Insert something at the end of the queue
	EnQueue(data T)
	//Remove and return item at the front of the queue
	DeQueue() T
	//Get item from the fron of the queue without removing it from the queue.
	Front() T
	//Get queue length
	Length() int
	//Check if queue is empty
	Empty() bool
	//Check if queue is full
	Full() bool
}
