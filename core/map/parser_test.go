package simulationMap

import (
	"reflect"
	sim "simulator/core"
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
	expected := sim.NewGridWorld(3)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(w, expected) {
		t.Fatalf("Parsed world does not match. Expected:\n%v\nActual:\n%v", expected, w)
	}
}
