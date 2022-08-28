package simulator

import "time"

type SimulationOptions struct {
	// waitForAction bool
	TickDuration time.Duration
}

func (opt *SimulationOptions) SetTickDuration(duration time.Duration) {
	opt.TickDuration = duration
}
