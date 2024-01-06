package tree

type BinaryTreeNode[T any] struct {
	data  T
	left  *BinaryTreeNode[T]
	right *BinaryTreeNode[T]
}

func NewBinaryTreeNode[T any](data T) BinaryTreeNode[T] {
	return BinaryTreeNode[T]{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (node BinaryTreeNode[T]) Data() T {
	return node.data
}

func (node BinaryTreeNode[T]) Children() []BinaryTreeNode[T] {
	children := make([]BinaryTreeNode[T], 0, 2)
	if node.left != nil {
		children = append(children, *node.left)
	}
	if node.right != nil {
		children = append(children, *node.right)
	}
	return children
}

func (node BinaryTreeNode[T]) Left() *BinaryTreeNode[T] {
	return node.left
}

func (node BinaryTreeNode[T]) Right() *BinaryTreeNode[T] {
	return node.right
}

func (node *BinaryTreeNode[T]) AppendLeft(data T) *BinaryTreeNode[T] {
	left := NewBinaryTreeNode(data)
	node.left = &left
	return &left
}

func (node *BinaryTreeNode[T]) AppendRight(data T) *BinaryTreeNode[T] {
	right := NewBinaryTreeNode(data)
	node.right = &right
	return &right
}

func (node *BinaryTreeNode[T]) AppendNode(data T) *BinaryTreeNode[T] {
	if node.left == nil {
		return node.AppendLeft(data)
	}
	if node.right == nil {
		return node.AppendRight(data)
	}
	panic("Can't add a third child node to a BinaryTreeNode")
}

func (node BinaryTreeNode[T]) Height() int {
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
In a binary tree, every node has up to 2 child nodes.
Binary tree is called "strict" if each node has exactly 2 or 0 children.
Binary tree is called "full" if each node has exactly 2 children and
all leaf nodes are at the same level.
Binary trees are used in compilers, compression algorithms and other data structures.
*/
type BinaryTree[T any] struct {
	root *BinaryTreeNode[T]
}

func NewBinaryTree[T any](rootValue T) BinaryTree[T] {
	root := NewBinaryTreeNode(rootValue)
	return BinaryTree[T]{
		root: &root,
	}
}

func (tree BinaryTree[T]) Root() *BinaryTreeNode[T] {
	return tree.root
}

func (tree BinaryTree[T]) Height() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.Height()
}
