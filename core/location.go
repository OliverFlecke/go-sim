package simulator

type Location struct {
	x int
	y int
}

func (loc Location) MoveInDirection(dir Direction) Location {
	result := loc
	switch dir {
	case NORTH:
		result.y += 1
	case SOUTH:
		result.y -= 1
	case EAST:
		result.x += 1
	case WEST:
		result.x -= 1
	}

	return result
}
