package objects

import "simulator/core/location"

type Goal struct {
	Location location.Location `json:"location"`
	Type     rune              `json:"type"`
}

// IMPL: WorldObject interface

func (g *Goal) GetLocation() location.Location {
	return g.Location
}

func (g *Goal) SetLocation(l location.Location) {
	g.Location = l
}

func (g *Goal) GetRune() rune {
	return g.Type
}

// CONSTRUCTORS

func NewGoal(location location.Location, goalType rune) *Goal {
	return &Goal{
		Location: location,
		Type:     goalType,
	}
}

func (g *Goal) GetType() rune {
	return g.Type
}
