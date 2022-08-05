package simulationMap

import (
	"reflect"
	simulator "simulator/core"
	"testing"
)

func TestParseStringToMap(t *testing.T) {
	world, err := GetStringFromFile("../../maps/00.map")
	expected := simulator.NewGridWorld(2)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(world.GetMap(), expected.GetMap()) {
		t.Fatalf("Parsed world does not match. Expected:\n%v\nActual:\n%v. Maps: Expected:\n%v\nActual:\n%v",
			expected.ToString(), world.ToString(), expected.GetMap(), world.GetMap())
	}
}
