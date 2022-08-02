module simulation/runner

go 1.18

replace simulator/core => ../core

require simulator/core v0.0.0-00010101000000-000000000000

require (
	atomicgo.dev/keyboard v0.2.8 // indirect
	github.com/containerd/console v1.0.3 // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
)
