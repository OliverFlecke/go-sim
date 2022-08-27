package simulator

import "time"

type SimulationEvent struct {
	CurrentTime time.Time
	Err         error
	Status      SimulationStatus
}

type SimulationStatus uint8

const (
	NONE SimulationStatus = iota
	RUNNING
	PAUSED
	COMPLETED
)
