#!make
.SILENT:
.DEFAULT_GOAL := help

help: ## Show this help
	@echo "Usage:\n  make <target>\n"
	@echo "Targets:"
	@grep -h -E '^[a-zA-Z_-].+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

run-compose: ## Run server and client by docker-compose
	docker-compose up --build

run-server: ## Run only server
	go run cmd/server/main.go

run-client: ## Run only client
	go run cmd/client/main.go

deps: ## Download dependencies
	go mod download && go mod tidy

lint: ## Check code (used golangci-lint)
	GO111MODULE=on golangci-lint run
