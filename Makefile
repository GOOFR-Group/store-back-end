CODEGEN_PATH = ./internal/oapi
CODEGEN_PACKAGE = oapi

.PHONY: all
all: dev production

# Builds

# production build
.PHONY: production
production: get-bins generate 
	go build -mod vendor -o dist/store ./cmd

# development build
.PHONY: dev
dev: get-bins generate 
	go build -mod vendor -tags "dev" -o dist/store-dev ./cmd

# Binaries

.PHONY: get-bins
get-bins: get-tgcon get-oapi

.PHONY: get-tgcon
get-tgcon:
	command -v tgcon || GO111MODULE=off go get github.com/amarjeetanandsingh/tgcon

.PHONY: get-oapi
get-oapi:
	command -v oapi-codegen || GO111MODULE=off go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

# Generators

# generate constants for each struct field tag matching database table columns
.PHONY: generate
generate: storage oapi-codegen

.PHONY: storage
storage:
	go generate internal/storage/storage.go

.PHONY: oapi-codegen
oapi-codegen:
	oapi-codegen --package ${CODEGEN_PACKAGE} -generate types,skip-prune ${CODEGEN_PATH}/store.yaml > ${CODEGEN_PATH}/types.gen.go
	oapi-codegen --package ${CODEGEN_PACKAGE} -generate chi-server  ${CODEGEN_PATH}/store.yaml > ${CODEGEN_PATH}/server.gen.go
	oapi-codegen --package ${CODEGEN_PACKAGE} -generate spec  ${CODEGEN_PATH}/store.yaml > ${CODEGEN_PATH}/spec.gen.go

# Tools

.PHONY: mod
mod:
	go mod vendor -v

.PHONY: clean
clean:
	rm -rf dist/*

lint: get-linter
	golangci-lint run --timeout=5m

get-linter:
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin
