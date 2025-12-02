GOBUILD = go build
GOFMT = go fmt
GOVET = go vet


TARGET = day2.out
VERSION = 1

.DEFAULT_GOAL := execute
.PHONY: build vet fmt execute all

all: execute

fmt:
	@$(GOFMT) ./...

vet: fmt
	@$(GOVET) ./...

build: vet
	$(GOBUILD) -o $(TARGET)

execute: build
	./$(TARGET)

