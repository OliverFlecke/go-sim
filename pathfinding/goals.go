package pathfinding

import (
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"
)

func GoalDependencies(w world.IWorld, start location.Location) SearchNode {
	mapper := func(
		w world.IWorld,
		l location.Location,
		node *SearchNode) *SearchNode {

		goal := getGoalAtLocation(w, l)

		if goal != nil && !w.IsGoalSolved(goal) {
			newNode := newSearchNode(goal)
			node.Children = append(node.Children, newNode)
			return newNode
		} else {
			return node
		}
	}

	return FindDepencyTree(w, start, mapper)
}

func getGoalAtLocation(w world.IWorld, l location.Location) *objects.Goal {
	for _, o := range w.GetObjectsAtLocation(l) {
		switch g := o.(type) {
		case *objects.Goal:
			if !w.IsGoalSolved(g) {
				return g
			}
		}
	}

	return nil
}
