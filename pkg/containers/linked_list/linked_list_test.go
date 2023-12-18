package linkedlist

import (
	tu "Wine1y/go-cs/internal/testing_utils"
	"testing"
)

func testLinkedListInsert(t *testing.T, list LinkedList[int]) {
	list.Append(1)
	list.Append(2)
	list.Append(3)

	tu.AssertEqualsNamed(t, "length", list.Length(), 3)
	tu.AssertEqualsNamed(t, "first item", list.First(), 1)
	tu.AssertEqualsNamed(t, "last item", list.Last(), 3)

	list.Insert(0, 10)
	tu.AssertEqualsNamed(t, "length", list.Length(), 4)
	tu.AssertEqualsNamed(t, "item at index 0", list.Get(0), 10)

	list.Insert(4, 20)
	tu.AssertEqualsNamed(t, "length", list.Length(), 5)
	tu.AssertEqualsNamed(t, "item at index 4", list.Get(4), 20)

	list.Insert(2, 30)
	tu.AssertEqualsNamed(t, "length", list.Length(), 6)
	tu.AssertEqualsNamed(t, "item at index 2", list.Get(2), 30)
}

func testLinkedListDelete(t *testing.T, list LinkedList[int]) {
	list.Append(1)
	list.Append(2)
	list.Append(3)

	tu.AssertEqualsNamed(t, "length", list.Length(), 3)
	tu.AssertEqualsNamed(t, "first item", list.First(), 1)
	tu.AssertEqualsNamed(t, "last item", list.Last(), 3)

	list.Delete(1)
	tu.AssertEqualsNamed(t, "length", list.Length(), 2)
	tu.AssertEqualsNamed(t, "first item", list.First(), 1)
	tu.AssertEqualsNamed(t, "last item", list.Last(), 3)

	list.Pop()
	tu.AssertEqualsNamed(t, "length", list.Length(), 1)
	tu.AssertEqualsNamed(t, "last item", list.Last(), 1)

	list.Pop()
	tu.AssertEqualsNamed(t, "length", list.Length(), 0)
}

func TestSinglyLinkedList(t *testing.T) {
	t.Run("testLinkedListInsert", func(t *testing.T) {
		testLinkedListInsert(t, NewSinglyLinkedList[int]())
	})

	t.Run("testLinkedListDelete", func(t *testing.T) {
		testLinkedListDelete(t, NewSinglyLinkedList[int]())
	})
}

func TestDoublyLinkedList(t *testing.T) {
	t.Run("testLinkedListInsert", func(t *testing.T) {
		testLinkedListInsert(t, NewDoublyLinkedList[int]())
	})

	t.Run("testLinkedListDelete", func(t *testing.T) {
		testLinkedListDelete(t, NewDoublyLinkedList[int]())
	})
}

func TestCircularLinkedList(t *testing.T) {
	t.Run("testLinkedListInsert", func(t *testing.T) {
		testLinkedListInsert(t, NewCircularLinkedList[int]())
	})

	t.Run("testLinkedListDelete", func(t *testing.T) {
		testLinkedListDelete(t, NewCircularLinkedList[int]())
	})
}
