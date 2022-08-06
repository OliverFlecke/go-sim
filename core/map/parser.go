package simulationMap

import (
	"io/ioutil"
	sim "simulator/core"
)

func ParseStringToWorld(text string) sim.IWorld {
	grid := make(map[sim.Location]sim.GridType)
	var x, y int

	for _, c := range text {
		loc := sim.NewLocation(x, y)
		switch c {
		case '\n':
			y += 1
			x = -1
		case ' ':
			grid[loc] = sim.EMPTY
		case '#':
			grid[loc] = sim.WALL
		}
		x += 1
	}

	return (*sim.World)(sim.NewWorld(grid))
}

func GetStringFromFile(filename string) (sim.IWorld, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return ParseStringToWorld(string(content)), nil
}
