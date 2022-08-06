package simulator

type GridType = byte

const (
	EMPTY GridType = iota
	WALL
)

func ToRune(g GridType) rune {
	switch g {
	case WALL:
		return '#'
	default:
		return ' '
	}
}
