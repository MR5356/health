.DEFAULT_GOAL := help

.PHONY: help
help:  ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: deps
deps:  ## Install dependencies
	go get -d -v -t ./...

.PHONY: test
test: deps ## Run tests
	@go test -v ./... -coverprofile=coverage.out && go tool cover -func=coverage.out