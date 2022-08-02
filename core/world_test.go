package simulator

import "testing"

func TestNewGridWorld(t *testing.T) {
	world := NewGridWorld(10)
	expectedSize := 100

	if len(world.grid) != expectedSize {
		t.Fatalf(`World does not have the right size. Expected %d but got %d`, expectedSize, len(world.grid))
	}
}
