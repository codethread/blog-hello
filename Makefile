PROJECT_NAME := "blogui"
PKG := "github.com/AHDesigns/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all dep build test lint

all: build

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: dep ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

format: dep ## Use go formatter to align style
	@go fmt

dep: ## Get the dependencies
	@go mod download

build: dep ## Build the binary file
	@go build $(PKG)

dev: dep ## Run the app in watch mode
	@air

checks: dep format lint test ## Run all checks

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
