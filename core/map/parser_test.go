package simulationMap

import (
	"reflect"
	sim "simulator/core"
	"simulator/core/location"
	"simulator/core/objects"
	"testing"
)

func TestParseStringToMap(t *testing.T) {
	w, err := ParseWorldFromFile("../../maps/00.map")
	world := w.(*sim.World)
	expected := sim.NewGridWorld(2)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(world.GetMap(), expected.GetMap()) {
		t.Fatalf("Parsed world does not match. Expected:\n%v\nActual:\n%v. Maps: Expected:\n%v\nActual:\n%v",
			expected.ToString(), world.ToString(), expected.GetMap(), world.GetMap())
	}
}

func TestPaseMapFile(t *testing.T) {
	w, err := ParseWorldFromFile("../../maps/02.map")

	objs := make(objects.ObjectMap)
	objs[objects.AGENT] = []objects.WorldObject{
		sim.NewAgentWithStartLocation("unused", '0', location.New(1, 1)),
	}
	objs[objects.BOX] = []objects.WorldObject{
		objects.NewBox(location.New(2, 2), 'A'),
	}
	objs[objects.GOAL] = []objects.WorldObject{
		objects.NewGoal(location.New(3, 3), 'a'),
	}
	expected := sim.NewWorld(sim.NewGrid(3), objs)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(w, expected) {
		t.Fatalf("Parsed world does not match. Expected:\n%v\nActual:\n%v", expected, w)
	}
}
