package simulator

import "time"

type SimulationEvent struct {
	CurrentTime time.Time
	Err         error
}
