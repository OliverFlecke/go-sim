module simulator/server

go 1.19

replace simulator/core => ../core

replace simulator/model/dto => ../model/dto

replace simulator/model/mapping => ../model/mapping

require (
	github.com/google/uuid v1.3.0
	google.golang.org/protobuf v1.28.1
	simulator/core v0.0.0-00010101000000-000000000000
	simulator/model/dto v0.0.0-00010101000000-000000000000
	simulator/model/mapping v0.0.0-00010101000000-000000000000
)
