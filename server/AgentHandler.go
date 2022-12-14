package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/logger"
	"simulator/model/dto"
	"simulator/model/mapping"
	"strconv"

	"google.golang.org/protobuf/encoding/protojson"
)

type AgentHandler struct{}

func NewAgentHandler() *AgentHandler {
	return &AgentHandler{}
}

func (h *AgentHandler) Handle(sim *simulator.Simulation) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var head string
		head, r.URL.Path = ShiftPath(r.URL.Path)
		id, err := strconv.Atoi(head)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid agent id %q", head), http.StatusBadRequest)
			return
		}

		a := getAgent(uint32(id), sim)

		head, r.URL.Path = ShiftPath(r.URL.Path)

		switch {
		case head == "" && r.Method == http.MethodPost:
			acts, err := parseAction(sim, r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				sim.SetActions(a, acts)
				sim.Run()
				w.WriteHeader(http.StatusNoContent)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func parseAction(sim *simulator.Simulation, body io.ReadCloser) ([]action.Action, error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("request body could not be read")
	}

	acts := &dto.ActionList{}
	if err := protojson.Unmarshal(bytes, acts); err != nil {
		logger.Error("Unable to parse text %v", err.Error())
		return nil, fmt.Errorf("request body could not be read")
	}

	return mapping.GetActions(acts, sim), nil
}
