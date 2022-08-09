package pathfinding

import (
	"simulator/core/location"
	"simulator/core/world"
)

type heuristic func(int64, location.Location, location.Location, world.IWorld) int64

func BFS(
	depth int64,
	location location.Location,
	goal location.Location,
	world world.IWorld) int64 {
	return depth + 1
}

func AStar(depth int64, location location.Location, goal location.Location, world world.IWorld) int64 {
	return (depth + 1) + int64(location.ManhattanDistance(goal))
}
