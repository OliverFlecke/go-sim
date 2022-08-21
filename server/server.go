// package server

package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/direction"
	"simulator/core/level"
	"simulator/core/logger"
	"simulator/core/objects"
	"simulator/core/utils"

	"github.com/gin-gonic/gin"
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

type ActionDto struct {
	Type      string `json:"type" binding:"required"`
	Direction string `json:"direction" binding:"required"`
}

type List struct {
	Messages []ActionDto `binding:"required"`
}

func directionFromString(dir string) (direction.Direction, error) {
	switch strings.ToLower(dir) {
	case "north":
		return direction.NORTH, nil
	case "south":
		return direction.SOUTH, nil
	case "east":
		return direction.EAST, nil
	case "west":
		return direction.WEST, nil
	default:
		return 0, fmt.Errorf("unknown direction")
	}
}

func addActions(c *gin.Context) {
	agentId := c.Param("agent")
	logger.Verbose("adding action for %v\n", agentId)

	data := new([]ActionDto)
	err := c.BindJSON(data)

	// data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// logger.Info("data: %v", data)

	logger.Verbose("data: %v\n", *data)

	a := sim.GetWorld().GetAgents()[0]
	acts := utils.Mapi(*data, func(_ int, dto ActionDto) action.Action {
		switch dto.Type {
		case "Move":
			dir, err := directionFromString(dto.Direction)
			if err != nil {
				fmt.Println(err.Error())
			}
			return action.NewMove(dir)
		default:
			return nil
		}
	})
	acts = utils.Filteri(acts, func(_ int, a action.Action) bool {
		return a != nil
	})

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
