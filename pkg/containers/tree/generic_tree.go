package tree

import (
	"github.com/Wine1y/go-cs/pkg/containers/queue"
	"github.com/Wine1y/go-cs/pkg/containers/stack"
	"github.com/Wine1y/go-cs/pkg/utils"
)

type GenericTreeNode[T any] struct {
	data        T
	firstChild  *GenericTreeNode[T]
	nextSibling *GenericTreeNode[T]
}

func NewGenericTreeNode[T any](data T) GenericTreeNode[T] {
	return GenericTreeNode[T]{
		data:        data,
		firstChild:  nil,
		nextSibling: nil,
	}
}

func (node GenericTreeNode[T]) Data() T {
	return node.data
}

func (node GenericTreeNode[T]) Children() []*GenericTreeNode[T] {
	if node.firstChild == nil {
		return make([]*GenericTreeNode[T], 0)
	}

	children := []*GenericTreeNode[T]{node.firstChild}

	for child := node.firstChild.nextSibling; child != nil; child = child.nextSibling {
		children = append(children, child)
	}
	return children
}

func (node *GenericTreeNode[T]) AppendNode(data T) *GenericTreeNode[T] {
	newChild := NewGenericTreeNode(data)
	if node.firstChild == nil {
		node.firstChild = &newChild
	} else {
		child := node.firstChild
		for child.nextSibling != nil {
			child = child.nextSibling
		}
		child.nextSibling = &newChild
	}
	return &newChild
}

func (node GenericTreeNode[T]) Height() int {
	maxHeight := 0
	for child := node.firstChild; child != nil; child = child.nextSibling {
		childMax := child.Height()
		if childMax > maxHeight {
			maxHeight = childMax
		}
	}
	return maxHeight + 1
}

/*
In a generic tree nodes can have any number of child nodes.
*/
type GenericTree[T any] struct {
	root *GenericTreeNode[T]
}

func NewGenericTree[T any](rootValue T) GenericTree[T] {
	root := NewGenericTreeNode(rootValue)
	return GenericTree[T]{
		root: &root,
	}
}

func (tree GenericTree[T]) Root() *GenericTreeNode[T] {
	return tree.root
}

func (tree GenericTree[T]) Height() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.Height()
}

func (tree GenericTree[T]) DFSIterator() utils.Iterator[*GenericTreeNode[T]] {
	iterator := newGenericTreeDFSIterator[T](tree.Root())
	return &iterator
}

func (tree GenericTree[T]) BFSIterator() utils.Iterator[*GenericTreeNode[T]] {
	iterator := newGenericTreeBFSIterator[T](tree.Root())
	return &iterator
}

type genericTreeDFSIterator[T any] struct {
	stack stack.Stack[*GenericTreeNode[T]]
}

type genericTreeBFSIterator[T any] struct {
	queue queue.Queue[*GenericTreeNode[T]]
}

func newGenericTreeDFSIterator[T any](root *GenericTreeNode[T]) genericTreeDFSIterator[T] {
	stack := stack.NewLinkedListStack[*GenericTreeNode[T]]()
	stack.Push(root)
	iterator := genericTreeDFSIterator[T]{
		stack: &stack,
	}
	return iterator
}

func newGenericTreeBFSIterator[T any](root *GenericTreeNode[T]) genericTreeBFSIterator[T] {
	queue := queue.NewLinkedListQueue[*GenericTreeNode[T]]()
	queue.EnQueue(root)
	iterator := genericTreeBFSIterator[T]{
		queue: &queue,
	}
	return iterator
}

func (iterator genericTreeDFSIterator[T]) Ended() bool {
	return iterator.stack.Empty()
}

func (iterator genericTreeBFSIterator[T]) Ended() bool {
	return iterator.queue.Empty()
}

func (iterator *genericTreeDFSIterator[T]) Next() *GenericTreeNode[T] {
	if iterator.Ended() {
		panic("Iterator has ended")
	}
	node := iterator.stack.Pop()
	children := node.Children()
	for i := len(children) - 1; i >= 0; i-- {
		iterator.stack.Push(children[i])
	}
	return node
}

func (iterator *genericTreeBFSIterator[T]) Next() *GenericTreeNode[T] {
	if iterator.Ended() {
		panic("Iterator has ended")
	}
	node := iterator.queue.DeQueue()
	for _, child := range node.Children() {
		iterator.queue.EnQueue(child)
	}
	return node
}
