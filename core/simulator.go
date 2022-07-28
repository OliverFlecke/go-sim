package simulator

type Direction byte

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type Location struct {
	x int64
	y int64
}
