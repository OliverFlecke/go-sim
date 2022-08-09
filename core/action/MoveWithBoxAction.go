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

func (action MoveWithBoxAction) Perform(a *agent.Agent, w world.IWorld) {
	newLoc := a.GetLocation().MoveInDirection(action.dir)
	err := isValidMoveWithBox(w, newLoc)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.MoveObject(a, newLoc)
	w.MoveObject(action.box, newLoc)
}
