package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	app := &App{
		SimulationHandler: NewSimulationHandler(),
		MapHandler:        &LevelHandler{},
	}
	port := 8080

	fmt.Printf("App listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), app)
}

type App struct {
	SimulationHandler *SimulationHandler
	MapHandler        *LevelHandler
}

func (h *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[APP] %s %s %s\n", time.Now().UTC().Format(time.RFC3339), r.Method, r.URL.Path)

	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch head {
	case "simulation":
		h.SimulationHandler.ServeHTTP(w, r)
	case "level":
		h.MapHandler.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}
