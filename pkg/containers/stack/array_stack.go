package stack

/*
Array stack can be full if underlying array is full.
When pushed, full ArrayStack will panic.
*/
type ArrayStack[T any] struct {
	array  []T
	size   int
	length int
}

func NewArrayStack[T any](size int) ArrayStack[T] {
	return ArrayStack[T]{
		array:  make([]T, size),
		size:   size,
		length: 0,
	}
}

func (stack *ArrayStack[T]) Push(data T) {
	if stack.Length() == stack.size {
		panic("Stack overflow: can't push into full stack")
	}
	stack.array[stack.length] = data
	stack.length++
}

func (stack *ArrayStack[T]) Pop() T {
	if stack.Length() == 0 {
		panic("Stack underflow: can't pop from empty stack")
	}
	popped := stack.array[stack.length-1]
	stack.length--
	return popped
}

func (stack *ArrayStack[T]) Top() T {
	if stack.Length() == 0 {
		panic("Stack underflow: can't get top item from empty stack")
	}
	return stack.array[stack.length-1]
}

func (stack *ArrayStack[T]) Length() int {
	return stack.length
}

func (stack *ArrayStack[T]) Empty() bool {
	return stack.Length() == 0
}

func (stack *ArrayStack[T]) Full() bool {
	return stack.length == stack.size
}
