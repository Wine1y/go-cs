package tree

import (
	"cmp"

	"github.com/Wine1y/go-cs/pkg/containers/stack"
	"github.com/Wine1y/go-cs/pkg/utils"
)

type BinarySearchTreeNode[T cmp.Ordered] struct {
	data  T
	left  *BinarySearchTreeNode[T]
	right *BinarySearchTreeNode[T]
}

func NewBinarySearchTreeNode[T cmp.Ordered](data T) BinarySearchTreeNode[T] {
	return BinarySearchTreeNode[T]{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (node BinarySearchTreeNode[T]) Data() T {
	return node.data
}

func (node BinarySearchTreeNode[T]) Children() []*BinarySearchTreeNode[T] {
	children := make([]*BinarySearchTreeNode[T], 0, 2)
	if node.left != nil {
		children = append(children, node.left)
	}
	if node.right != nil {
		children = append(children, node.right)
	}
	return children
}

func (node BinarySearchTreeNode[T]) Left() *BinarySearchTreeNode[T] {
	return node.left
}

func (node BinarySearchTreeNode[T]) Right() *BinarySearchTreeNode[T] {
	return node.right
}

func (node *BinarySearchTreeNode[T]) AppendLeft(data T) *BinarySearchTreeNode[T] {
	if node.left == nil {
		newChild := NewBinarySearchTreeNode(data)
		node.left = &newChild
		return &newChild
	}
	return node.left.AppendNode(data)
}

func (node *BinarySearchTreeNode[T]) AppendRight(data T) *BinarySearchTreeNode[T] {
	if node.right == nil {
		newChild := NewBinarySearchTreeNode(data)
		node.right = &newChild
		return &newChild
	}
	return node.right.AppendNode(data)
}

func (node *BinarySearchTreeNode[T]) AppendNode(data T) *BinarySearchTreeNode[T] {
	if data == node.data {
		return node
	}
	if data < node.Data() {
		return node.AppendLeft(data)
	}
	return node.AppendRight(data)
}

func (node BinarySearchTreeNode[T]) FindNode(data T) *BinarySearchTreeNode[T] {
	if data < node.Data() {
		if node.Left() == nil {
			return nil
		}
		return node.Left().FindNode(data)
	}
	if data > node.Data() {
		if node.Right() == nil {
			return nil
		}
		return node.Right().FindNode(data)
	}
	return &node
}

func (node BinarySearchTreeNode[T]) FindMin() *BinarySearchTreeNode[T] {
	if node.Left() == nil {
		return &node
	}
	return node.Left().FindMin()
}

func (node BinarySearchTreeNode[T]) FindMax() *BinarySearchTreeNode[T] {
	if node.Right() == nil {
		return &node
	}
	return node.Right().FindMax()
}

func (node *BinarySearchTreeNode[T]) DeleteNode(data T) *BinarySearchTreeNode[T] {
	if data < node.Data() && node.Left() != nil {
		node.left = node.Left().DeleteNode(data)
	}
	if data > node.Data() && node.Right() != nil {
		node.right = node.Right().DeleteNode(data)
	}
	if data == node.Data() {
		children := node.Children()
		switch len(children) {
		case 2:
			leftLargest := node.Left().FindMax()
			node.data = leftLargest.Data()
			node.left = node.Left().DeleteNode(leftLargest.Data())
		case 1:
			return children[0]
		case 0:
			return nil
		}
	}
	return node
}

func (node BinarySearchTreeNode[T]) Height() int {
	maxHeight := 0
	if node.left != nil {
		lHeight := node.left.Height()
		if lHeight > maxHeight {
			maxHeight = lHeight
		}
	}
	if node.right != nil {
		rHeight := node.right.Height()
		if rHeight > maxHeight {
			maxHeight = rHeight
		}
	}
	return maxHeight + 1
}

/*
BinarySearchTree (BST) is a BinaryTree with data constraints.
In BST, all elements of a node's left subtree should be less than the node itself
And all elements of the right subtree should be greater.
These constraints allow BST to be searched in O(logN) time instead of O(N).
*/
type BinarySearchTree[T cmp.Ordered] struct {
	root *BinarySearchTreeNode[T]
}

func NewBinarySearchTree[T cmp.Ordered](rootValue T) BinarySearchTree[T] {
	root := NewBinarySearchTreeNode(rootValue)
	return BinarySearchTree[T]{
		root: &root,
	}
}

func (tree BinarySearchTree[T]) Root() *BinarySearchTreeNode[T] {
	return tree.root
}

func (tree BinarySearchTree[T]) Height() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.Height()
}

func (tree *BinarySearchTree[T]) AppendNode(data T) *BinarySearchTreeNode[T] {
	if tree.Root() == nil {
		node := NewBinarySearchTreeNode(data)
		tree.root = &node
		return &node
	}
	return tree.Root().AppendNode(data)
}

func (tree *BinarySearchTree[T]) DeleteNode(data T) {
	if tree.Root() != nil {
		tree.root = tree.root.DeleteNode(data)
	}
}

func (tree *BinarySearchTree[T]) FindNode(data T) *BinarySearchTreeNode[T] {
	if tree.Root() == nil {
		return nil
	}
	return tree.Root().FindNode(data)
}

func (tree BinarySearchTree[T]) FindMin() *BinarySearchTreeNode[T] {
	if tree.Root() == nil {
		return nil
	}
	return tree.Root().FindMin()
}

func (tree BinarySearchTree[T]) FindMax() *BinarySearchTreeNode[T] {
	if tree.Root() == nil {
		return nil
	}
	return tree.Root().FindMax()
}

func (tree BinarySearchTree[T]) PreOrderIterator() utils.Iterator[*BinarySearchTreeNode[T]] {
	iterator := binarySearchTreePreOrderIterator[T]{
		binarySearchTreeIterator: newBinarySearchTreeIterator[T](tree.Root()),
	}
	iterator.stack.Push(tree.Root())
	return &iterator
}

func (tree BinarySearchTree[T]) InOrderIterator() utils.Iterator[*BinarySearchTreeNode[T]] {
	iterator := binarySearchTreeInOrderIterator[T]{
		binarySearchTreeIterator: newBinarySearchTreeIterator[T](tree.Root()),
	}
	return &iterator
}

func (tree BinarySearchTree[T]) PostOrderIterator() utils.Iterator[*BinarySearchTreeNode[T]] {
	iterator := binarySearchTreePostOrderIterator[T]{
		binarySearchTreeIterator: newBinarySearchTreeIterator[T](tree.Root()),
	}
	iterator.stack.Push(tree.Root())
	return &iterator
}

type binarySearchTreeIterator[T cmp.Ordered] struct {
	stack stack.Stack[*BinarySearchTreeNode[T]]
	head  *BinarySearchTreeNode[T]
}

type binarySearchTreePreOrderIterator[T cmp.Ordered] struct {
	binarySearchTreeIterator[T]
}

type binarySearchTreeInOrderIterator[T cmp.Ordered] struct {
	binarySearchTreeIterator[T]
}

type binarySearchTreePostOrderIterator[T cmp.Ordered] struct {
	binarySearchTreeIterator[T]
}

func newBinarySearchTreeIterator[T cmp.Ordered](head *BinarySearchTreeNode[T]) binarySearchTreeIterator[T] {
	stack := stack.NewLinkedListStack[*BinarySearchTreeNode[T]]()
	return binarySearchTreeIterator[T]{
		stack: &stack,
		head:  head,
	}
}

func (iterator binarySearchTreeIterator[T]) Ended() bool {
	return iterator.stack.Empty()
}

func (iterator binarySearchTreeInOrderIterator[T]) Ended() bool {
	return iterator.stack.Empty() && iterator.head == nil
}

func (iterator *binarySearchTreePreOrderIterator[T]) Next() *BinarySearchTreeNode[T] {
	if iterator.Ended() {
		panic("Iterator has ended")
	}
	node := iterator.stack.Pop()
	if node.Right() != nil {
		iterator.stack.Push(node.Right())
	}
	if node.Left() != nil {
		iterator.stack.Push(node.Left())
	}
	return node
}

func (iterator *binarySearchTreeInOrderIterator[T]) Next() *BinarySearchTreeNode[T] {
	if iterator.Ended() {
		panic("Iterator has ended")
	}
	for {
		if iterator.head == nil {
			node := iterator.stack.Pop()
			iterator.head = node.Right()
			return node
		}
		iterator.stack.Push(iterator.head)
		iterator.head = iterator.head.Left()
	}
}

func (iterator *binarySearchTreePostOrderIterator[T]) Next() *BinarySearchTreeNode[T] {
	if iterator.Ended() {
		panic("Iterator has ended")
	}
	for {
		node := iterator.stack.Top()
		processed := node.Right() == iterator.head || node.Left() == iterator.head
		leaf := node.Right() == nil && node.Left() == nil
		if leaf || processed {
			iterator.stack.Pop()
			iterator.head = node
			return node
		}
		if node.Right() != nil {
			iterator.stack.Push(node.Right())
		}
		if node.Left() != nil {
			iterator.stack.Push(node.Left())
		}
	}
}
