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

func AssertEqualSlices[T comparable](t *testing.T, a []T, b []T) {
	if len(a) != len(b) {
		t.Fatalf("Slices has different length. a: %d, b: %d", len(a), len(b))
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Fatalf("Item at %d does not match. a: %v, b: %v", i, a[i], b[i])
		}
	}
}
