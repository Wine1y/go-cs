package linkedlist

import (
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	deleted := list.Delete(1)
	tu.AssertEqualsNamed(t, "SLL delete", deleted.Data(), 2)
	list.Insert(1, 2)

	list.Pop()
	list.Append(3)

	expected := [3]int{1, 2, 3}
	i := 0
	for iterator := list.Iterator(); !iterator.Ended(); {
		tu.AssertEqualsNamed(t, "SLL iterate", iterator.Next().Data(), expected[i])
		i++
	}

}

func TestDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	deleted := list.Delete(1)
	tu.AssertEqualsNamed(t, "DLL delete", deleted.Data(), 2)
	list.Insert(1, 2)

	list.Pop()
	list.Append(3)

	expected := [3]int{1, 2, 3}
	i := 0
	for iterator := list.Iterator(); !iterator.Ended(); {
		tu.AssertEqualsNamed(t, "DLL iterate forwards", iterator.Next().Data(), expected[i])
		i++
	}

	i = list.Length() - 1
	for iterator := list.BackwardsIterator(); !iterator.Ended(); {
		tu.AssertEqualsNamed(t, "DLL iterate backwards", iterator.Next().Data(), expected[i])
		i--
	}
}

func TestCircularLinkedList(t *testing.T) {
	list := NewCircularLinkedList[int]()

	list.Append(1)
	tu.AssertEqualsNamed(t, "CLL single node points to itself", list.First().Next(), list.First())
	list.Append(2)
	list.Append(3)

	deleted := list.Delete(1)
	tu.AssertEqualsNamed(t, "CLL delete", deleted.Data(), 2)
	list.Insert(1, 2)

	list.Pop()
	list.Append(3)

	expected := [6]int{1, 2, 3, 1, 2, 3}
	iterator := list.Iterator()
	for i := 0; i < len(expected); i++ {
		tu.AssertEqualsNamed(t, "CLL iterate", iterator.Next().Data(), expected[i])
	}
}
