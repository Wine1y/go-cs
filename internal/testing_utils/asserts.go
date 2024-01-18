package testingutils

import "testing"

func Assert(t *testing.T, name string, condition bool) {
	t.Helper()
	if !condition {
		t.Fatalf("Expected: %v", name)
	}
}

func AssertEquals[T comparable](t *testing.T, real, expected T) {
	t.Helper()
	if expected != real {
		t.Fatalf("Expected: %v, Actual: %v\n", expected, real)
	}
}

func AssertEqualsNamed[T comparable](t *testing.T, name string, real, expected T) {
	t.Helper()
	if expected != real {
		t.Fatalf("Expected %s: %v, Actual: %v\n", name, expected, real)
	}
}

func AssertPanics(t *testing.T, name string, function func()) {
	t.Helper()
	defer func() { _ = recover() }()
	function()
	t.Fatalf("Expected to panic: %s\n", name)
}
