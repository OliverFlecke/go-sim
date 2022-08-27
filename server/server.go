package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/logger"
	"simulator/core/objects"
	"simulator/model/dto"
	"simulator/model/mapping"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

var simulations map[string]*simulator.Simulation

type App struct {
	SimulationHandler *SimulationHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)

	switch head {
	case "simulation":
		res.Header().Set("Access-Control-Allow-Origin", "*")
		h.SimulationHandler.ServeHttp(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func main() {
	simulations = make(map[string]*simulator.Simulation)

	app := &App{
		SimulationHandler: NewSimulationHandler(),
	}
	fmt.Printf("App listening on port 8080\n")
	http.ListenAndServe(":8080", app)

	// r := gin.Default()
	// r.Use(func(ctx *gin.Context) {
	// 	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// })

	// r.GET("/simulation/:sim/stream", StreamHandler)
	// r.GET("/simulation/:sim/map", getMapOfWorld)
	// r.POST("/simulation/create", startSimulationHandler)
	// r.POST("/simulation/:sim/agent/:agent", addActions)

	// r.Run() // listen and serve on 0.0.0.0:8080
}

func getSimulation(c *gin.Context) *simulator.Simulation {
	simId := c.Param("sim")
	sim := simulations[simId]
	if sim == nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no simulation with id '%s' exists", simId))
		return nil
	}

	return sim
}

func parseAction(sim *simulator.Simulation, body io.ReadCloser) ([]action.Action, error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	acts := &dto.ActionList{}
	if err := protojson.Unmarshal(bytes, acts); err != nil {
		logger.Error("Unable to parse text %v", err.Error())
		return nil, err
	}

	return mapping.GetActions(acts, sim), nil
}

func addActions(c *gin.Context) {
	sim := getSimulation(c)
	if sim == nil {
		return
	}

	agentId, err := strconv.ParseUint(c.Param("agent"), 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("id for agent could not be parsed: '%s'", c.Param("agent")))
		return
	}

	a := getAgent(uint32(agentId), sim)

	acts, err := parseAction(sim, c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sim.SetActions(a, acts)
	sim.Run()
	c.Status(http.StatusNoContent)
}

func StreamHandler(c *gin.Context) {
	sim := getSimulation(c)
	if sim == nil {
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	c.Stream(func(w io.Writer) bool {
		if e, ok := <-sim.GetEvents(); ok {
			w := sim.GetWorld()

			switch e.Status {
			case simulator.RUNNING:
				c.SSEvent("move", gin.H{
					"agents": w.GetAgents(),
					"boxes":  w.GetObjects(objects.BOX),
					"goals":  w.GetObjects(objects.GOAL),
				})
				return true
			case simulator.COMPLETED:
				c.SSEvent("complete", nil)
				return false
			}
		}

		return false
	})
}

func getMapOfWorld(c *gin.Context) {
	sim := getSimulation(c)
	if sim == nil {
		return
	}

	w := sim.GetWorld()
	objs := make(map[string][]objects.WorldObject)
	objs["agents"] = w.GetObjects(objects.AGENT)
	objs["goals"] = w.GetObjects(objects.GOAL)
	objs["boxes"] = w.GetObjects(objects.BOX)

	c.JSON(http.StatusOK, gin.H{
		"state": objs,
		"grid":  w.GetStaticMapAsString(),
	})
}

func startSimulationHandler(c *gin.Context) {
	id := startSimulation("SAanagram.map")

	c.JSON(http.StatusAccepted, id)
}
