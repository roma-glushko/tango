.PHONY: help
help:
	@echo "🛠 Available Commands"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

PROJECT_DIR := $(shell git rev-parse --show-toplevel)
BIN_DIR := $(PROJECT_DIR)/bin
TMP_DIR := $(PROJECT_DIR)/tmp/bin

export PATH := $(BIN_DIR):$(PATH)

.PHONY: tools
tools: ## Install static checkers & other binaries
	@echo "🚚 Downloading tools.."
	@mkdir -p $(TMP_DIR)
	@test -f $(TMP_DIR)/goreleaser || GOBIN=$(TMP_DIR) go install github.com/goreleaser/goreleaser/v2@v2.13.3
	@test -f $(TMP_DIR)/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(TMP_DIR) v2.9.0

.PHONY: lint
lint: tools ## Lint the codebase
	@go mod tidy
	@$(TMP_DIR)/golangci-lint run ./...

.PHONY: build
build: lint ## Build tango binary
	@go build -o bin/tango

.PHONY: release
release: tools ## Release a new version of Tango
	@$(TMP_DIR)/goreleaser

.PHONY: run
run: ## Run tango in dev mode
	@go run main.go

.PHONY: test
test: ## Run tests
	@go test ./test/ -run 'TestCreate(Custom|Browser)Report'

.PHONY: test-all
test-all: ## Run all tests (requires GeoIP database)
	@go test ./test/

.PHONY: clean
clean:
	@rm -rf bin dist tmp