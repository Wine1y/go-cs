package stringmatching

import (
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func testStringMatching(t *testing.T, matchingFunction StringMatchingFunction) {
	i, found := matchingFunction("This is a test string!", "test")
	tu.AssertEqualsNamed(t, "Matching function returns correct index", i, 10)
	tu.AssertEqualsNamed(t, "Matching function returns correct found value", found, true)

	i, found = matchingFunction("This is a test string!", "not a substring")
	tu.AssertEqualsNamed(t, "Matching function returns correct index", i, -1)
	tu.AssertEqualsNamed(t, "Matching function returns correct found value", found, false)

	i, found = matchingFunction("0This is a test string!", "0")
	tu.AssertEqualsNamed(t, "Matching function returns correct index", i, 0)
	tu.AssertEqualsNamed(t, "Matching function returns correct found value", found, true)

	i, found = matchingFunction("This is a test string!0", "0")
	tu.AssertEqualsNamed(t, "Matching function returns correct index", i, 22)
	tu.AssertEqualsNamed(t, "Matching function returns correct found value", found, true)
}

func TestBruteForceMatching(t *testing.T) {
	testStringMatching(t, BruteForceMatching)
}

func TestRabinKarpMatching(t *testing.T) {
	testStringMatching(t, RabinKarpMatching)
}

func TestUnknownMatching(t *testing.T) {
	testStringMatching(t, UnknownMatching)
}

func TestKnuthMorrisPrattMatching(t *testing.T) {
	testStringMatching(t, KnuthMorrisPrattMatching)
}
