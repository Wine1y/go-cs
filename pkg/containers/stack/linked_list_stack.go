package stack

import linkedlist "github.com/Wine1y/go-cs/pkg/containers/linked_list"

/*
Linked list stack can't be full because linked list have no max capacity.
Linked list stack uses extra space to store pointers
*/
type LinkedListStack[T any] struct {
	list *linkedlist.SinglyLinkedList[T]
}

func NewLinkedListStack[T any]() LinkedListStack[T] {
	return LinkedListStack[T]{
		list: linkedlist.NewSinglyLinkedList[T](),
	}
}

func (stack *LinkedListStack[T]) Push(data T) {
	stack.list.Insert(0, data)
}

func (stack *LinkedListStack[T]) Pop() T {
	if stack.Length() == 0 {
		panic("Stack underflow: can't pop from empty stack")
	}
	popped := stack.list.Delete(0)
	return popped.Data()
}

func (stack *LinkedListStack[T]) Top() T {
	if stack.Length() == 0 {
		panic("Stack underflow: can't get top item from empty stack")
	}
	return stack.list.Get(0).Data()
}

func (stack *LinkedListStack[T]) Length() int {
	return stack.list.Length()
}

func (stack *LinkedListStack[T]) Empty() bool {
	return stack.Length() == 0
}

func (stack *LinkedListStack[T]) Full() bool {
	return false
}
