GOCMD=$(shell echo go)
GOLINT=$(shell echo golangci-lint)

fmt:
	@echo "+ $@"
	@$(GOCMD) fmt ./...

lint: 
	@echo "+ $@"
	@${GOLINT} run

test:
	@echo "+ $@"
	@$(GOCMD) test ./... -cover

build:
	@echo "+ $@"
	@$(GOCMD) build cmd/example.go

all: fmt lint test build

run:
	@echo "+ $@"
	@$(GOCMD) run cmd/example.go
	
