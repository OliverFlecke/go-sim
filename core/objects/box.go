package objects

import "simulator/core/location"

type Box struct {
	location location.Location
	boxType  rune
}

// IMPL: WorldObject interface

func (b Box) GetLocation() location.Location {
	return b.location
}

// Constructor
func NewBox(location location.Location, boxType rune) *Box {
	return &Box{
		location: location,
		boxType:  boxType,
	}
}

// Getters

func (b Box) GetType() rune {
	return b.boxType
}
