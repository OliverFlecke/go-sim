package pathfinding

import sim "simulator/core"

type heuristic func(int64, sim.Location, sim.Location, *sim.IWorld) int64

func BFS(
	depth int64,
	location sim.Location,
	goal sim.Location,
	world *sim.IWorld) int64 {
	return depth + 1
}

func AStar(depth int64, location sim.Location, goal sim.Location, world *sim.IWorld) int64 {
	return (depth + 1) + int64(location.ManhattanDistance(goal))
}
