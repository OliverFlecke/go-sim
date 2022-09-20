package main

import (
	"simulator/core/action"
	"simulator/core/direction"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/utils"
	"simulator/core/world"
	"simulator/pathfinding"
)

func LocationsToActions(locs []location.Location) []action.Action {
	return utils.Mapi(location.PathToDirections(locs), func(_ int, dir direction.Direction) action.Action {
		return action.NewMove(dir)
	})
}

func LocationsToMoveWithMoxActions(
	locs []location.Location,
	box *objects.Box) []action.Action {
	return utils.Mapi(location.PathToDirections(locs),
		func(_ int, dir direction.Direction) action.Action {
			return action.NewMoveWithBox(dir, box)
		})
}

func FindPathWhileCarringBox(
	w world.IWorld,
	box *objects.Box,
	end location.Location) ([]location.Location, pathfinding.SearchStats, error) {
	return pathfinding.FindPath(
		w,
		box.GetLocation(),
		end,
		pathfinding.AStar,
		func(l location.Location, w world.IWorld) bool {
			for _, o := range w.GetObjectsAtLocation(l) {
				switch o.(type) {
				case *objects.Box:
					return false
				}
			}

			return true
		})
}
