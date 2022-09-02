package strategy

import (
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"
	"simulator/pathfinding"
)

func getObjectAtLocationOfType[T objects.WorldObject](
	w world.IWorld,
	l location.Location) *T {
	for _, o := range w.GetObjectsAtLocation(l) {
		switch g := o.(type) {
		case T:
			return &g
		}
	}

	return nil
}

func GetBoxWithDependencies(
	w world.IWorld,
	start location.Location) {
	mapper := func(
		w world.IWorld,
		l location.Location,
		node *pathfinding.SearchNode) *pathfinding.SearchNode {

		box := getObjectAtLocationOfType[*objects.Box](w, l)

		if box != nil {
			newNode := pathfinding.NewSearchNode(*box)
			node.Children = append(node.Children, newNode)
			return newNode
		} else {
			return node
		}
	}

	tree := pathfinding.FindDepencyTree(w, start, mapper)
	pathfinding.PrintTree(&tree, 0)
}

func FindNearestBox(w world.IWorld, goal *objects.Goal) (*objects.Box, error) {
	// fmt.Printf("Searching for boxes\n")
	// GetBoxWithDependencies(w, goal.Location)

	obj, err := pathfinding.FindClosestObject(w, goal.GetLocation(),
		func(l location.Location) objects.WorldObject {
			var box *objects.Box = nil
			var otherGoal *objects.Goal

			for _, obj := range w.GetObjectsAtLocation(l) {
				switch v := obj.(type) {
				case *objects.Box:
					if v.Matches(*goal) {
						box = v
					}
				case *objects.Goal:
					otherGoal = v
				}
			}

			if box == nil || (otherGoal != nil && box.Matches(*otherGoal)) {
				return nil
			}

			return box
		})

	if err != nil {
		return nil, err
	}

	return obj.(*objects.Box), nil
}
