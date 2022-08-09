package pathfinding

import (
	sim "simulator/core"
	"simulator/core/location"
)

type filter func(location.Location, sim.IWorld) bool
