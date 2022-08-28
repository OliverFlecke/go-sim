package level

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"simulator/core/agent"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"
	"strconv"
	"strings"
)

func parseGridWorld(text string) world.Grid {
	grid := make(world.Grid)
	var x, y int

	for _, c := range text {
		loc := location.New(x, y)
		switch c {
		case '\n':
			y += 1
			x = -1
		case ' ':
			// Don't store empty locations. Grid should be represented by a sparse matrix
			// grid[loc] = world.EMPTY
		case '#':
			grid[loc] = world.WALL
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

func parseLocation(re *regexp.Regexp, match []string) (location.Location, error) {
	x, err := parseInt(match[re.SubexpIndex("x")])
	if err != nil {
		return location.Location{}, err
	}
	y, err := parseInt(match[re.SubexpIndex("y")])
	if err != nil {
		return location.Location{}, err
	}

	return location.New(x, y), nil
}

func parseObjects(str string) (objects.ObjectMap, error) {
	result := make(objects.ObjectMap)

	counts := make(map[objects.WorldObjectKey]uint32)
	counts[objects.AGENT] = 0
	counts[objects.BOX] = 0
	counts[objects.GOAL] = 0

	re := regexp.MustCompile(`(?P<type>[a-z]+) (?P<id>\w) (?P<x>\d+),(?P<y>\d+)`)
	typeIdx := re.SubexpIndex("type")
	for i, line := range strings.Split(str, "\n") {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			r := rune(match[re.SubexpIndex("id")][0])
			loc, err := parseLocation(re, match)
			if err != nil {
				return nil, fmt.Errorf("unable to parse location on line %d with %s", i, line)
			}

			switch match[typeIdx] {
			case "agent":
				id := counts[objects.AGENT]
				counts[objects.AGENT] += 1
				agent := agent.NewAgent(id, r, loc)
				result[objects.AGENT] = append(result[objects.AGENT], agent)
			case "box":
				id := counts[objects.BOX]
				counts[objects.BOX] += 1
				box := objects.NewBoxWithId(id, loc, r)
				result[objects.BOX] = append(result[objects.BOX], box)
			case "goal":
				goal := objects.NewGoal(loc, r)
				result[objects.GOAL] = append(result[objects.GOAL], goal)
			default:
				return nil, fmt.Errorf("unable to handle type: '%s'", match[typeIdx])
			}
		}
	}

	return result, nil
}

func ParseWorldFromString(name, content string) (world.IWorld, error) {
	splits := strings.Split(content, "\n\n")

	grid := parseGridWorld(splits[0])
	var w world.IWorld
	if len(splits) > 1 {
		objs, err := parseObjects(splits[1])
		if err != nil {
			return nil, err
		}

		w = world.NewWorld(name, grid, objs)
	} else {
		w = world.NewWorld(name, grid, make(objects.ObjectMap))
	}

	return w, nil
}

func ParseWorldFromFile(dir, filename string) (world.IWorld, error) {
	content, err := os.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		return nil, err
	}

	return ParseWorldFromString(filename, string(content))
}
