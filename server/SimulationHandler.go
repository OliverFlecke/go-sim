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

func (h *SimulationHandler) ServeHttp(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)

	if head == "create" && req.Method == http.MethodPost {
		// TODO: Should return error if map cannot be found
		// TODO: Should read map from request body
		id, _ := h.startSimulation("04.map")
		res.Write([]byte(id))
		return
	}

	sim := h.simulations[head]
	if sim == nil {
		http.Error(res, fmt.Sprintf("invalid simulation id %q", head), http.StatusBadRequest)
		return
	}

	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	case "map":

	case "raw":
		name := sim.GetWorld().GetName()
		content, err := os.ReadFile(filepath.Join(PATH_TO_LEVELS, name))
		if err != nil {
			fmt.Println(err.Error())
			http.Error(res, fmt.Sprintf("no level found for simulation %s", sim.Id), http.StatusBadRequest)
			return
		}
		res.Write([]byte(content))
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
