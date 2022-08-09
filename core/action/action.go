package action

import (
	"fmt"
	"simulator/core/agent"
	"simulator/core/direction"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"
)

type Action interface {
	Perform(*agent.Agent, *world.IWorld)
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

func isValidMove(w world.IWorld, agent *agent.Agent, dir direction.Direction) bool {
	newLocation := agent.GetLocation().MoveInDirection(dir)
	return w.GetLocation(newLocation) == world.EMPTY
}

func (m MoveAction) Perform(a *agent.Agent, w *world.IWorld) {
	if !isValidMove(*w, a, m.dir) {
		return
	}
	a.SetLocation(a.GetLocation().MoveInDirection(m.dir))
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

func isValidMoveWithBox(newL location.Location, w *world.IWorld) error {
	for _, v := range (*w).GetObjectsAtLocation(newL) {
		switch o := v.(type) {
		case *objects.Box:
			return fmt.Errorf("a box is already at loc (%v): %v", newL, o)
		case *agent.Agent:
			return fmt.Errorf("an agent is already at loc (%v): %v", newL, o)
		}
	}

	return nil
}

func (action MoveWithBoxAction) Perform(a *agent.Agent, w *world.IWorld) {
	err := isValidMoveWithBox(a.GetLocation().MoveInDirection(action.dir), w)
	if err != nil {
		fmt.Println(err)
		return
	}

	a.SetLocation(a.GetLocation().MoveInDirection(action.dir))
	(*w).MoveObject(action.box, action.box.GetLocation().MoveInDirection(action.dir))
}
