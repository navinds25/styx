COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date +%F)
BUILD=$(shell echo "${BUILDNUMBER}")
CWD=$(shell pwd)
NAME=styx
GO_LDFLAGS=-ldflags "-X main.Version=build="$(BUILD)"|commit="$(COMMIT)"|date="$(DATE)""

all: clean tidy proto build

.PHONY: build
build:
	mkdir bin | tee /dev/stderr
	cd bin && go build ${GO_LDFLAGS} ../cmd/styxnode/styxnode.go
	cd bin && go build ${GO_LDFLAGS} ../cmd/styxmaster/styxmaster.go

.PHONY: local-docker
local-docker:
	cd build && ./build_docker.sh

.PHONY: proto
proto:
	cd api/nodeconfig && protoc -I. --go_out=plugins=grpc,paths=source_relative:. *.proto
	cd api/filetransfer && protoc -I. --go_out=plugins=grpc,paths=source_relative:. *.proto

.PHONY: certs
certs:
	cd tools/certificates && ./generate_certs.sh

.PHONY: clean
clean:
	find api -name *.pb.go -exec rm {} \;
	rm -rfv bin | tee /dev/stderr ; rm -v styx.log | tee /dev/stderr
	rm -v coverage.txt | tee /dev/stderr;
	find . -type f \( -name "*.pem" -o -name "*.csr" -o -name "host_key" \) -exec rm {} \;
	go mod tidy

.PHONY: test
test:
	go test -v -race github.com/navinds25/styx/pkg/nodeconfig...

.PHONY: fmt
fmt:
	gofmt -s -l . | grep -v '.pb.go' | grep -v vendor | tee /dev/stderr
	golint ./... | grep -v '.pb.go' | grep -v vendor | tee /dev/stderr
	go vet $(shell go list ./... | grep -v vendor) | grep -v '.pb.go' | tee /dev/stderr

.PHONY: cover
cover: ## Runs go test with coverage
	@echo "" > coverage.txt
	@for d in $(shell go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out -covermode=atomic "$$d"; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi; \
	done;

