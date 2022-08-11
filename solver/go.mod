module simulation/solver

go 1.18

replace simulator/core => ../core

require (
	simulator/core v0.0.0-00010101000000-000000000000
	simulator/pathfinding v0.0.0-00010101000000-000000000000
)

require (
	github.com/deckarep/golang-set/v2 v2.1.0 // indirect
	github.com/ethereum/go-ethereum v1.10.21 // indirect
)

replace simulator/pathfinding => ../pathfinding
