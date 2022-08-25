package action

import (
	"fmt"
	"simulator/core/agent"
	"simulator/core/direction"
	"simulator/core/world"
)

// MoveAction action
type MoveAction struct {
	dir direction.Direction
}

func NewMove(dir direction.Direction) *MoveAction {
	return &MoveAction{
		dir: dir,
	}
}

func (a *MoveAction) GetDirection() direction.Direction {
	return a.dir
}

func (action *MoveAction) IsValid(a *agent.Agent, w world.IWorld) bool {
	newLoc := a.GetLocation().MoveInDirection(action.dir)

	return w.GetLocation(newLoc) == world.EMPTY
}

// IMPL: Action interface

func (m MoveAction) Perform(a *agent.Agent, w world.IWorld) ActionResult {
	if !m.IsValid(a, w) {
		return failure(fmt.Errorf("new location is not free"))
	}

	newLoc := a.GetLocation().MoveInDirection(m.dir)
	w.MoveObject(a, newLoc)

	return success()
}

func (m MoveAction) ToString() string {
	return fmt.Sprintf("Move - dir: %s", m.dir.ToString())
}
