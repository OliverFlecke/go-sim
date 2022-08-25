package mapping

import (
	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/direction"
	"simulator/core/objects"
	"simulator/core/utils"
	"simulator/model/dto"
)

func GetActions(acts *dto.ActionList, sim *simulator.Simulation) []action.Action {
	return utils.Mapi(acts.Actions, func(_ int, d *dto.Action) action.Action {
		switch a := d.Action.(type) {
		case *dto.Action_Move:
			return toMoveAction(a)
		case *dto.Action_MoveWithBox:
			var box *objects.Box
			for _, x := range sim.GetWorld().GetObjects(objects.BOX) {
				b := x.(*objects.Box)
				if b.GetId() == a.MoveWithBox.BoxId {
					box = b
					break
				}
			}
			return action.NewMoveWithBox(direction.Direction(a.MoveWithBox.Direction), box)

		default:
			// Should never be hit
			return nil
		}
	})
}

func toMoveAction(a *dto.Action_Move) *action.MoveAction {
	return action.NewMove(direction.Direction(a.Move.Direction))
}

func ToDtos(acts []action.Action) *dto.ActionList {
	return &dto.ActionList{
		Actions: utils.Mapi(acts, func(_ int, in action.Action) *dto.Action {
			switch a := in.(type) {
			case *action.MoveAction:
				return &dto.Action{
					Action: &dto.Action_Move{
						Move: &dto.Move{
							Direction: dto.Direction(a.GetDirection()),
						},
					},
				}
			case *action.MoveWithBoxAction:
				return &dto.Action{
					Action: &dto.Action_MoveWithBox{
						MoveWithBox: &dto.MoveWithBox{
							Direction: dto.Direction(a.GetDirection()),
							BoxId:     a.GetBoxId(),
						},
					},
				}
			default:
				return nil
			}
		}),
	}
}
