package heap

import (
	"sort"
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func TestMaxBinaryHeap(t *testing.T) {
	heap := NewBinaryHeap[int](MaxHeapOrder)
	tu.AssertEqualsNamed(t, "Heap is initially empty", heap.Length(), 0)

	tu.AssertPanics(
		t,
		"Empty heap panics on \"Pop\"",
		func() { heap.Pop() },
	)

	tu.AssertPanics(
		t,
		"Empty heap panics on \"Replace\"",
		func() { heap.Replace(5) },
	)

	heap.Push(10)
	heap.Push(15)
	heap.Push(5)
	tu.AssertEqualsNamed(t, "Heap length increases when pushed", heap.Length(), 3)
	tu.AssertEqualsNamed(t, "Heap returns right value when peeked", heap.Peek(), 15)

	tu.AssertEqualsNamed(t, "Heap returns right value when popped", heap.Pop(), 15)
	tu.AssertEqualsNamed(t, "Heap length decreases when popped", heap.Length(), 2)

	heap.Replace(8)
	tu.AssertEqualsNamed(t, "Heap \"Replace\" works correctly when new > old", heap.Peek(), 8)

	heap.Replace(3)
	tu.AssertEqualsNamed(t, "Heap \"Replace\" works correctly when new < old", heap.Peek(), 5)

	values := []int{10, 3, 7, 4, 6}
	heapified := HeapifyBinaryHeap[int](values, MaxHeapOrder)

	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	for _, value := range values {
		tu.AssertEqualsNamed(t, "Heapify returns correct heap", heapified.Pop(), value)
	}
}

func TestMinBinaryHeap(t *testing.T) {
	heap := NewBinaryHeap[int](MinHeapOrder)
	tu.AssertEqualsNamed(t, "Heap is initially empty", heap.Length(), 0)

	tu.AssertPanics(
		t,
		"Empty heap panics on \"Pop\"",
		func() { heap.Pop() },
	)

	tu.AssertPanics(
		t,
		"Empty heap panics on \"Replace\"",
		func() { heap.Replace(5) },
	)

	heap.Push(10)
	heap.Push(15)
	heap.Push(5)
	tu.AssertEqualsNamed(t, "Heap length increases when pushed", heap.Length(), 3)
	tu.AssertEqualsNamed(t, "Heap returns right value when peeked", heap.Peek(), 5)

	tu.AssertEqualsNamed(t, "Heap returns right value when popped", heap.Pop(), 5)
	tu.AssertEqualsNamed(t, "Heap length decreases when popped", heap.Length(), 2)

	heap.Replace(20)
	tu.AssertEqualsNamed(t, "Heap \"Replace\" works correctly when new > old", heap.Peek(), 15)

	heap.Replace(8)
	tu.AssertEqualsNamed(t, "Heap \"Replace\" works correctly when new < old", heap.Peek(), 8)

	values := []int{10, 3, 7, 4, 6}
	heapified := HeapifyBinaryHeap[int](values, MinHeapOrder)

	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	for _, value := range values {
		tu.AssertEqualsNamed(t, "Heapify returns correct heap", heapified.Pop(), value)
	}
}
