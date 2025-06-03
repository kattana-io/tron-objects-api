GO_PATH := $(shell go env GOPATH)

lint: check-lint dep
	golangci-lint run --timeout=5m -c .golangci.yml

check-lint:
	@which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_PATH)/bin v1.64.8

dep:
	@go mod tidy
	@go mod download

update:
	@go mod tidy
	@go mod vendor