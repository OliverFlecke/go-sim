package simulator

import "time"

type SimulationOptions struct {
	// waitForAction bool
	tickDuration time.Duration
}

func (opt *SimulationOptions) SetTickDuration(duration time.Duration) {
	opt.tickDuration = duration
}
