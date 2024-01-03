package queue

import (
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func TestCircularArrayQueue(t *testing.T) {
	queue := NewCircularArrayQueue[int](3)
	tu.AssertEqualsNamed(t, "CircularArrayQueue is initially empty", queue.Empty(), true)

	tu.AssertPanics(
		t,
		"CircularArrayQueue panics on \"DeQueue\" underflow",
		func() { queue.DeQueue() },
	)

	tu.AssertPanics(
		t,
		"CircularArrayQueue panics on \"Front\" underflow",
		func() { queue.Front() },
	)

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	tu.AssertEqualsNamed(t, "CircularArrayQueue length increases when enqued", queue.Length(), 3)
	tu.AssertEqualsNamed(t, "CircularArrayQueue is not empty when enqued", queue.Empty(), false)
	tu.AssertEqualsNamed(t, "CircularArrayQueue is full when length == size", queue.Full(), true)

	tu.AssertPanics(
		t,
		"CircularArrayQueue panics on overflow",
		func() { queue.EnQueue(4) },
	)

	tu.AssertEqualsNamed(t, "CircularArrayQueue dequeue returns right value", queue.DeQueue(), 1)
	tu.AssertEqualsNamed(t, "CircularArrayQueue returns right front value", queue.Front(), 2)
	tu.AssertEqualsNamed(t, "CircularArrayQueue length decreases when dequeued", queue.Length(), 2)

	queue.DeQueue()
	queue.EnQueue(10)
	queue.EnQueue(20)
	tu.AssertEqualsNamed(t, "CircularArrayQueue returns right length when front > end", queue.Length(), 3)
}

func TestCircularSliceQueue(t *testing.T) {
	queue := NewCircularSliceQueueWithSize[int](3)
	tu.AssertEqualsNamed(t, "CircularSliceQueue is initially empty", queue.Empty(), true)

	tu.AssertPanics(
		t,
		"CircularSliceQueue panics on \"DeQueue\" underflow",
		func() { queue.DeQueue() },
	)

	tu.AssertPanics(
		t,
		"CircularSliceQueue panics on \"Front\" underflow",
		func() { queue.Front() },
	)

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	tu.AssertEqualsNamed(t, "CircularSliceQueue length increases when enqued", queue.Length(), 3)
	tu.AssertEqualsNamed(t, "CircularSliceQueue is not empty when enqued", queue.Empty(), false)
	tu.AssertEqualsNamed(t, "CircularSliceQueue is full when length == size", queue.Full(), true)

	queue.EnQueue(4)
	tu.AssertEqualsNamed(t, "CircularSliceQueue grows when overflowed", queue.Length(), 4)

	tu.AssertEqualsNamed(t, "CircularSliceQueue dequeue returns right value", queue.DeQueue(), 1)
	tu.AssertEqualsNamed(t, "CircularSliceQueue returns right front value", queue.Front(), 2)
	tu.AssertEqualsNamed(t, "CircularSliceQueue length decreases when dequeued", queue.Length(), 3)

	queue.DeQueue()
	queue.DeQueue()
	queue.EnQueue(10)
	queue.EnQueue(20)
	tu.AssertEqualsNamed(t, "CircularSliceQueue returns right length when front > end", queue.Length(), 3)
}

func TestLinkedListQueue(t *testing.T) {
	queue := NewLinkedListQueue[int]()
	tu.AssertEqualsNamed(t, "LinkedListQueue is initially empty", queue.Empty(), true)

	tu.AssertPanics(
		t,
		"LinkedListQueue panics on \"DeQueue\" underflow",
		func() { queue.DeQueue() },
	)

	tu.AssertPanics(
		t,
		"LinkedListQueue panics on \"Front\" underflow",
		func() { queue.Front() },
	)

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	tu.AssertEqualsNamed(t, "LinkedListQueue length increases when enqued", queue.Length(), 3)
	tu.AssertEqualsNamed(t, "LinkedListQueue is not empty when enqued", queue.Empty(), false)
	tu.AssertEqualsNamed(t, "LinkedListQueue can't be full", queue.Full(), false)

	tu.AssertEqualsNamed(t, "LinkedListQueue dequeue returns right value", queue.DeQueue(), 1)
	tu.AssertEqualsNamed(t, "LinkedListQueue returns right front value", queue.Front(), 2)
	tu.AssertEqualsNamed(t, "LinkedListQueue length decreases when dequeued", queue.Length(), 2)

	queue.DeQueue()
	queue.EnQueue(10)
	queue.EnQueue(20)
	tu.AssertEqualsNamed(t, "LinkedListQueue returns right length when front > end", queue.Length(), 3)
}
