package stack

/*
Slice stack (dynamic array stack) can be full if underlying slice is full.
But instead of throwing exception when pushed, the full SliceStack will grow (reallocate).
*/
type SliceStack[T any] struct {
	slice []T
}

func NewSliceStack[T any]() SliceStack[T] {
	return SliceStack[T]{
		slice: make([]T, 0, 1),
	}
}

func NewSliceStackWithSize[T any](size int) SliceStack[T] {
	return SliceStack[T]{
		slice: make([]T, 0, size),
	}
}

func (stack *SliceStack[T]) Push(data T) {
	stack.slice = append(stack.slice, data)
}

func (stack *SliceStack[T]) Pop() T {
	if stack.Length() == 0 {
		panic("Stack underflow: can't pop from empty stack")
	}
	popped := stack.slice[len(stack.slice)-1]
	stack.slice = stack.slice[:len(stack.slice)-1]
	return popped
}

func (stack *SliceStack[T]) Top() T {
	if stack.Length() == 0 {
		panic("Stack underflow: can't get top item from empty stack")
	}
	return stack.slice[len(stack.slice)-1]
}

func (stack *SliceStack[T]) Length() int {
	return len(stack.slice)
}

func (stack *SliceStack[T]) Empty() bool {
	return stack.Length() == 0
}

func (stack *SliceStack[T]) Full() bool {
	return len(stack.slice) == cap(stack.slice)
}
