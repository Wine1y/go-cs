package linkedlist

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
	Next() *LinkedListNode[T]
}
