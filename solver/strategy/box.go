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
	goal *objects.Goal) []Mission {
	mapper := func(
		w world.IWorld,
		l location.Location,
		node *pathfinding.SearchNode) (*pathfinding.SearchNode, bool) {

		box := getObjectAtLocationOfType[*objects.Box](w, l)

		if box != nil {
			newNode := pathfinding.NewSearchNode(*box)
			node.Children = append(node.Children, newNode)

			return newNode, (*box).Matches(*goal)
		} else {
			return node, false
		}
	}

	tree := pathfinding.FindDepencyTree(w, goal.GetLocation(), mapper)
	pathfinding.PrintTree(&tree, 0)

	return toMissions(&tree, goal)
}

type Mission interface {
}

type StoreBoxMission struct {
	Box *objects.Box
}

type MoveBoxToGoalMission struct {
	Box  *objects.Box
	Goal *objects.Goal
}

func toMissions(
	node *pathfinding.SearchNode,
	goal *objects.Goal) []Mission {
	missions := make([]Mission, 0)

	queue := make([]*pathfinding.SearchNode, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, child := range current.Children {
			var m Mission
			if len(child.Children) == 0 {
				m = MoveBoxToGoalMission{
					Box:  child.Item.(*objects.Box),
					Goal: goal,
				}
			} else {
				m = StoreBoxMission{
					Box: child.Item.(*objects.Box),
				}
			}
			missions = append(missions, m)
		}

		queue = append(queue, current.Children...)
	}

	return missions
}

func FindNearestBox(w world.IWorld, goal *objects.Goal) (*objects.Box, error) {
	// fmt.Printf("Searching for boxes\n")
	// GetBoxWithDependencies(w, goal)

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
