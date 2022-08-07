package simulator

import (
	"simulator/core/direction"
)

type Action interface {
	Perform(*Agent, *IWorld)
}

type Move struct {
	dir direction.Direction
}

func NewMoveAction(dir direction.Direction) *Move {
	return &Move{
		dir: dir,
	}
}

func (m Move) Perform(a *Agent, w *IWorld) {
	a.MoveInWorld(*w, m.dir)
}
