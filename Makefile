
SHELL := /bin/bash
CURRENT_PATH = $(shell pwd)
APP_NAME = bridge
APP_VERSION = 1.0.0

# build with verison infos
VERSION_DIR = main
BUILD_DATE = $(shell date +%FT%T)
GIT_COMMIT = $(shell git log --pretty=format:'%h' -n 1)
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

GOLDFLAGS += -X "${VERSION_DIR}.date=${BUILD_DATE}"
GOLDFLAGS += -X "${VERSION_DIR}.commit=${GIT_COMMIT}"
GOLDFLAGS += -X "${VERSION_DIR}.branch=${GIT_BRANCH}"
GOLDFLAGS += -X "${VERSION_DIR}.version=${APP_VERSION}"

STATIC_LDFLAGS += ${GOLDFLAGS}
STATIC_LDFLAGS += -linkmode external -extldflags -static

GO = GO111MODULE=on go
TEST_PKGS := $(shell $(GO) list ./...  | grep -v 'mock_*')

RED=\033[0;31m
GREEN=\033[0;32m
BLUE=\033[0;34m
NC=\033[0m

.PHONY: test

help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## make test: Run go unittest
test:
	go generate ./...
	@$(GO) test ${TEST_PKGS} -count=1

## make test-coverage: Test project with cover
test-coverage:
	go generate ./...
	@go test -short -coverprofile cover.out -covermode=atomic ${TEST_PKGS}
	@cat cover.out >> coverage.txt

packr:
	cd pkg/repo && packr

prepare:
	cd scripts && bash prepare.sh

## make install: Go install the project
install:
	cd pkg/repo && packr
	$(GO) install -ldflags '${GOLDFLAGS}' ./cmd/${APP_NAME}
	@printf "${GREEN}Build bridge successfully!${NC}\n"

build:
	cd pkg/repo && packr
	@mkdir -p bin
	$(GO) build -ldflags '${GOLDFLAGS}' ./cmd/${APP_NAME}
	@mv ./bridge bin
	@printf "${GREEN}Build bridge successfully!${NC}\n"

docker-build: packr
	$(GO) install -ldflags '${STATIC_LDFLAGS}' ./cmd/${APP_NAME}
	@echo "Build bridge successfully"

## make build-linux: Go build linux executable file
build-linux:
	cd scripts && bash cross_compile.sh linux-amd64 ${CURRENT_PATH}

plugin:
	@mkdir -p build
	cd internal/monitor/chain/klaytn && make plugin

## make release: Build release before push
release-binary:
	@cd scripts && bash release_binary.sh

## make linter: Run golanci-lint
linter:
	golangci-lint run
