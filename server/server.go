// package server

package main

import (
	"io"
	"net/http"
	"time"

	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/level"
	"simulator/core/logger"
	"simulator/core/objects"
	"simulator/model/dto"
	"simulator/model/mapping"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

var sim *simulator.Simulation

func main() {
	mapName := "../maps/04.map"
	w, _ := level.ParseWorldFromFile(mapName)

	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(250 * time.Millisecond)
	sim = simulator.NewSimulation(w, opt)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/stream", StreamHandler)
	r.GET("/simulation/map", getMapOfWorld)
	r.POST("/agent/:agent", addActions)

	// genAction()
	go sim.Run(nil)

	r.Run() // listen and serve on 0.0.0.0:8080
}

type ActionDto struct {
	Type      string `json:"type" binding:"required"`
	Direction string `json:"direction" binding:"required"`
}

type List struct {
	Messages []ActionDto `binding:"required"`
}

func parseAction(c *gin.Context) ([]action.Action, error) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	logger.Info("got bytes from body: %v\n", string(bytes))
	acts := &dto.ActionList{}
	if err := protojson.Unmarshal(bytes, acts); err != nil {
		logger.Error("Unable to parse text %v", err.Error())
		return nil, err
	}

	return mapping.GetActions(acts, sim), nil
}

func addActions(c *gin.Context) {
	agentId := c.Param("agent")
	logger.Verbose("adding action for %v\n", agentId)
	a := sim.GetWorld().GetAgents()[0]

	acts, err := parseAction(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sim.SetActions(a, acts)
	c.Status(http.StatusNoContent)
}

func StreamHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Stream(func(w io.Writer) bool {
		if _, ok := <-sim.GetEvents(); ok {
			w := sim.GetWorld()
			c.SSEvent("move", gin.H{
				"agents": w.GetAgents(),
				"boxes":  w.GetObjects(objects.BOX),
				"goals":  w.GetObjects(objects.GOAL),
			})

			// genAction()
			return true
		}

		return false
	})
}

func getMapOfWorld(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	w := sim.GetWorld()
	objs := make(map[string][]objects.WorldObject)
	objs["agents"] = w.GetObjects(objects.AGENT)
	objs["goals"] = w.GetObjects(objects.GOAL)
	objs["boxes"] = w.GetObjects(objects.BOX)

	c.JSON(200, gin.H{
		"state": objs,
		"grid":  w.GetStaticMapAsString(),
	})
}
