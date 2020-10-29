.PHONY: proto
proto: clear hello
	protoc -I. \
   -I/usr/local/include \
   -I${GOPATH}/src \
   -I${GOPATH}/src/github.com/googleapis/googleapis \
   -I${PWD}/proto \
   --go_out=. \
   --micro_out=. \
   --hello_out=. \
   --hello_opt=paths=source_relative,a=b,c=d \
   proto/*

clear:
	rm -rf proto/*.go

hello:
	cd protoc-gen-hello && go install .
	cd protoc-gen-micro && go install .
