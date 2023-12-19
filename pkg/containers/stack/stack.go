package stack

/*
Stacks are used to store data, items are added to and popped from the stack at one end.
Stacks can't be randomly accessed.
You can only "Push" something to the stack and "Pop" something from it.
You can't insert something in the middle of the stack.
Stacks are usefull for some balancing or parsing operations.
*/
type Stack[T any] interface {
	//Push something on top of the stack
	Push(data T)
	//Remove item from top of the stack
	Pop() T
	//Get item from top of the stack without removing it from the stack
	Top() T
	//Get stack length
	Length() int
	//Check if stack is empty
	Empty() bool
	//Check if stack is full
	Full() bool
}
