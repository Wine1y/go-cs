package tree

/*
Trees are used to store data in a hierarchical order using nodes.
Each node can have child nodes. Nodes that have no children are called "leaf nodes".
Links between nodes are called "edges"
In trees nodes are not stored one after another.
*/
type Tree[T any] interface {
	//Get tree root node
	Root() *TreeNode[T]
	//Get height of the tree
	Height() int
}

type TreeNode[T any] interface {
	Data() T
	Children() []*TreeNode[T]
	Height() int
	AppendNode(data T) TreeNode[T]
}
