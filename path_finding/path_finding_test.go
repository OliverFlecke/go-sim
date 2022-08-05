package pathfinding

import (
	sim "simulator/core"
	"testing"
)

func TestFindPath(t *testing.T) {
	world := sim.NewGridWorld(4)
	start := sim.Location{}
	goal := sim.NewLocation(2, 2)

	path, err := FindPath(world, start, goal)

	if err != nil {
		t.Fatal(err.Error())
	}
	if len(path) != 5 {
		t.Logf("World: %s", world.ToStringWithPath(path))
		t.Fatalf("Did not find a shortest path. Length is %d\nPath: %v", len(path), path)
	}
}
