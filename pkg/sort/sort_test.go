package sort

import (
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func testInPlaceSorting(t *testing.T, fn InPlaceSortingFunction[int]) {
	slice := []int{2, 4, 1, 3, 5, 1, -2, -5}
	sortedSlice := []int{-5, -2, 1, 1, 2, 3, 4, 5}
	fn(slice)
	for i := 0; i < len(slice); i++ {
		tu.AssertEqualsNamed(t, "Slice is sorted in ascending order", slice[i], sortedSlice[i])
	}
}

func testSorting(t *testing.T, fn SortingFunction[int]) {
	slice := []int{2, 4, 1, 3, 5, 1, -2, -5}
	sortedSlice := []int{-5, -2, 1, 1, 2, 3, 4, 5}
	slice = fn(slice)
	for i := 0; i < len(slice); i++ {
		tu.AssertEqualsNamed(t, "Slice is sorted in ascending order", slice[i], sortedSlice[i])
	}
}

func TestBubbleSort(t *testing.T) {
	testInPlaceSorting(t, BubbleSort[int])
}

func TestSelectionSort(t *testing.T) {
	testInPlaceSorting(t, SelectionSort[int])
}

func TestInsertionSort(t *testing.T) {
	testInPlaceSorting(t, InsertionSort[int])
}
