package linkedlist

import (
	"reflect"
)

/*
Linked lists are used to store data, they can't be accessed in a constant time like arrays,
but they can be expanded in a constant time, which makes them perfect for a highly
dynamic data that grows and shrinks more often than it's accessed.
In linked lists elements are not stored one after another.
*/
type LinkedList[T any] interface {
	//Insert the specified item at the specified position
	Insert(index int, data T)
	//Get item from the specified position
	Get(index int) LinkedListNode[T]
	//Delete item at the specified position
	Delete(index int) LinkedListNode[T]
	//Get list length
	Length() int

	//Insert the specified item at the end of the list
	Append(data T)
	//Delete item from the end of the list
	Pop() LinkedListNode[T]
	//Get first item in the list
	First() LinkedListNode[T]
	//Get last item in the list
	Last() LinkedListNode[T]
}

type LinkedListNode[T any] interface {
	Data() T
	Next() LinkedListNode[T]
}

type linkedListIterator[T any] struct {
	node LinkedListNode[T]
}

func newLinkedListIterator[T any](start LinkedListNode[T]) linkedListIterator[T] {
	return linkedListIterator[T]{
		node: start,
	}
}

func (iterator linkedListIterator[T]) Ended() bool {
	return iterator.node == nil || reflect.ValueOf(iterator.node).IsNil()
}

func (iterator *linkedListIterator[T]) Next() LinkedListNode[T] {
	if iterator.Ended() {
		panic("Iterator has ended")
	}
	node := iterator.node
	iterator.node = node.Next()
	return node
}
