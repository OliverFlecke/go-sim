package action

import (
	"fmt"
	"simulator/core/agent"
	"simulator/core/direction"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"
)

// Pickup action
type MoveWithBoxAction struct {
	dir direction.Direction
	box *objects.Box
}

func NewMoveWithBox(
	dir direction.Direction,
	box *objects.Box) *MoveWithBoxAction {
	return &MoveWithBoxAction{
		dir: dir,
		box: box,
	}
}

func isValidMoveWithBox(w world.IWorld, newL location.Location) error {
	if !isValidMove(w, newL) {
		return fmt.Errorf("agent cannot move here")
	}

	for _, v := range w.GetObjectsAtLocation(newL) {
		switch o := v.(type) {
		case *objects.Box:
			return fmt.Errorf("a box is already at loc (%v): %v", newL, o)
		case *agent.Agent:
			return fmt.Errorf("an agent is already at loc (%v): %v", newL, o)
		}
	}

	return nil
}

// IMPL: Action interface

func (action MoveWithBoxAction) Perform(a *agent.Agent, w world.IWorld) ActionResult {
	newLoc := a.GetLocation().MoveInDirection(action.dir)
	err := isValidMoveWithBox(w, newLoc)
	if err != nil {
		return failure(err)
	}

	w.MoveObject(a, newLoc)
	w.MoveObject(action.box, newLoc)

	return success()
}

func (a MoveWithBoxAction) ToString() string {
	return fmt.Sprintf("MoveWithBox - location %v, dir %s", a.box.GetLocation(), a.dir.ToString())
}
