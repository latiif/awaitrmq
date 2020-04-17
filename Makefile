BINARY = awaitrmq
VERIFY_REPORT = vet.report
TEST_REPORT = tests.xml
GOARCH = amd64

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Symlink into GOPATH
GITHUB_USERNAME=latiif
BUILD_DIR=$(shell pwd)
BIN_DIR=${BUILD_DIR}/bin
CURRENT_DIR=$(shell pwd)
BUILD_DIR_LINK=$(shell readlink ${BUILD_DIR})

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

all: clean test vet linux darwin windows

linux: 
	cd ${BUILD_DIR}; \
	CGO_ENABLED=0  GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BIN_DIR}/${BINARY}-linux-${GOARCH} . ; \
	cd - >/dev/null

darwin:
	cd ${BUILD_DIR}; \
	CGO_ENABLED=0  GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BIN_DIR}/${BINARY}-darwin-${GOARCH} . ; \
	cd - >/dev/null

windows:
	cd ${BUILD_DIR}; \
	CGO_ENABLED=0  GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BIN_DIR}/${BINARY}-windows-${GOARCH}.exe . ; \
	cd - >/dev/null

docker:
	cd ${BUILD_DIR}; \
	docker build -f Dockerfile -t ${GITHUB_USERNAME}/${BINARY}:${COMMIT} -t ${GITHUB_USERNAME}/${BINARY}:${BRANCH} -t ${GITHUB_USERNAME}/${BINARY}:latest .
	docker push ${GITHUB_USERNAME}/${BINARY}:${COMMIT}
	docker push ${GITHUB_USERNAME}/${BINARY}:${BRANCH}
	docker push ${GITHUB_USERNAME}/${BINARY}:latest
	cd - >/dev/null

vet:
	-cd ${BUILD_DIR}; \
	go mod verify > ${VERIFY_REPORT} 2>&1 ; \
	cd - >/dev/null

fmt:
	cd ${BUILD_DIR}; \
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	-rm -f ${TEST_REPORT}
	-rm -f ${VERIFY_REPORT}
	-rm -f ${BIN_DIR}/${BINARY}-*
	-rmdir ${BIN_DIR}

.PHONY: linux darwin windows test vet fmt clean