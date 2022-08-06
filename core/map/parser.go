package simulationMap

import (
	"fmt"
	"io/ioutil"
	"regexp"
	sim "simulator/core"
	"simulator/core/objects"
	obj "simulator/core/objects"
	"strconv"
	"strings"
)

func parseGridWorld(text string) sim.Grid {
	grid := make(sim.Grid)
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

	return grid
}

func parseInt(str string) (int, error) {
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

func parseLocation(re *regexp.Regexp, match []string) (sim.Location, error) {
	x, err := parseInt(match[re.SubexpIndex("x")])
	if err != nil {
		return sim.Location{}, err
	}
	y, err := parseInt(match[re.SubexpIndex("y")])
	if err != nil {
		return sim.Location{}, err
	}

	return sim.NewLocation(x, y), nil
}

type ObjectMap map[objects.WorldObjectKey][]objects.WorldObject

func parseObjects(str string) (ObjectMap, error) {
	result := make(ObjectMap)

	re := regexp.MustCompile(`(?P<type>[a-z]+) (?P<id>\w) (?P<x>\d+),(?P<y>\d+)`)
	typeIdx := re.SubexpIndex("type")
	for i, line := range strings.Split(str, "\n") {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			callsign := rune(match[re.SubexpIndex("id")][0])
			loc, err := parseLocation(re, match)
			if err != nil {
				return nil, fmt.Errorf("unable to parse location on line %d with %s", i, line)
			}

			switch match[typeIdx] {
			case "agent":
				agent := sim.NewAgentWithStartLocation("unused", callsign, loc)
				result[obj.AGENT] = append(result[obj.AGENT], agent)
			case "box":
				box := obj.NewBox(loc, callsign)
				result[obj.BOX] = append(result[obj.BOX], box)
			case "goal":
				goal := obj.NewGoal(loc)
				result[obj.GOAL] = append(result[obj.GOAL], goal)
			default:
				return nil, fmt.Errorf("unable to handle type: '%s'", match[typeIdx])
			}
		}
	}

	return result, nil
}

func parseWorldFromString(content string) (sim.IWorld, error) {
	splits := strings.Split(content, "\n\n")

	grid := parseGridWorld(splits[0])
	if len(splits) > 1 {
		// fmt.Println(splits[1])
		parseObjects(splits[1])
	}

	return sim.NewWorld(grid), nil
}

func ParseWorldFromFile(filename string) (sim.IWorld, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return parseWorldFromString(string(content))
}
