package sort

import (
	"slices"
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func getDefaultSets() [][]int {
	return [][]int{
		{2, 4, 1, 3, 5, 1, -2, -5},
	}
}

func testInPlaceSorting(t *testing.T, fn InPlaceSortingFunction[int], testingSets ...[]int) {
	for _, set := range testingSets {
		sorted := make([]int, len(set))
		copy(sorted, set)

		slices.Sort(sorted)
		fn(set)

		for i := 0; i < len(set); i++ {
			tu.AssertEqualsNamed(t, "Slice is sorted in ascending order", set[i], sorted[i])
		}
	}
}

func testSorting(t *testing.T, fn SortingFunction[int], testingSets ...[]int) {
	for _, set := range testingSets {
		sorted := make([]int, len(set))
		copy(sorted, set)

		slices.Sort(sorted)
		set = fn(set)

		for i := 0; i < len(set); i++ {
			tu.AssertEqualsNamed(t, "Slice is sorted in ascending order", set[i], sorted[i])
		}
	}
}

func TestBubbleSort(t *testing.T) {
	testInPlaceSorting(t, BubbleSort[int], getDefaultSets()...)
}

func TestSelectionSort(t *testing.T) {
	testInPlaceSorting(t, SelectionSort[int], getDefaultSets()...)
}

func TestInsertionSort(t *testing.T) {
	testInPlaceSorting(t, InsertionSort[int], getDefaultSets()...)
}

func TestShellSort(t *testing.T) {
	testInPlaceSorting(t, ShellSort[int], getDefaultSets()...)
}

func TestMergeSort(t *testing.T) {
	testSorting(t, MergeSort[int], getDefaultSets()...)
}

func TestQuickSort(t *testing.T) {
	testInPlaceSorting(t, QuickSort[int], getDefaultSets()...)
}

func TestHeapSort(t *testing.T) {
	testInPlaceSorting(t, HeapSort[int], getDefaultSets()...)
}

func TestTreeSort(t *testing.T) {
	testInPlaceSorting(t, TreeSort[int], getDefaultSets()...)
}

func TestCountingSort(t *testing.T) {
	testSorting(t, CountingSort, []int{2, 4, 1, 3, 5, 1, 2, 10})
}

func TestRadixSort(t *testing.T) {
	testSorting(t, CountingSort, []int{2, 4, 1, 3, 5, 1, 2, 10})
}
