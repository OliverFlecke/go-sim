package main

import (
	"fmt"
	"net/http"
	simulator "simulator/core"
	"strconv"
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

		switch r.Method {
		case http.MethodPost:
			acts, _ := parseAction(sim, r.Body)
			sim.SetActions(a, acts)
			sim.Run()
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
