package objects

import (
	"simulator/core/location"
	"unicode"
)

type Box struct {
	id       uint32
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

var boxCount uint32 = 0

// Constructor
func NewBox(location location.Location, boxType rune) *Box {
	id := boxCount
	boxCount += 1
	return NewBoxWithId(id, location, boxType)
}

func NewBoxWithId(id uint32, l location.Location, t rune) *Box {
	return &Box{
		id:       id,
		Location: l,
		Type:     t,
	}
}

// Getters

func (b Box) GetId() uint32 {
	return b.id
}

func (b Box) GetType() rune {
	return b.Type
}

func (b *Box) Matches(g Goal) bool {
	return unicode.ToLower(b.GetType()) == unicode.ToLower(g.GetType())
}
