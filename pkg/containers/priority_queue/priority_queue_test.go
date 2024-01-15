package priorityqueue

import (
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func testAscendingPriorityQueue(t *testing.T, queue AscendingPriorityQueue[int]) {
	tu.AssertEqualsNamed(t, "Priority queue is initially empty", queue.Length(), 0)

	tu.AssertPanics(
		t,
		"Empty priority queue panics on \"DeleteMin\"",
		func() { queue.DeleteMin() },
	)

	tu.AssertPanics(
		t,
		"Empty priority queue panics on \"GetMin\"",
		func() { queue.GetMin() },
	)

	queue.Insert(10)
	queue.Insert(15)
	queue.Insert(5)
	tu.AssertEqualsNamed(t, "Priority queue length increases when inserted", queue.Length(), 3)
	tu.AssertEqualsNamed(t, "Priority queue returns right \"GetMin\" value", queue.GetMin(), 5)

	tu.AssertEqualsNamed(t, "Priority queue returns right \"DeleteMin\" value", queue.DeleteMin(), 5)
	tu.AssertEqualsNamed(t, "Priority queue length decreases when deleted", queue.Length(), 2)

	expValues := []int{10, 15}
	for _, value := range expValues {
		tu.AssertEqualsNamed(t, "Priority queue have correct order", queue.DeleteMin(), value)
	}
}

func testDescendingPriorityQueue(t *testing.T, queue DescendingPriorityQueue[int]) {
	tu.AssertEqualsNamed(t, "Priority queue is initially empty", queue.Length(), 0)

	tu.AssertPanics(
		t,
		"Empty priority queue panics on \"DeleteMax\"",
		func() { queue.DeleteMax() },
	)

	tu.AssertPanics(
		t,
		"Empty priority queue panics on \"GetMax\"",
		func() { queue.GetMax() },
	)

	queue.Insert(10)
	queue.Insert(15)
	queue.Insert(5)
	tu.AssertEqualsNamed(t, "Priority queue length increases when inserted", queue.Length(), 3)
	tu.AssertEqualsNamed(t, "Priority queue returns right \"GetMax\" value", queue.GetMax(), 15)

	tu.AssertEqualsNamed(t, "Priority queue returns right \"DeleteMax\" value", queue.DeleteMax(), 15)
	tu.AssertEqualsNamed(t, "Priority queue length decreases when deleted", queue.Length(), 2)

	expValues := []int{10, 5}
	for _, value := range expValues {
		tu.AssertEqualsNamed(t, "Priority queue have correct order", queue.DeleteMax(), value)
	}
}

func testDoubleEndedPriorityQueue(t *testing.T, queue DoubleEndedPriorityQueue[int]) {
	t.Helper()
	testAscendingPriorityQueue(t, queue)
	for queue.Length() != 0 {
		queue.DeleteMin()
	}
	testDescendingPriorityQueue(t, queue)
}

func TestSlicePriorityQueue(t *testing.T) {
	queue := NewSlicePriorityQueue[int]()
	testDoubleEndedPriorityQueue(t, &queue)
}

func TestSortedSlicePriorityQueue(t *testing.T) {
	queue := NewSortedSlicePriorityQueue[int]()
	testDoubleEndedPriorityQueue(t, &queue)
}

func TestLinkedListPriorityQueue(t *testing.T) {
	queue := NewLinkedListPriorityQueue[int]()
	testDoubleEndedPriorityQueue(t, &queue)
}

func TestSortedLinkedListPriorityQueue(t *testing.T) {
	queue := NewSortedLinkedListPriorityQueue[int]()
	testDoubleEndedPriorityQueue(t, &queue)
}

func TestBinarySearchTreePriorityQueue(t *testing.T) {
	queue := NewBinarySearchTreePriorityQueue[int]()
	testDoubleEndedPriorityQueue(t, &queue)
}

func TestHeapAscendingPriorityQueue(t *testing.T) {
	queue := NewHeapAscendingPriorityQueue[int]()
	testAscendingPriorityQueue(t, &queue)
}

func TestHeapDescendingPriorityQueue(t *testing.T) {
	queue := NewHeapDescendingPriorityQueue[int]()
	testDescendingPriorityQueue(t, &queue)
}
