package simulator

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
