package tree

import (
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func TestGenericTree(t *testing.T) {
	tu.AssertEqualsNamed(t, "Empty tree height == 0", GenericTree[int]{}.Height(), 0)

	tree := NewGenericTree[int](1)
	tu.AssertEqualsNamed(t, "Tree with only root node height == 1", tree.Height(), 1)

	tree.Root().AppendNode(2)
	tree.Root().AppendNode(3)
	subNode := tree.Root().AppendNode(4)
	tu.AssertEqualsNamed(t, "Tree with appended nodes height is right", tree.Height(), 2)

	subNode.AppendNode(5)
	tu.AssertEqualsNamed(t, "Tree with appended subnodes height is right", tree.Height(), 3)

	rootChilds := tree.Root().Children()
	childExpValues := []int{2, 3, 4}
	tu.AssertEqualsNamed(t, "Node children has correct length", len(rootChilds), len(childExpValues))

	for i := range rootChilds {
		tu.AssertEqualsNamed(t, "Tree node returns correct children", rootChilds[i].Data(), childExpValues[i])
	}
}

func TestBinaryTree(t *testing.T) {
	tu.AssertEqualsNamed(t, "Empty tree height == 0", BinaryTree[int]{}.Height(), 0)

	tree := NewBinaryTree[int](1)
	tu.AssertEqualsNamed(t, "Tree with only root node height == 1", tree.Height(), 1)

	left := tree.Root().AppendNode(2)
	tree.Root().AppendNode(3)
	tu.AssertEqualsNamed(t, "AppendNode appends left node first", tree.Root().Left().Data(), 2)
	tu.AssertEqualsNamed(t, "AppendNode appends right node second", tree.Root().Right().Data(), 3)
	tu.AssertEqualsNamed(t, "Tree with appended nodes height is right", tree.Height(), 2)

	tu.AssertPanics(
		t,
		"BinaryTreeNode panics when trying to append a third child node",
		func() { tree.Root().AppendNode(4) },
	)

	left.AppendLeft(4)
	tu.AssertEqualsNamed(t, "Tree with appended subnodes height is right", tree.Height(), 3)

	rootChilds := tree.Root().Children()
	childExpValues := []int{2, 3}
	tu.AssertEqualsNamed(t, "Node children has correct length", len(rootChilds), len(childExpValues))

	for i := range rootChilds {
		tu.AssertEqualsNamed(t, "Tree node returns correct children", rootChilds[i].Data(), childExpValues[i])
	}
}

func TestBinarySearchTree(t *testing.T) {
	tu.AssertEqualsNamed(t, "Empty tree height == 0", BinarySearchTree[int]{}.Height(), 0)

	tree := NewBinarySearchTree[int](10)
	tu.AssertEqualsNamed(t, "Tree with only root node height == 1", tree.Height(), 1)

	left := tree.AppendNode(5)
	tree.AppendNode(20)
	tu.AssertEqualsNamed(t, "Tree with appended nodes height is right", tree.Height(), 2)

	tu.AssertEqualsNamed(t, "AppendNode appends smaller nodes to the left subtree", tree.Root().Left().Data(), 5)
	tu.AssertEqualsNamed(t, "AppendNode appends greater nodes to the right subtree", tree.Root().Right().Data(), 20)

	tree.AppendNode(7)
	tu.AssertEqualsNamed(t, "AppendNode appends subnodes correctly", left.Right().Data(), 7)
	tu.AssertEqualsNamed(t, "Tree with appended subnodes height is right", tree.Height(), 3)

	tu.AssertEqualsNamed(t, "FindNode returns node when it exists", tree.FindNode(7).Data(), 7)
	tree.DeleteNode(7)
	tu.AssertEqualsNamed(t, "FindNode returns nil for non-existing nodes", tree.FindNode(7), nil)

	rootChilds := tree.Root().Children()
	childExpValues := []int{5, 20}
	tu.AssertEqualsNamed(t, "Node children has correct length", len(rootChilds), len(childExpValues))

	for i := range rootChilds {
		tu.AssertEqualsNamed(t, "Tree node returns correct children", rootChilds[i].Data(), childExpValues[i])
	}

	tree.DeleteNode(10)
	tu.AssertEqualsNamed(t, "DeleteNode deletes node with 2 child nodes correctly", tree.Root().Data(), 5)
	tu.AssertEqualsNamed(t, "DeleteNode deletes node with 2 child nodes correctly", tree.Root().Right().Data(), 20)

	tree.DeleteNode(5)
	tu.AssertEqualsNamed(t, "DeleteNode deletes node with 1 child node correctly", tree.Root().Data(), 20)

	tree.DeleteNode(20)
	tu.AssertEqualsNamed(t, "DeleteNode deletes leaf root node correctly", tree.Root(), nil)
	tu.AssertEqualsNamed(t, "DeleteNode deletes leaf root node correctly", tree.Height(), 0)
}
