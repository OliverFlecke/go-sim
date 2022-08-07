package simulator

import (
	"simulator/core/direction"
	"simulator/core/objects"
)

type Action interface {
	Perform(*Agent, *IWorld)
}

// MoveAction action
type MoveAction struct {
	dir direction.Direction
}

func NewActionMove(dir direction.Direction) *MoveAction {
	return &MoveAction{
		dir: dir,
	}
}

func (m MoveAction) Perform(a *Agent, w *IWorld) {
	a.MoveInWorld(*w, m.dir)
}

// Pickup action
type MoveWithBoxAction struct {
	dir direction.Direction
	box *objects.Box
}

func NewActionMoveWithBox(
	dir direction.Direction,
	box *objects.Box) *MoveWithBoxAction {
	return &MoveWithBoxAction{
		dir: dir,
		box: box,
	}
}

func (action MoveWithBoxAction) Perform(a *Agent, w *IWorld) {
	a.MoveInWorld(*w, action.dir)
	(*w).MoveObject(action.box, action.box.GetLocation().MoveInDirection(action.dir))
}
