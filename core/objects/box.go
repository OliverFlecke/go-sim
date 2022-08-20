package objects

import (
	"simulator/core/location"
	"unicode"
)

type Box struct {
	Location location.Location `json:"location"`
	Type     rune              `json:"type"`
}

// IMPL: WorldObject interface

func (b *Box) GetLocation() location.Location {
	return b.Location
}

func (b *Box) SetLocation(l location.Location) {
	b.Location = l
}

func (b *Box) GetRune() rune {
	return unicode.ToUpper(b.Type)
}

// Constructor
func NewBox(location location.Location, boxType rune) *Box {
	return &Box{
		Location: location,
		Type:     boxType,
	}
}

// Getters

func (b Box) GetType() rune {
	return b.Type
}

func (b *Box) Matches(g Goal) bool {
	return unicode.ToLower(b.GetType()) == unicode.ToLower(g.GetType())
}
