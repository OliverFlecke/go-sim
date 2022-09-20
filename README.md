# Simulation engine for a Sokoban-like world

This repository contains a simulation engine to run [Sokoban](https://en.wikipedia.org/wiki/Sokoban) like simulations for general pathfinding and problem solving, both for single-agent and multi-agent situations.

Currently only simple move actions is supported on a grid-like world, but the plan is to introduce other interesting features to create new, more complicated problems.

## Packages

The repository contains the following packages:

- `core` is the main simulation engine which can represent the world and a simulation engine to execute actions in
- `server` is a HTTP based server to communicate with the engine
  - The [simulation-ui](https://github.com/OliverFlecke/simulation-ui) contains a simple UI to render a graphical representation of the simulation.
- `pathfinding` for general pathfinding algorithms
- `solver` is an attempt to solve the existing problems that has been created
- `level` directory contains sample levels
- `converter` for converting levels from an older map type

## Build Protobufs

All Protobuf models can be build with the following command.
All models will then be placed in the `simulator/models/dto` package.

```sh
protoc -I=model/dto/proto --go_out=.. model/dto/proto/*.proto
```
