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

func TestToStringWithAgents(t *testing.T) {
	world := NewGridWorld(3)
	agents := []Agent{
		*NewAgent("Test agent"),
	}
	agents[0].location = Location{x: 1, y: 1}
	expected := "   \n 0 \n   "

	actual := world.ToStringWithAgents(agents)

	if actual != expected {
		t.Fatalf(`World with agents does not look as expected '%s'. Actual: '%s'`, expected, actual)
	}
}
