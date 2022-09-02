package strategy

import (
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"
	"simulator/pathfinding"
)

func GetBoxWithDependencies() {

}

func FindNearestBox(w world.IWorld, goal *objects.Goal) (*objects.Box, error) {
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
