package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/agent"
	"simulator/core/level"
	"simulator/core/logger"
	"simulator/core/world"
	"simulator/model/mapping"

	"google.golang.org/protobuf/encoding/protojson"
)

const URL = "http://localhost:8080"

func downloadLevel(id string) world.IWorld {
	url := fmt.Sprintf("%s/simulation/%s/map", URL, id)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Could not download map %s", err)
	}

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Could not read body of response")
	}

	w, err := level.ParseWorldFromString(id, string(bs))
	if err != nil {
		log.Fatalf("World could not be parsed")
	}

	return w
}

func sendActions(
	sim *simulator.Simulation,
	a *agent.Agent,
	acts []action.Action) {
	httpposturl := fmt.Sprintf("%s/simulation/%s/agent/%d",
		URL, sim.Id, a.GetId())

	dtos := mapping.ToDtos(acts)
	data, err := protojson.Marshal(dtos)
	if err != nil {
		logger.Error("%v\n", err.Error())
		return
	}

	request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(data))
	if err != nil {
		logger.Error("%v\n", err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		logger.Error("%v\n", err.Error())
		return
	}
}
