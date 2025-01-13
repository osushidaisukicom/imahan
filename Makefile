.DEFAULT_GOAL := help

GO_VERSION := 1.23

# BUILD_COMMAND を docker にしたら docker で動くかも
BUILD_COMMAND := buildah
REGISTORY_ENDPOINT := docker://localhost:5000

.PHONY: help
# https://qiita.com/itoi10/items/5766df81fa28348f3fad
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: fmt
fmt: ## Format
	@go fmt ./...

