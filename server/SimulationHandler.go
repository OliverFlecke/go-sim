package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	simulator "simulator/core"
	"simulator/core/level"
	"simulator/core/objects"
	"simulator/model/dto"
	"simulator/model/mapping"

	"google.golang.org/protobuf/encoding/protojson"
)

type SimulationHandler struct {
	simulations  map[string]*simulator.Simulation
	agentHandler *AgentHandler
}

func NewSimulationHandler() *SimulationHandler {
	return &SimulationHandler{
		simulations:  make(map[string]*simulator.Simulation),
		agentHandler: NewAgentHandler(),
	}
}

func (h *SimulationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch {
	case head == "create" && r.Method == http.MethodPost:
		h.handleCreate(w, r)
	case head != "":
		sim := h.simulations[head]
		if sim == nil {
			http.Error(w, fmt.Sprintf("invalid simulation id %q", head), http.StatusBadRequest)
			return
		}

		head, r.URL.Path = ShiftPath(r.URL.Path)
		switch {
		case head == "agent":
			h.agentHandler.Handle(sim).ServeHTTP(w, r)
		case head == "stream" && r.Method == http.MethodGet:
			h.streamEvents(sim).ServeHTTP(w, r)
		case head == "level" && r.Method == http.MethodGet:
			h.sendRawLevelContent(sim).ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}

	default:
		http.NotFound(w, r)
	}
}

func (h *SimulationHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	// TODO: Should return error if map cannot be found

	bs, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}
	opt := &dto.CreateSimulationOptions{}
	protojson.Unmarshal(bs, opt)

	// TODO: level name should be checked for whether the extension is there
	id, err := h.startSimulation(opt.Level)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Level '%s' does not exists", opt.Level), http.StatusBadRequest)
		return
	}

	w.Write([]byte(id))
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

		for e := range sim.GetEvents() {
			switch e.Status {
			case simulator.COMPLETED:
				SendSSE(w, "complete", nil)
				return
			case simulator.RUNNING:
				state := createWorldState(sim)
				bs, err := protojson.Marshal(state)
				if err != nil {
					fmt.Println(err.Error())
					http.Error(w, "could not serialize state", http.StatusInternalServerError)
					return
				}

				SendSSE(w, "move", bs)
			}
		}
	})
}

func createWorldState(sim *simulator.Simulation) *dto.WorldState {
	state := &dto.WorldState{
		Agents: make([]*dto.Agent, 0),
	}

	for _, x := range sim.GetWorld().GetAgents() {
		state.Agents = append(state.Agents, mapping.AgentToDto(x))
	}
	for _, x := range sim.GetWorld().GetBoxes() {
		state.Boxes = append(state.Boxes, mapping.BoxToDto(&x))
	}
	for _, x := range sim.GetWorld().GetObjects(objects.GOAL) {
		g := x.(*objects.Goal)
		state.Goals = append(state.Goals, mapping.GoalToDto(g))
	}

	return state
}
