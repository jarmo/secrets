BINARY = secrets
GOARCH = amd64
GO_BUILD = GOARCH=${GOARCH} go build -mod=vendor
PREFIX ?= ${GOPATH}

all: test clean linux darwin windows

clean:
	rm -rf bin/

vendor:
	go mod vendor

linux: vendor
	GOOS=linux ${GO_BUILD} -o bin/linux_${GOARCH}/${BINARY}

darwin: vendor
	GOOS=darwin ${GO_BUILD} -o bin/darwin_${GOARCH}/${BINARY}

windows: vendor
	GOOS=windows ${GO_BUILD} -o bin/windows_${GOARCH}/${BINARY}.exe

test: vendor
	script/run_tests.sh

install:
	cp -Rf bin/ "${PREFIX}/bin"

release: all
	script/release.sh

.PHONY: all test clean vendor linux darwin windows install
