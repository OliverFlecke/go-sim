package pathfinding

import (
	sim "simulator/core"
	"simulator/core/location"
	"testing"
)

func TestFindPath(t *testing.T) {
	world := sim.NewGridWorld(4)
	start := location.Location{}
	goal := location.NewLocation(2, 2)

	path, stats, err := FindPath(world, start, goal, BFS)

	if err != nil {
		t.Fatal(err.Error())
	}
	if len(path) != 5 {
		t.Logf("World:\n%s", world.ToStringWithPath(path))
		t.Fatalf("Did not find a shortest path. Length is %d\nPath: %v", len(path), path)
	}

	if stats.visited > 13 {
		t.Fatalf("Too many states where visited %d", stats.visited)
	}
}

func TestFindPathWithAStar(t *testing.T) {
	world := sim.NewGridWorld(10)
	start := location.Location{}
	goal := location.NewLocation(9, 9)

	path, _, err := FindPath(world, start, goal, AStar)

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("World:\n%s", world.ToStringWithPath(path))
	if len(path) != 19 {
		t.Fatalf("Did not find a shortest path. Length is %d\nPath: %v", len(path), path)
	}
}
