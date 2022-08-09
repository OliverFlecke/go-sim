package action

import (
	"fmt"
	"simulator/core/agent"
	"simulator/core/direction"
	"simulator/core/location"
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

func isValidMove(w world.IWorld, newLoc location.Location) bool {
	return w.GetLocation(newLoc) == world.EMPTY
}

func (m MoveAction) Perform(a *agent.Agent, w world.IWorld) ActionResult {
	newLoc := a.GetLocation().MoveInDirection(m.dir)
	if !isValidMove(w, newLoc) {
		return failure(fmt.Errorf("new location is not free"))
	}
	a.SetLocation(newLoc)

	return success()
}
