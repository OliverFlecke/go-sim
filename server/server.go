// package server

package main

import (
	"io"
	"time"

	simulator "simulator/core"
	maps "simulator/core/map"
	"simulator/core/objects"

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

	go sim.Run(nil)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func StreamHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Stream(func(w io.Writer) bool {
		if e, ok := <-sim.GetEvents(); ok {
			c.SSEvent("tick", e.CurrentTime)
			return true
		}

		return false
	})
}

func getMapOfWorld(c *gin.Context) {
	objs := make(map[string][]objects.WorldObject)
	objs["agent"] = sim.GetWorld().GetObjects(objects.AGENT)
	objs["goal"] = sim.GetWorld().GetObjects(objects.GOAL)
	objs["box"] = sim.GetWorld().GetObjects(objects.BOX)

	c.JSON(200, gin.H{
		"world": objs,
		"grid":  sim.GetWorld().GetStaticMapAsString(),
	})
}
