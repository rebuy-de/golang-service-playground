# golang-service-playground
Playground for evaluating proper golang http service structure

## Environment

* Glide needs Golang vendoring experiment enabled 
  + you need Golang version `>=1.5`
  +`export GO15VENDOREXPERIMENT=1`
* project source must be cloned to `${GOPATH}/src/github.com/rebuy-de/golang-service-playground`


## Build

1. clone repo
2. install [glide](https://github.com/Masterminds/glide): `go install github.com/Masterminds/glide`
3. install dependencies: `glide install`
4. build: `go build`

## Run on Docker

1. `make`

## Update dependencies

1. `glide update`

