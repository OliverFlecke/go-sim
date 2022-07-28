package simulator

import (
	"regexp"
	"testing"
)

func TestRun(t *testing.T) {
	value := Run()
	want := regexp.MustCompile(`Hello from simulator`)

	if !want.MatchString(value) {
		t.Fatalf(`Hello`)
	}
}

func TestMove(t *testing.T) {
	agent := newAgent("test agent")
	directions := []Direction{NORTH, EAST, SOUTH, WEST}
	locations := []Location{{x: 0, y: 1}, {x: 1, y: 1}, {x: 1, y: 0}, {x: 0, y: 0}}

	for i, dir := range directions {
		move(agent, dir)
		if agent.location != locations[i] {
			t.Fatal(`Wrong location for agent`)
		}
	}
}
