package linkedlist

import "fmt"

type singlyLinkedListNode[T any] struct {
	data T
	next *singlyLinkedListNode[T]
}

/*
In a singly linked list, each element have a pointer to the next element.
Singly linked list can't be efficiently traversed backwards
*/
type SinglyLinkedList[T any] struct {
	first  *singlyLinkedListNode[T]
	length int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		first:  nil,
		length: 0,
	}
}

func (list *SinglyLinkedList[T]) Insert(index int, item T) {
	node := singlyLinkedListNode[T]{data: item, next: nil}
	if index == 0 {
		if list.first != nil {
			node.next = list.first
		}
		list.first = &node
	} else {
		previous := list.getNode(index - 1)
		node.next = previous.next
		previous.next = &node
	}
	list.length++
}

func (list SinglyLinkedList[T]) Get(index int) T {
	return list.getNode(index).data
}

func (list *SinglyLinkedList[T]) Delete(index int) T {
	if index < 0 || index >= list.length {
		panic(fmt.Sprintf("List index out of range - %v", index))
	}
	var node *singlyLinkedListNode[T]
	if index == 0 {
		node = list.first
		list.first = node.next
	} else {
		previous := list.getNode(index - 1)
		node = previous.next
		previous.next = node.next
	}
	list.length--
	return node.data
}

func (list SinglyLinkedList[T]) Length() int {
	return list.length
}

func (list *SinglyLinkedList[T]) Append(item T) {
	list.Insert(list.Length(), item)
}

func (list *SinglyLinkedList[T]) Pop() T {
	return list.Delete(list.Length() - 1)
}

func (list SinglyLinkedList[T]) First() T {
	return list.Get(0)
}

func (list SinglyLinkedList[T]) Last() T {
	return list.Get(list.Length() - 1)
}

func (list SinglyLinkedList[T]) getNode(index int) *singlyLinkedListNode[T] {
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
