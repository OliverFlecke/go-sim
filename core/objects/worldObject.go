package objects

import simulator "simulator/core"

type WorldObject interface {
	GetLocation() simulator.Location
}

type WorldObjectKey byte

const (
	AGENT = iota
	BOX
	GOAL
)
