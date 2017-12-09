BINARY = secrets
GOARCH = amd64
PREFIX ?= ${GOPATH}

all: test clean linux darwin windows

clean:
	rm -rf bin/

linux:
	GOOS=linux GOARCH=${GOARCH} go build -o bin/linux_${GOARCH}/${BINARY}

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build -o bin/darwin_${GOARCH}/${BINARY}

windows:
	GOOS=windows GOARCH=${GOARCH} go build -o bin/windows_${GOARCH}/${BINARY}.exe

test:
	go test -v ./...

install:
	cp -Rf bin/ ${PREFIX}/bin

release: all
	script/release.sh

.PHONY: all test clean linux darwin windows install
