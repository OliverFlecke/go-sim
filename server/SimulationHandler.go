package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	simulator "simulator/core"
	"simulator/core/level"
)

type SimulationHandler struct {
	simulations map[string]*simulator.Simulation
}

func NewSimulationHandler() *SimulationHandler {
	return &SimulationHandler{
		simulations: make(map[string]*simulator.Simulation),
	}
}

func (h *SimulationHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[APP] %s %s %s\n", time.Now().UTC().Format(time.RFC3339), r.Method, r.URL.Path)

	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	if head == "create" && r.Method == http.MethodPost {
		// TODO: Should return error if map cannot be found
		// TODO: Should read map from request body
		id, _ := h.startSimulation("04.map")
		w.Write([]byte(id))
		return
	}

	sim := h.simulations[head]
	if sim == nil {
		http.Error(w, fmt.Sprintf("invalid simulation id %q", head), http.StatusBadRequest)
		return
	}

	head, r.URL.Path = ShiftPath(r.URL.Path)
	switch head {
	case "stream":
		h.streamEvents(sim).ServeHTTP(w, r)
	case "agent": // TODO:

	case "map":
		h.sendRawLevelContent(sim).ServeHTTP(w, r)
	}
}

func (h *SimulationHandler) startSimulation(levelName string) (string, error) {
	id := generateId()
	w, err := level.ParseWorldFromFile(PATH_TO_LEVELS, levelName)
	if err != nil {
		return "", err
	}

	// TODO: Get these options from request
	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(50 * time.Millisecond)

	// TODO: There should be a constructor to take the ID
	sim := simulator.NewSimulation(w, opt)
	sim.Id = id
	h.simulations[id] = sim

	return id, nil
}

func (h *SimulationHandler) sendRawLevelContent(sim *simulator.Simulation) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		name := sim.GetWorld().GetName()
		content, err := os.ReadFile(filepath.Join(PATH_TO_LEVELS, name))
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, fmt.Sprintf("no level found for simulation %s", sim.Id), http.StatusBadRequest)
			return
		}
		w.Write([]byte(content))
	})
}

func (h *SimulationHandler) streamEvents(sim *simulator.Simulation) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, _ := w.(http.Flusher)

		for e := range sim.GetEvents() {
			switch e.Status {
			case simulator.COMPLETED:
				fmt.Fprint(w, "complete")
				flusher.Flush()
				return
			case simulator.RUNNING:
				fmt.Fprintf(w, "data: %s\n\n", e.CurrentTime)
				flusher.Flush()
			}
		}
	})
}
