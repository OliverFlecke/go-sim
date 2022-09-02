package strategy

import (
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"
	"simulator/pathfinding"
)

// This strategy simply finds the goal closest to the given start
// location using BFS
func ClosestGoal(
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

func GoalByDependencies(
	w world.IWorld,
	start location.Location) *objects.Goal {
	mapper := func(
		w world.IWorld,
		l location.Location,
		node *pathfinding.SearchNode) *pathfinding.SearchNode {

		goal := getObjectAtLocationOfType[*objects.Goal](w, l)

		if goal != nil && !w.IsGoalSolved(*goal) {
			newNode := pathfinding.NewSearchNode(*goal)
			node.Children = append(node.Children, newNode)
			return newNode
		} else {
			return node
		}
	}

	tree := pathfinding.FindDepencyTree(w, start, mapper)
	// pathfinding.PrintTree(&tree, 0)

	for _, g := range pathfinding.SearchNodeToList(&tree) {
		return g.(*objects.Goal)
	}

	return nil
}
