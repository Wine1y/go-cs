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
