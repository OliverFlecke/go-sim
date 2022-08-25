package mapping

import (
	"simulator/core/direction"
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
