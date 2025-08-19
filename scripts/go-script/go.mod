module fern-sdk-sandbox-sandbox-go-simple

go 1.23.0

toolchain go1.24.5

// Local SDK dependency
replace github.com/jsklan/fern-sdk-sandbox => ../../sdks/go

require (
	github.com/jsklan/fern-sdk-sandbox v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/net v0.42.0 // indirect
)
