package linkedlist

import "fmt"

type SinglyLinkedListNode[T any] struct {
	data T
	next *SinglyLinkedListNode[T]
}

func (node SinglyLinkedListNode[T]) Data() T {
	return node.data
}

func (node SinglyLinkedListNode[T]) Next() *SinglyLinkedListNode[T] {
	return node.next
}

/*
In a singly linked list, each element have a pointer to the next element.
Singly linked list can't be efficiently traversed backwards
*/
type SinglyLinkedList[T any] struct {
	first  *SinglyLinkedListNode[T]
	length int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		first:  nil,
		length: 0,
	}
}

func (list *SinglyLinkedList[T]) Insert(index int, data T) {
	node := SinglyLinkedListNode[T]{data: data, next: nil}
	if index == 0 {
		if list.first != nil {
			node.next = list.first
		}
		list.first = &node
	} else {
		previous := list.Get(index - 1)
		node.next = previous.next
		previous.next = &node
	}
	list.length++
}

func (list SinglyLinkedList[T]) Get(index int) *SinglyLinkedListNode[T] {
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

func (list *SinglyLinkedList[T]) Delete(index int) *SinglyLinkedListNode[T] {
	if index < 0 || index >= list.length {
		panic(fmt.Sprintf("List index out of range - %v", index))
	}
	var node *SinglyLinkedListNode[T]
	if index == 0 {
		node = list.first
		list.first = node.next
	} else {
		previous := list.Get(index - 1)
		node = previous.next
		previous.next = node.next
	}
	list.length--
	return node
}

func (list SinglyLinkedList[T]) Length() int {
	return list.length
}

func (list *SinglyLinkedList[T]) Append(data T) {
	list.Insert(list.Length(), data)
}

func (list *SinglyLinkedList[T]) Pop() *SinglyLinkedListNode[T] {
	return list.Delete(list.Length() - 1)
}

func (list SinglyLinkedList[T]) First() *SinglyLinkedListNode[T] {
	return list.Get(0)
}

func (list SinglyLinkedList[T]) Last() *SinglyLinkedListNode[T] {
	return list.Get(list.Length() - 1)
}
