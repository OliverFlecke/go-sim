// package server

package main

import (
	"io"
	"math/rand"
	"time"

	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/direction"
	maps "simulator/core/map"
	"simulator/core/objects"
	"simulator/core/utils"

	"github.com/gin-gonic/gin"
)

var sim *simulator.Simulation

func main() {
	mapName := "../maps/04.map"
	w, _ := maps.ParseWorldFromFile(mapName)

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

	genAction()
	go sim.Run(nil)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func genAction() {
	// TEMP: generating random actions to keep sim running
	a := sim.GetWorld().GetAgents()[0]
	if len(sim.GetActions(a)) > 0 {
		return
	}

	potential := utils.Filteri(
		utils.Mapi(direction.All, func(_ int, dir direction.Direction) action.MoveAction {
			return *action.NewMove(dir)
		}),
		func(_ int, act action.MoveAction) bool {
			return act.IsValid(a, sim.GetWorld())
		})
	// logger.Verbose("Agent: %v, Actions: %v\n", a, potential)
	sim.SetActions(a, []action.Action{potential[rand.Intn(len(potential))]})
}

func StreamHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Stream(func(w io.Writer) bool {
		if _, ok := <-sim.GetEvents(); ok {
			// c.SSEvent("tick", e.CurrentTime)
			a := sim.GetWorld().GetAgents()[0]
			c.SSEvent("move", a)

			genAction()
			return true
		}

		return false
	})
}

func getMapOfWorld(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	objs := make(map[string][]objects.WorldObject)
	objs["agent"] = sim.GetWorld().GetObjects(objects.AGENT)
	objs["goal"] = sim.GetWorld().GetObjects(objects.GOAL)
	objs["box"] = sim.GetWorld().GetObjects(objects.BOX)

	c.JSON(200, gin.H{
		"world": objs,
		"grid":  sim.GetWorld().GetStaticMapAsString(),
	})
}
