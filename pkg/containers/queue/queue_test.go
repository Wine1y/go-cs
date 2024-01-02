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
