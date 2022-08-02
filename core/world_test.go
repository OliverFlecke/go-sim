package simulator

import "testing"

func TestNewGridWorld(t *testing.T) {
	world := NewGridWorld(10)
	expectedSize := 100

	if len(world.grid) != expectedSize {
		t.Fatalf(`World does not have the right size. Expected %d but got %d`, expectedSize, len(world.grid))
	}
}

func TestToString(t *testing.T) {
	world := NewGridWorld(3)
	expected := "   \n   \n   "

	actual := world.ToString()
	if actual != expected {
		t.Fatalf("World does not look the way it should! Actual '%s'", actual)
	}
}
