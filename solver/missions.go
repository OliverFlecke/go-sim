package main

import (
	simulator "simulator/core"
	"simulator/core/agent"
	"simulator/core/objects"
	"simulator/pathfinding"
)

func StoreBox(
	sim *simulator.Simulation,
	a *agent.Agent,
	box *objects.Box) {
	storage, _ := pathfinding.FindStorageLocation(
		sim.GetWorld(), box.GetLocation())

	locs, _, _ := pathfinding.FindPath(
		sim.GetWorld(),
		a.GetLocation(),
		box.GetLocation(),
		pathfinding.AStar, nil)
	acts := LocationsToActions(locs)
	sim.SetActions(a, acts)

	locs, _, _ = FindPathWhileCarringBox(
		sim.GetWorld(),
		box,
		*storage)
	sim.SetActions(a, LocationsToMoveWithMoxActions(locs, box))
}
