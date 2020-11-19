Project=github.com/pubgo/xprotogen
GoPath=$(shell go env GOPATH)
Version=$(shell git tag --sort=committerdate | tail -n 1)
GoROOT=$(shell go env GOROOT)
BuildTime=$(shell date "+%F %T")
CommitID=$(shell git rev-parse HEAD)
GO := GO111MODULE=on go

LDFLAGS += -X "${Project}/version.GoROOT=${GoROOT}"
LDFLAGS += -X "${Project}/version.BuildTime=${BuildTime}"
LDFLAGS += -X "${Project}/version.GoPath=${GoPath}"
LDFLAGS += -X "${Project}/version.CommitID=${CommitID}"
LDFLAGS += -X "${Project}/version.Project=${Project}"
LDFLAGS += -X "${Project}/version.Version=${Version:-v0.0.1}"

.PHONY: build
build:
	go build -ldflags '${LDFLAGS}' -mod vendor -v -o main main.go

.PHONY: proto
proto: clear gen
	protoc -I. \
   -I/usr/local/include \
   -I${GOPATH}/src \
   -I${GOPATH}/src/github.com/googleapis/googleapis \
   -I${GOPATH}/src/github.com/gogo/protobuf \
   --go_out=plugins=grpc:. \
   --micro_out=. \
	examples/proto/hello/*

	protoc -I. \
   -I/usr/local/include \
   -I${GOPATH}/src \
   -I${GOPATH}/src/github.com/googleapis/googleapis \
   -I${GOPATH}/src/github.com/gogo/protobuf \
   --go_out=plugins=grpc:. \
   --micro_out=. \
	examples/proto/login/*

.PHONY: clear
clear:
	rm -rf examples/proto/*.go
	rm -rf examples/proto/**/*.go

.PHONY: gen
gen:
	cd examples/protoc-gen-micro && go install .
