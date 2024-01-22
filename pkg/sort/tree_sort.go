package sort

import (
	"cmp"

	"github.com/Wine1y/go-cs/pkg/containers/tree"
)

func TreeSort[T cmp.Ordered](slice []T) {
	if len(slice) == 0 {
		return
	}
	counters := make(map[T]int, len(slice))
	bst := tree.NewBinarySearchTree[T](slice[0])
	counters[slice[0]] = 1

	for i := 1; i < len(slice); i++ {
		bst.AppendNode(slice[i])
		counter, ok := counters[slice[i]]
		if ok {
			counters[slice[i]] = counter + 1
		} else {
			counters[slice[i]] = 1
		}
	}
	iterator := bst.InOrderIterator()
	for i := 0; !iterator.Ended(); {
		value := iterator.Next().Data()
		for inserted := 0; inserted < counters[value]; inserted++ {
			slice[i] = value
			i++
		}
	}
}
