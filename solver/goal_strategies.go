package main

import (
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"
	"simulator/pathfinding"
)

// This strategy simply finds the goal closest to the given start
// location using BFS
func closestGoal(
	w world.IWorld,
	start location.Location) *objects.Goal {
	result, _ := pathfinding.FindClosestObject(w, start,
		func(l location.Location) objects.WorldObject {
			for _, x := range w.GetObjectsAtLocation(l) {
				switch g := x.(type) {
				case *objects.Goal:
					if !w.IsGoalSolved(g) {
						return g
					}
				}
			}
			return nil
		})

	return result.(*objects.Goal)
}

func goalByDependencies(
	w world.IWorld,
	start location.Location) *objects.Goal {

	tree := pathfinding.GoalDependencies(w, start)
	for _, g := range pathfinding.GoalTreeToList(&tree) {
		return g
	}

	return nil
}
