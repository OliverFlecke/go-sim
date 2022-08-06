package objects

import sim "simulator/core"

type Box struct {
	location sim.Location
	boxType  rune
}

// IMPL: WorldObject interface

func (b Box) GetLocation() sim.Location {
	return b.location
}

// Constructor
func NewBox(location sim.Location, boxType rune) *Box {
	return &Box{
		location: location,
		boxType:  boxType,
	}
}

// Getters

func (b Box) GetType() rune {
	return b.boxType
}
