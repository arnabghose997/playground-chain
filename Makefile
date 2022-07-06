export GO111MODULE=on

GOBIN = $(shell go env GOPATH)/bin

build: go.sum
		go build -mod=readonly -o ${GOBIN}/playground-chaind  ./cmd/playground-chaind

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify