package objects

import "simulator/core/location"

type Goal struct {
	location location.Location
}

// IMPL: WorldObject interface

func (g Goal) GetLocation() location.Location {
	return g.location
}

// CONSTRUCTORS

func NewGoal(location location.Location) *Goal {
	return &Goal{
		location: location,
	}
}
