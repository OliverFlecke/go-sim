package pathfinding

import (
	"simulator/core/location"
	"simulator/core/world"
)

type filter func(location.Location, world.IWorld) bool
