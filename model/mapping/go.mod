module simulator/model/mapping

replace simulator/model/dto => ../dto

replace simulator/core => ../../core

go 1.19

require (
	github.com/stretchr/testify v1.8.0
	simulator/core v0.0.0-00010101000000-000000000000
	simulator/model/dto v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
