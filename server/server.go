// package server

package main

import (
	"io"
	"time"

	simulator "simulator/core"
	maps "simulator/core/map"

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
	go sim.Run(nil)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func StreamHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// stream := make(chan time.Time)
	// go func() {
	// 	defer close(stream)
	// 	for {
	// 		stream <- time.Now()
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	c.Stream(func(w io.Writer) bool {
		if e, ok := <-sim.GetEvents(); ok {
			c.SSEvent("tick", e.CurrentTime)
			return true
		}
		// if msg, ok := <-stream; ok {
		// 	c.SSEvent("message", msg)
		// 	return true
		// }

		return false
	})
}
