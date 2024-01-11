package linkedlist

import (
	"fmt"

	"github.com/Wine1y/go-cs/pkg/utils"
)

/*
In a circular linked list, each element have a pointer to the next element.
There is no end in a circular linked list.
The last element has its next pointer set to the first element.
Circulat linked lists can be traversed infinitely.
*/
type CircularLinkedList[T any] struct {
	first  *singlyLinkedListNode[T]
	length int
}

func NewCircularLinkedList[T any]() *CircularLinkedList[T] {
	return &CircularLinkedList[T]{
		first:  nil,
		length: 0,
	}
}

func (list *CircularLinkedList[T]) Insert(index int, data T) {
	node := singlyLinkedListNode[T]{data: data, next: nil}
	node.next = &node
	if index == 0 {
		if list.length > 0 {
			node.next = list.first
			last := list.getNode(list.length - 1)
			last.next = &node
		}
		list.first = &node
	} else {
		previous := list.getNode(index - 1)
		node.next = previous.next
		previous.next = &node
	}
	list.length++
}

func (list CircularLinkedList[T]) Get(index int) LinkedListNode[T] {
	return list.getNode(index)
}

func (list CircularLinkedList[T]) getNode(index int) *singlyLinkedListNode[T] {
	if index < 0 || index >= list.length {
		panic(fmt.Sprintf("List index out of range - %v", index))
	}

	i := 0
	node := list.first
	for {
		if i == index {
			return node
		}
		node = node.next
		i++
	}
}

func (list *CircularLinkedList[T]) Delete(index int) LinkedListNode[T] {
	if index < 0 || index >= list.length {
		panic(fmt.Sprintf("List index out of range - %v", index))
	}
	var node *singlyLinkedListNode[T]
	if index == 0 {
		node = list.first
		if list.length > 1 {
			list.first = node.next
			last := list.getNode(list.length - 1)
			last.next = list.first
		} else {
			list.first = nil
		}
	} else {
		previous := list.getNode(index - 1)
		node = previous.next
		previous.next = node.next
	}
	list.length--
	return node
}

func (list CircularLinkedList[T]) Length() int {
	return list.length
}

func (list *CircularLinkedList[T]) Append(data T) {
	list.Insert(list.Length(), data)
}

func (list *CircularLinkedList[T]) Pop() LinkedListNode[T] {
	return list.Delete(list.Length() - 1)
}

func (list CircularLinkedList[T]) First() LinkedListNode[T] {
	return list.Get(0)
}

func (list CircularLinkedList[T]) Last() LinkedListNode[T] {
	return list.Get(list.Length() - 1)
}

func (list CircularLinkedList[T]) Iterator() utils.Iterator[LinkedListNode[T]] {
	iterator := newLinkedListIterator[T](list.First())
	return &iterator
}
