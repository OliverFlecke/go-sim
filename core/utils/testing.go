package utils

import (
	"reflect"
	"testing"
)

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Received %v (type %v), expected %v (type %v)",
			a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

func AssertEqualSlices[T comparable](t *testing.T, actual []T, expected []T) {
	if len(actual) != len(expected) {
		t.Fatalf("Slices has different length. a: %d, b: %d", len(actual), len(expected))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("Item at %d does not match. a: %v, b: %v", i, actual[i], expected[i])
		}
	}
}
