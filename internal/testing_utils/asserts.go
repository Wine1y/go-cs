package testingutils

import "testing"

func AssertEquals[T comparable](t *testing.T, real, expected T) {
	if expected != real {
		t.Fatalf("Expected: %v, Actual: %v\n", expected, real)
	}
}

func AssertEqualsNamed[T comparable](t *testing.T, name string, real, expected T) {
	if expected != real {
		t.Fatalf("Expected %s: %v, Actual: %v\n", name, expected, real)
	}
}
