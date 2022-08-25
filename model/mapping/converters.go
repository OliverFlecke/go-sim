package mapping

import (
	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/direction"
	"simulator/core/utils"
	"simulator/model/dto"
)

func GetActions(acts *dto.ActionList, sim *simulator.Simulation) []action.Action {
	return utils.Mapi(acts.Actions, func(_ int, d *dto.Action) action.Action {
		switch a := d.Action.(type) {
		case *dto.Action_Move:
			return toMoveAction(a)
		case *dto.Action_MoveWithBox:
			box := sim.GetWorld().GetBoxes()[0]
			return action.NewMoveWithBox(direction.Direction(a.MoveWithBox.Direction), &box)

		default:
			// Should never be hit
			return nil
		}
	})
}

func toMoveAction(a *dto.Action_Move) *action.MoveAction {
	return action.NewMove(direction.Direction(a.Move.Direction))
}
