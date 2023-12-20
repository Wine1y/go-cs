package stack

import (
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func TestSliceStack(t *testing.T) {
	stack := NewSliceStackWithSize[int](3)
	tu.AssertEqualsNamed(t, "SliceStack is initially empty", stack.Empty(), true)
	tu.AssertPanics(
		t,
		"SliceStack panics on \"Pop\" underflow",
		func() { stack.Pop() },
	)
	tu.AssertPanics(
		t,
		"SliceStack panics on \"Top\" underflow",
		func() { stack.Top() },
	)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	tu.AssertEqualsNamed(t, "SliceStack length increases when pushed", stack.Length(), 3)
	tu.AssertEqualsNamed(t, "SliceStack is not empty when pushed", stack.Empty(), false)
	tu.AssertEqualsNamed(t, "SliceStack is full when length == size", stack.Full(), true)

	stack.Push(4)
	tu.AssertEqualsNamed(t, "SliceStack will grow when overflowed", stack.Length(), 4)

	tu.AssertEqualsNamed(t, "SliceStack popped returns right value", stack.Pop(), 4)
	tu.AssertEqualsNamed(t, "SliceStack returns right top value", stack.Top(), 3)
	tu.AssertEqualsNamed(t, "SliceStack length decreases when popped", stack.Length(), 3)
}

func TestArrayStack(t *testing.T) {
	stack := NewArrayStack[int](3)
	tu.AssertEqualsNamed(t, "ArrayStack is initially empty", stack.Empty(), true)
	tu.AssertPanics(
		t,
		"ArrayStack panics on \"Pop\" underflow",
		func() { stack.Pop() },
	)
	tu.AssertPanics(
		t,
		"ArrayStack panics on \"Top\" underflow",
		func() { stack.Top() },
	)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	tu.AssertEqualsNamed(t, "ArrayStack length increases when pushed", stack.Length(), 3)
	tu.AssertEqualsNamed(t, "ArrayStack is not empty when pushed", stack.Empty(), false)
	tu.AssertEqualsNamed(t, "ArrayStack is full when length == size", stack.Full(), true)

	tu.AssertPanics(
		t,
		"ArrayStack panics on overflow",
		func() { stack.Push(4) },
	)

	tu.AssertEqualsNamed(t, "ArrayStack popped returns right value", stack.Pop(), 3)
	tu.AssertEqualsNamed(t, "ArrayStack returns right top value", stack.Top(), 2)
	tu.AssertEqualsNamed(t, "ArrayStack length decreases when popped", stack.Length(), 2)
}

func TestLinkedListStack(t *testing.T) {
	stack := NewLinkedListStack[int]()
	tu.AssertEqualsNamed(t, "LinkedListStack is initially empty", stack.Empty(), true)
	tu.AssertPanics(
		t,
		"LinkedListStack panics on \"Pop\" underflow",
		func() { stack.Pop() },
	)
	tu.AssertPanics(
		t,
		"LinkedListStack panics on \"Top\" underflow",
		func() { stack.Top() },
	)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	tu.AssertEqualsNamed(t, "LinkedListStack length increases when pushed", stack.Length(), 3)
	tu.AssertEqualsNamed(t, "LinkedListStack is not empty when pushed", stack.Empty(), false)
	tu.AssertEqualsNamed(t, "LinkedListStack can't be full", stack.Full(), false)

	tu.AssertEqualsNamed(t, "LinkedListStack popped returns right value", stack.Pop(), 3)
	tu.AssertEqualsNamed(t, "LinkedListStack returns right top value", stack.Top(), 2)
	tu.AssertEqualsNamed(t, "LinkedListStack length decreases when popped", stack.Length(), 2)
}
