package objects

import (
	"simulator/core/location"
	"unicode"
)

type Box struct {
	location location.Location
	boxType  rune
}

// IMPL: WorldObject interface

func (b *Box) GetLocation() location.Location {
	return b.location
}

func (b *Box) SetLocation(l location.Location) {
	b.location = l
}

func (b *Box) GetRune() rune {
	return unicode.ToUpper(b.boxType)
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

func (b *Box) Matches(g Goal) bool {
	return unicode.ToLower(b.GetType()) == unicode.ToLower(g.GetType())
}
