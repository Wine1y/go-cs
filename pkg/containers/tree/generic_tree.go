package tree

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

func (node GenericTreeNode[T]) Children() []GenericTreeNode[T] {
	if node.firstChild == nil {
		return make([]GenericTreeNode[T], 0)
	}

	children := []GenericTreeNode[T]{*node.firstChild}

	for child := node.firstChild.nextSibling; child != nil; child = child.nextSibling {
		children = append(children, *child)
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
