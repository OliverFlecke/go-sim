package simulator

import (
	"simulator/core/agent"
	"simulator/core/location"
	"testing"
)

func TestNewGridWorld(t *testing.T) {
	world := NewGridWorld(10)
	expectedSize := 144

	if len(world.grid) != expectedSize {
		t.Fatalf(`World does not have the right size. Expected %d but got %d`, expectedSize, len(world.grid))
	}
}

func TestToString(t *testing.T) {
	world := NewGridWorld(3)
	expected := "#####\n#   #\n#   #\n#   #\n#####"

	actual := world.ToString()
	if actual != expected {
		t.Fatalf("World does not look the way it should! Actual '%s'", actual)
	}
}

func TestToStringWithAgents(t *testing.T) {
	world := NewGridWorld(3)
	agents := []agent.Agent{
		*agent.NewAgentWithStartLocation("Test agent", '0', location.New(2, 2)),
	}
	expected := "#####\n#   #\n# 0 #\n#   #\n#####"

	actual := world.ToStringWithAgents(agents)

	if actual != expected {
		t.Fatalf(`World with agents does not look as expected '%s'. Actual: '%s'`, expected, actual)
	}
}

func TestWorldImplementsIWord(t *testing.T) {
	var w IWorld = (*World)(NewGridWorld(1)) // Verify that *T implements I.
	if w.GetLocation(location.New(1, 1)) != EMPTY {
		t.Fatal("Incorrect location returned from interface")
	}
}
