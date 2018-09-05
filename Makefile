COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date +%F)
BUILD=$(shell echo "${BUILDNUMBER}")
NAME="styx"
GO_LDFLAGS=-ldflags "-X main.Version=build="$(BUILD)"|commit="$(COMMIT)"|date="$(DATE)""

all: clean test fmt lint vet megacheck build proto cover

.PHONY: build
build:
	go build ${GO_LDFLAGS} -o ${NAME}

.PHONY: proto
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go get -u -v github.com/golang/protobuf/protoc-gen-go
	for file in $$(git ls-files '*.proto'); do \
		protoc -I $$(dirname $$file) --go_out=plugins=grpc:$$(dirname $$file) $$file; \
	done

.PHONY: clean
clean:
	rm -v ${NAME} | tee /dev/stderr ; rm -v coverage.txt | tee /dev/stderr

.PHONY: test
test:
	go test github.com/navinds25/maitre/pkg/saltfunc/

.PHONY: fmt
fmt:
	gofmt -s -l . | grep -v '.pb.go' | grep -v vendor | tee /dev/stderr

.PHONY: lint
lint:
	golint ./... | grep -v '.pb.go' | grep -v vendor | tee /dev/stderr

.PHONY: vet
vet:
	go vet $(shell go list ./... | grep -v vendor) | grep -v '.pb.go' | tee /dev/stderr

.PHONY: megacheck
megacheck:
	megacheck $(shell go list ./... | grep -v vendor) | grep -v '.pb.go' | tee /dev/stderr

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

