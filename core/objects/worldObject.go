package objects

import (
	"simulator/core/location"
)

type WorldObject interface {
	GetLocation() location.Location
	GetRune() rune
}

type WorldObjectKey byte

const (
	AGENT = iota
	BOX
	GOAL
)

type ObjectMap map[WorldObjectKey][]WorldObject
