# Project settings
BINARY_NAME=armis
PKG=./...
MAIN=./cmd/main.go
BUILD_DIR=build

# Git info
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)


# Default target
run: ## Run the app
	@clear
	@go run $(MAIN)

all: build ## Build the project (default)

compile: ## Compile (without installing)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN)

build: clean compile
	./$(BUILD_DIR)/$(BINARY_NAME)


test: ## Run tests
	go test -v $(PKG)

dep: ## Download and vendor dependencies
	go mod tidy
	go mod vendor

clean: ## Remove build artifacts
	rm -rf $(BUILD_DIR)

push: ## Push current branch to origin
	git push origin $(BRANCH)

pull: ## Pull current branch from origin
	git pull origin $(BRANCH)

# lint: ## Lint the project (requires golangci-lint)
# 	golangci-lint run ./...

info: ## Show build info
	@echo "Branch:   $(BRANCH)"

help: ## Show this help
	@echo "Usage: make [target]"
	@awk 'BEGIN {FS = ":.*##"; printf "\nAvailable targets:\n"} /^[a-zA-Z0-9_-]+:.*##/ {printf "  %-12s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
