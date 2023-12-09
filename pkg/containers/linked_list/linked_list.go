package linkedlist

/*
Linked lists are used to store data, they can't be accessed in a constant time like arrays,
but they can be expanded in a constant time, which makes them perfect for a highly
dynamic data that grows and shrinks more often the it's accessed.
In a linked lists, elements are not stored one after another.
*/
type LinkedList[T any] interface {
	//Insert the specified item at the specified position
	Insert(index int, item T)
	//Get item from the specified position
	Get(index int) T
	//Delete item at the specified position
	Delete(index int) T
	//Get list length
	Length() int

	//Insert the specified item at the end of the list
	Append(item T)
	//Delete item from the end of the list
	Pop() T
	//Get first item in the list
	First() T
	//Get last item in the list
	Last() T
}
