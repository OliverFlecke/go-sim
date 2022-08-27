package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	app := &App{
		SimulationHandler: NewSimulationHandler(),
	}
	port := 8080

	fmt.Printf("App listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), app)
}

type App struct {
	SimulationHandler *SimulationHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	fmt.Printf("[APP] %s %s %s\n", time.Now().UTC().Format(time.RFC3339), r.Method, r.URL.Path)

	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch head {
	case "simulation":
		res.Header().Set("Access-Control-Allow-Origin", "*")
		h.SimulationHandler.ServeHTTP(res, r)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}
