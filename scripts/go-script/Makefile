# Makefile for fern-sdk-sandbox Go SDK Sandbox

.PHONY: help build run clean tidy mod-download test

help: ## Display this help message
    @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the sandbox binary
    go build -o bin/sandbox main.go

run: ## Run the sandbox
    go run main.go

clean: ## Clean build artifacts
    rm -rf bin/
    go clean

tidy: ## Tidy go modules
    go mod tidy

mod-download: ## Download dependencies
    go mod download

test: ## Run tests (if any)
    go test ./...

install-deps: ## Install development dependencies
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: ## Run linter
    golangci-lint run

fmt: ## Format code
    go fmt ./...

vet: ## Run go vet
    go vet ./...

check: fmt vet lint ## Run all code quality checks

.DEFAULT_GOAL := help
