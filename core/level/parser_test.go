package level

import (
	"reflect"
	"simulator/core/agent"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStringToMap(t *testing.T) {
	wp, err := ParseWorldFromFile("../../maps", "00.map")
	w := wp.(*world.World)
	expected := world.NewGridWorld(2)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(w.GetMap(), expected.GetMap()) {
		t.Fatalf("Parsed world does not match. Maps: Expected:\n%v\nActual:\n%v", expected.GetMap(), w.GetMap())
	}
}

func TestPaseMapFile(t *testing.T) {
	w, err := ParseWorldFromFile("../../maps/", "02.map")

	objs := make(objects.ObjectMap)
	objs[objects.AGENT] = []objects.WorldObject{
		agent.NewAgentWithStartLocation('0', location.New(1, 1)),
	}
	objs[objects.BOX] = []objects.WorldObject{
		objects.NewBox(location.New(2, 2), 'A'),
	}
	objs[objects.GOAL] = []objects.WorldObject{
		objects.NewGoal(location.New(3, 3), 'a'),
	}
	expected := world.NewWorld("", world.NewGrid(3), objs)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(w, expected) {
		t.Fatalf("Parsed world does not match. Expected:\n%v\nActual:\n%v", expected, w)
	}
}

// Error tests
func TestParseWorldFromFileNotFound(t *testing.T) {
	_, err := ParseWorldFromFile("", "non_existing_file")

	assert.EqualError(t, err, "open non_existing_file: no such file or directory")
}

func TestParseInt(t *testing.T) {
	v, err := parseInt("1234")
	if err != nil {
		t.Fatalf("Received unexpected error %v", err)
	}

	assert.Equal(t, 1234, v, "Parsed value should match")
}

func TestParseIntWithInvalidValue(t *testing.T) {
	_, err := parseInt("random string")
	assert.EqualError(t, err, `strconv.ParseInt: parsing "random string": invalid syntax`)
}
