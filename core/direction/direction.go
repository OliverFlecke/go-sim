package direction

type Direction byte

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func (d Direction) ToString() string {
	switch d {
	case NORTH:
		return "N"
	case EAST:
		return "E"
	case SOUTH:
		return "S"
	case WEST:
		return "W"
	default:
		return "Unknown direction"
	}
}
