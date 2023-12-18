package linkedlist

import "fmt"

type doublyLinkedListNode[T any] struct {
	data     T
	previous *doublyLinkedListNode[T]
	next     *doublyLinkedListNode[T]
}

/*
In a doubly linked list, each element have a pointer to the previous and to the next element.
Doubly linked lists can be traversed backwards but require more memory due to one more pointer.
*/
type DoublyLinkedList[T any] struct {
	first  *doublyLinkedListNode[T]
	last   *doublyLinkedListNode[T]
	length int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

func (list *DoublyLinkedList[T]) Insert(index int, item T) {
	node := doublyLinkedListNode[T]{data: item, previous: nil, next: nil}
	if index == 0 {
		if list.first != nil {
			node.next = list.first
			list.first.previous = &node
		}
		list.first = &node
	} else {
		previous := list.getNode(index - 1)

		node.next = previous.next
		node.previous = previous
		previous.next = &node

		if node.next != nil {
			node.next.previous = &node
		}
	}
	if index == list.length {
		list.last = &node
	}
	list.length++
}

func (list DoublyLinkedList[T]) Get(index int) T {
	return list.getNode(index).data
}

func (list *DoublyLinkedList[T]) Delete(index int) T {
	if index < 0 || index >= list.length {
		panic(fmt.Sprintf("List index out of range - %v", index))
	}
	var node *doublyLinkedListNode[T]
	if index == 0 {
		node = list.first
		list.first = node.next
		if list.first != nil {
			list.first.previous = nil
		}
	} else {
		previous := list.getNode(index - 1)
		node = previous.next
		previous.next = node.next
		if node.next != nil {
			node.next.previous = previous
		}
	}
	if index == list.length-1 {
		list.last = node.previous
	}
	list.length--
	return node.data
}

func (list DoublyLinkedList[T]) Length() int {
	return list.length
}

func (list *DoublyLinkedList[T]) Append(item T) {
	list.Insert(list.Length(), item)
}

func (list *DoublyLinkedList[T]) Pop() T {
	return list.Delete(list.Length() - 1)
}

func (list DoublyLinkedList[T]) First() T {
	return list.Get(0)
}

func (list DoublyLinkedList[T]) Last() T {
	return list.Get(list.Length() - 1)
}

func (list DoublyLinkedList[T]) getNode(index int) *doublyLinkedListNode[T] {
	if index < 0 || index >= list.length {
		panic(fmt.Sprintf("List index out of range - %v", index))
	}

	var (
		i         int
		node      *doublyLinkedListNode[T]
		backwards bool
	)

	if index < list.length/2 {
		i = 0
		node = list.first
		backwards = false
	} else {
		i = list.length - 1
		node = list.last
		backwards = true
	}

	for {
		if i == index {
			return node
		}
		if backwards {
			node = node.previous
			i--
		} else {
			node = node.next
			i++
		}
	}
}
