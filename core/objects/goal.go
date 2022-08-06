package objects

import sim "simulator/core"

type Goal struct {
	location sim.Location
}

// IMPL: WorldObject interface

func (g Goal) GetLocation() sim.Location {
	return g.location
}

// CONSTRUCTORS

func NewGoal(location sim.Location) *Goal {
	return &Goal{
		location: location,
	}
}
