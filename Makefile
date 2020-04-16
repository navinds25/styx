COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date +%F)
BUILD=$(shell echo "${BUILDNUMBER}")
CWD=$(shell pwd)
NAME=styx
GO_LDFLAGS=-ldflags "-X main.Version=build="$(BUILD)"|commit="$(COMMIT)"|date="$(DATE)""

all: clean proto tidy build

.PHONY: build
build:
	mkdir bin | tee /dev/stderr
	cd bin && go build ${GO_LDFLAGS} ../cmd/styxnode/styxnode.go
	cd bin && go build ${GO_LDFLAGS} ../cmd/styxmaster/styxmaster.go

.PHONY: proto
proto:
	cd api/nodeconfig && protoc -I. --go_out=plugins=grpc,paths=source_relative:. *.proto
	cd api/filetransfer && protoc -I. --go_out=plugins=grpc,paths=source_relative:. *.proto

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	find api -name *.pb.go -exec rm {} \;
	rm -rfv bin | tee /dev/stderr ; rm -v styx.log | tee /dev/stderr

.PHONY: test
test:
	go test -race github.com/navinds25/maitre/pkg/saltfunc/

.PHONY: fmt
fmt:
	gofmt -s -l . | grep -v '.pb.go' | grep -v vendor | tee /dev/stderr

.PHONY: lint
lint:
	golint ./... | grep -v '.pb.go' | grep -v vendor | tee /dev/stderr

.PHONY: vet
vet:
	go vet $(shell go list ./... | grep -v vendor) | grep -v '.pb.go' | tee /dev/stderr

.PHONY: hostkey
hostkey:
	mkdir -p etc/ssh/
	ssh-keygen -A -f ${CWD}

.PHONY: sshkey
sshkey:
	ssh-keygen -t rsa -N "" -f host_key && rm -v host_key.pub

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

