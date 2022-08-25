package mapping

import (
	"simulator/core/action"
	"simulator/core/direction"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/model/dto"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestDtoToMoveAction(t *testing.T) {
	dtoDirs := []dto.Direction{
		dto.Direction_NORTH,
		dto.Direction_EAST,
		dto.Direction_SOUTH,
		dto.Direction_WEST,
	}
	expected := []direction.Direction{
		direction.NORTH,
		direction.EAST,
		direction.SOUTH,
		direction.WEST,
	}

	for i, dir := range dtoDirs {
		d := &dto.Action_Move{
			Move: &dto.Move{
				Direction: dir,
			},
		}

		a := toMoveAction(d)
		Equal(t, a.GetDirection(), expected[i])
	}
}

func TestMoveActionToDto(t *testing.T) {
	acts := []action.Action{
		action.NewMove(direction.NORTH),
		action.NewMove(direction.EAST),
		action.NewMove(direction.SOUTH),
		action.NewMove(direction.WEST),
	}
	expectedDirections := []dto.Direction{
		dto.Direction_NORTH,
		dto.Direction_EAST,
		dto.Direction_SOUTH,
		dto.Direction_WEST,
	}

	converted := ToDtos(acts)

	for i, c := range converted.Actions {
		IsType(t, &dto.Action_Move{}, c.GetAction())
		Equal(t, expectedDirections[i], c.GetMove().Direction)
	}
}

func TestMoveWithBoxActionToDto(t *testing.T) {
	l := location.New(0, 0)
	acts := []action.Action{
		action.NewMoveWithBox(direction.NORTH, objects.NewBox(l, 'a')),
		action.NewMoveWithBox(direction.EAST, objects.NewBox(l, 'a')),
		action.NewMoveWithBox(direction.SOUTH, objects.NewBox(l, 'a')),
		action.NewMoveWithBox(direction.WEST, objects.NewBox(l, 'a')),
	}
	expectedDirections := []dto.Direction{
		dto.Direction_NORTH,
		dto.Direction_EAST,
		dto.Direction_SOUTH,
		dto.Direction_WEST,
	}

	converted := ToDtos(acts)

	for i, c := range converted.Actions {
		IsType(t, &dto.Action_MoveWithBox{}, c.GetAction())
		Equal(t, expectedDirections[i], c.GetMoveWithBox().Direction)
		Equal(t, uint32(i), c.GetMoveWithBox().BoxId)
	}
}
