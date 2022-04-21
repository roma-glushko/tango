.PHONY: help
help:
	@echo "🛠 Available Commands"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

PROJECT_DIR := $(shell git rev-parse --show-toplevel)
BIN_DIR := $(PROJECT_DIR)/bin

export PATH := $(BIN_DIR):$(PATH)

bin/packr2:
	@GOBIN=$(BIN_DIR) go install github.com/gobuffalo/packr/v2/packr2

bin/goreleaser:
	@GOBIN=$(BIN_DIR) go install github.com/goreleaser/goreleaser@latest

.PHONY: install
install: go.mod bin/packr2 bin/goreleaser ## Install project dependencies
	@go get -t -v ./...

.PHONY: lint
lint: ## Lint the codebase
	@go mod tidy
	@go fmt ./...
	@go vet ./...

.PHONY: build
build: lint bin/packr2 ## Build tango binary
	@$(BIN_DIR)/packr2
	@go build -o bin/tango

.PHONY: release
release: bin/goreleaser ## Release a new version of Tango
	@export PATH="$(BIN_DIR):$$PATH"
	@$(BIN_DIR)/goreleaser

.PHONY: run
run: ## Run tango in dev mode
	@go run main.go

.PHONY: test
test: ## Run tests
	@go test ./test/
