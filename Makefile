SHELL := /bin/bash
GO_VERSION=`cat GO_VERSION`
DOCKER_WORKDIR=.docker-cache
DOCKER_RUN=docker run --rm -v "$$PWD/.:${DOCKER_WORKDIR}" -v "`go env GOPATH`/pkg/mod/.:/go/pkg/mod:ro" -w ${DOCKER_WORKDIR}
DOCKER_GO_BUILD=go build -mod=readonly -a -installsuffix cgo -ldflags "$$LD_FLAGS"

test-main:
	go test main.go main_test.go

test-coverage:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

format:
	goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

check-go:
	golangci-lint run

build-docker:
	cp ${BUILD_DIR}/gotify-linux-amd64 ./docker/skillbird-go-rest-api
	cd ${DOCKER_DIR} && \
		docker build \
		-t skillbird/rest-api:latest \
		-t skillbird/rest-api:${VERSION} \
		-t skillbird/rest-api:$(shell echo $(VERSION) | cut -d '.' -f -2) .
	rm ${DOCKER_DIR}skillbird-go-rest-api

build-bin:
	go build

.PHONY:	test-main	test-coverage	format	check-go	build-docker	build-bin
