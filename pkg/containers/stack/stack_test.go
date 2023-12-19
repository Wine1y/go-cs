package stack

import (
	tu "Wine1y/go-cs/internal/testing_utils"
	"testing"
)

func TestSliceStack(t *testing.T) {
	stack := NewSliceStackWithSize[int](3)
	tu.AssertEqualsNamed(t, "SliceStack is initially empty", stack.Empty(), true)

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
