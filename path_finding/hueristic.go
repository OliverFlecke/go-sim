package pathfinding

import (
	sim "simulator/core"
	"simulator/core/location"
)

type heuristic func(int64, location.Location, location.Location, sim.IWorld) int64

func BFS(
	depth int64,
	location location.Location,
	goal location.Location,
	world sim.IWorld) int64 {
	return depth + 1
}

func AStar(depth int64, location location.Location, goal location.Location, world sim.IWorld) int64 {
	return (depth + 1) + int64(location.ManhattanDistance(goal))
}
