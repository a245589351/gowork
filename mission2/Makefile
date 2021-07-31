GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	docker run --rm --user `id -u ${USER}` -v `pwd`:`pwd` -w `pwd` znly/protoc -I. --go_out=plugins=grpc:. proto/*.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello *.go

.PHONY: dockerBuild
dockerBuild:
	docker build -t ve/hello:latest .

.PHONY: run
run:
	make proto
	make build
	make dockerBuild
