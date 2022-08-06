package objects

import "simulator/core/location"

type Goal struct {
	location location.Location
	goalType rune
}

// IMPL: WorldObject interface

func (g Goal) GetLocation() location.Location {
	return g.location
}

func (g Goal) GetRune() rune {
	return g.goalType
}

// CONSTRUCTORS

func NewGoal(location location.Location, goalType rune) *Goal {
	return &Goal{
		location: location,
		goalType: goalType,
	}
}
