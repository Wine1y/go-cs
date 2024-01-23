package search

import (
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func testSearch(t *testing.T, searchFunction SearchingFunction[int]) {
	slice := []int{0, -5, 1, -4, 2, -3, 42, 3, -2, 4, -1, 5, 0}
	i, found := searchFunction(slice, 42)
	tu.AssertEqualsNamed(t, "Search function returns correct found value", found, true)
	tu.AssertEqualsNamed(t, "Search function returns correct index", i, 6)

	i, found = searchFunction(slice, 43)
	tu.AssertEqualsNamed(t, "Search function returns correct found value", found, false)
	tu.AssertEqualsNamed(t, "Search function returns correct index", i, -1)

	edgeSliceL := []int{10, 3, 4, 7, 5, 8, 11}
	iL, foundL := searchFunction(edgeSliceL, 10)
	tu.AssertEqualsNamed(t, "Search function returns correct found value", foundL, true)
	tu.AssertEqualsNamed(t, "Search function returns correct index", iL, 0)

	edgeSliceR := []int{3, 4, 7, 5, 8, 11, 10}
	iR, foundR := searchFunction(edgeSliceR, 10)
	tu.AssertEqualsNamed(t, "Search function returns correct found value", foundR, true)
	tu.AssertEqualsNamed(t, "Search function returns correct index", iR, 6)
}

func testOrderedSearch(t *testing.T, searchFunction SearchingFunction[int]) {
	slice := []int{0, 1, 2, 3, 4, 5, 42, 43, 44, 45, 46, 47, 48}
	i, found := searchFunction(slice, 42)
	tu.AssertEqualsNamed(t, "Search function returns correct found value", found, true)
	tu.AssertEqualsNamed(t, "Search function returns correct index", i, 6)

	i, found = searchFunction(slice, 49)
	tu.AssertEqualsNamed(t, "Search function returns correct found value", found, false)
	tu.AssertEqualsNamed(t, "Search function returns correct index", i, -1)

	edgeSliceL := []int{0, 1, 2, 3, 4, 5, 100, 110, 120, 130, 140, 150}
	iL, foundL := searchFunction(edgeSliceL, 0)
	tu.AssertEqualsNamed(t, "Search function returns correct found value", foundL, true)
	tu.AssertEqualsNamed(t, "Search function returns correct index", iL, 0)

	edgeSliceR := []int{0, 1, 2, 3, 4, 5, 100, 110, 120, 130, 140, 150}
	iR, foundR := searchFunction(edgeSliceR, 150)
	tu.AssertEqualsNamed(t, "Search function returns correct found value", foundR, true)
	tu.AssertEqualsNamed(t, "Search function returns correct index", iR, 11)
}

func TestLinearSearch(t *testing.T) {
	testSearch(t, LinearSearch[int])
}

func TestOrderedLinearSearch(t *testing.T) {
	testOrderedSearch(t, OrderedLinearSearch[int])
}

func TestBinarySearch(t *testing.T) {
	testOrderedSearch(t, BinarySearch[int])
}

func TestInterpolationSearch(t *testing.T) {
	testOrderedSearch(t, InterpolationSearch)
}

func TestJumpSearch(t *testing.T) {
	testOrderedSearch(t, JumpSearch[int])
}
