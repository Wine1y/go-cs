package priorityqueue

import (
	"cmp"

	"github.com/Wine1y/go-cs/pkg/containers/tree"
)

type BinarySearchTreePriorityQueue[T cmp.Ordered] struct {
	tree.BinarySearchTree[T]
	length int
}

func NewBinarySearchTreePriorityQueue[T cmp.Ordered]() BinarySearchTreePriorityQueue[T] {
	return BinarySearchTreePriorityQueue[T]{
		BinarySearchTree: tree.BinarySearchTree[T]{},
		length:           0,
	}
}

func (queue *BinarySearchTreePriorityQueue[T]) Insert(data T) {
	queue.BinarySearchTree.AppendNode(data)
	queue.length++
}

func (queue BinarySearchTreePriorityQueue[T]) Length() int {
	return queue.length
}

func (queue BinarySearchTreePriorityQueue[T]) GetMin() T {
	return queue.BinarySearchTree.FindMin().Data()
}

func (queue *BinarySearchTreePriorityQueue[T]) DeleteMin() T {
	minNode := queue.FindMin()
	queue.BinarySearchTree.DeleteNode(minNode.Data())
	queue.length--
	return minNode.Data()
}

func (queue BinarySearchTreePriorityQueue[T]) GetMax() T {
	return queue.BinarySearchTree.FindMax().Data()
}

func (queue *BinarySearchTreePriorityQueue[T]) DeleteMax() T {
	maxNode := queue.FindMax()
	queue.BinarySearchTree.DeleteNode(maxNode.Data())
	queue.length--
	return maxNode.Data()
}
