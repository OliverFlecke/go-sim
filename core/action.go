package simulator

import (
	"fmt"
	"simulator/core/direction"
	"simulator/core/location"
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

func isValidMoveWithBox(newL location.Location, w *IWorld) error {
	for _, v := range (*w).GetObjectsAtLocation(newL) {
		switch o := v.(type) {
		case *objects.Box:
			return fmt.Errorf("a box is already at loc (%v): %v", newL, o)
		case *Agent:
			return fmt.Errorf("an agent is already at loc (%v): %v", newL, o)
		}
	}

	return nil
}

func (action MoveWithBoxAction) Perform(a *Agent, w *IWorld) {
	err := isValidMoveWithBox(a.GetLocation().MoveInDirection(action.dir), w)
	if err != nil {
		fmt.Println(err)
		return
	}

	a.MoveInWorld(*w, action.dir)
	(*w).MoveObject(action.box, action.box.GetLocation().MoveInDirection(action.dir))
}
